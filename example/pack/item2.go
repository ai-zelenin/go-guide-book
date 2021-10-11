package factory

import "fmt"

type Item2 struct {
	dependency *Dependency
	Data       string
}

func NewItem2(dependency *Dependency) *Item2 {
	return &Item2{
		dependency: dependency,
		Data:       "Item2 data",
	}
}

func (i *Item2) PrintData() error {
	data, err := i.dependency.GetData()
	if err != nil {
		return err
	}
	// Print data to stdout
	fmt.Printf("%s - %s", i.Data, data)
	return nil
}

func GetData2() (string, error) {
	return "some data", nil
}
