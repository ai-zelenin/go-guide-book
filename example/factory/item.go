package factory

import "fmt"

type Item struct {
	Dependency *Dependency
}

func NewItem(dependency *Dependency) *Item {
	return &Item{Dependency: dependency}
}

func (i *Item) DoSomething() error {
	data, err := i.Dependency.GetData()
	if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}
