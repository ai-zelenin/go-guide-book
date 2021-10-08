package factory

import (
	"fmt"
	"time"
)

type ExampleDI1 struct {
}

func NewExampleDI1() *ExampleDI {
	return &ExampleDI{}
}

func (u *ExampleDI1) Print() error {
	for {
		d := NewDependency()
		data, err := d.LoadData()
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
		fmt.Println("print - ", data)
	}
}
