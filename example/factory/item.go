package factory

import "fmt"

type Item struct {
	dependency *Dependency
	Data       string
}

func NewItem(dependency *Dependency) *Item {
	return &Item{
		dependency: dependency,
		Data:       "Item data",
	}
}

func (i *Item) PrintData() error {
	data, err := i.dependency.GetData()
	if err != nil {
		return err
	}
	// Print data to stdout
	fmt.Printf("%s - %s", i.Data, data)
	return nil
}
