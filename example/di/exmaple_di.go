package factory

import (
	"fmt"
	"time"
)

type ExampleDI struct {
	dep Dependency
}

func NewExampleDI(dep Dependency) *ExampleDI {
	return &ExampleDI{
		dep: dep,
	}
}

func (u *ExampleDI) Print() error {
	for {
		data, err := u.dep.LoadData()
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
		fmt.Println("print - ", data)
	}
}
