package factory

import "fmt"

type Item1 struct {
	dependency *Dependency
	Data       string
}

func NewItem1(dependency *Dependency) *Item1 {
	return &Item1{
		dependency: dependency,
		Data:       "Item1 data",
	}
}

func (i *Item1) PrintData() error {
	data, err := i.dependency.GetData()
	if err != nil {
		return err
	}
	// Print data to stdout
	fmt.Printf("%s - %s", i.Data, data)
	return nil
}

func (d *Dependency) GetData1() (string, error) {
	return "some data from Dependency 1", nil
}
