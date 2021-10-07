package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var src string
var dst string

func main() {
	flag.StringVar(&src, "src", ".", "path to file or dir")
	flag.StringVar(&dst, "dst", ".", "path to file or dir")
	flag.Parse()

	err := filepath.Walk(src, func(p string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return handlePath(p, info)
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func handlePath(path string, info fs.FileInfo) error {
	if info.IsDir() {
		return nil
	}
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	result, err := handleFile(file, path, info)
	if err != nil {
		return err
	}
	var resultPath string
	if strings.HasSuffix(dst, "/") {
		err = os.MkdirAll(filepath.Join(dst, filepath.Dir(path)), os.ModePerm)
		if err != nil {
			return err
		}
		resultPath = filepath.Join(dst, filepath.Dir(path), info.Name())
	} else {
		resultPath = dst
	}
	fmt.Printf("%s -> %s \n", path, resultPath)
	data, err := ioutil.ReadAll(result)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(resultPath, data, os.ModePerm)
}

func handleFile(reader io.Reader, path string, info fs.FileInfo) (io.Reader, error) {
	parent, _, ext := fileNameSplit(path, info)
	switch strings.ToLower(ext) {
	case ".md":
		linkRegexp, err := regexp.Compile(`\[[^\]]*\]\([^)]*\)`)
		if err != nil {
			return nil, err
		}
		ptRegext, err := regexp.Compile(`\([^)]*\)`)
		if err != nil {
			return nil, err
		}
		buf := bytes.NewBufferString("")
		data, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}
		res := linkRegexp.ReplaceAllFunc(data, func(i []byte) []byte {
			linkData := ptRegext.Find(i)
			link := strings.TrimPrefix(string(linkData), "(")
			link = strings.TrimSuffix(link, ")")
			name := strings.TrimSuffix(string(i), string(linkData))
			result, err := createSpoilerCodeByLink(name, parent, link)
			if err != nil {
				log.Println(err)
				return i
			}
			if result == nil {
				return i
			}
			return result
		})
		buf.Write(res)
		return buf, nil
	}
	return reader, nil
}

func createSpoilerCodeByLink(origin, rootDir, link string) ([]byte, error) {
	sourceFileExt := filepath.Ext(link)
	if strings.ToLower(sourceFileExt) == ".md" {
		return nil, nil
	}
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	codeExt := strings.TrimPrefix(filepath.Ext(link), ".")
	err = os.Chdir(rootDir)
	if err != nil {
		return nil, err
	}
	fileData, err := ioutil.ReadFile(link)
	if err != nil {
		return nil, err
	}
	err = os.Chdir(pwd)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	//buf.WriteString("<details>\n")
	//buf.WriteString(fmt.Sprintf("<summary>%s</summary>\n\n", origin))
	buf.WriteString(fmt.Sprintf("```%s\n %s\n ```\n", codeExt, fileData))
	//buf.WriteString("</details>")
	return buf.Bytes(), nil
}

func fileNameSplit(path string, info fs.FileInfo) (parent string, name string, ext string) {
	basename := info.Name()
	ext = filepath.Ext(basename)
	name = strings.TrimSuffix(basename, ext)
	parent = filepath.Dir(path)
	return parent, name, ext
}
