package factory

import (
	"fmt"
	"time"
)

var dep = NewDependency()

type ExampleDI2 struct {
}

func NewExampleDI2() *ExampleDI {
	return &ExampleDI{}
}

func (u *ExampleDI2) Print() error {
	for {
		data, err := dep.LoadData()
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
		fmt.Println("print - ", data)
	}
}
