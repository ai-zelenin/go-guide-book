package factory

type ExampleFactory struct {
	dep *Dependency
}

func NewExampleFactory(dep *Dependency) *ExampleFactory {
	return &ExampleFactory{
		dep: dep,
	}
}

func (f *ExampleFactory) CreateItem() *Item {
	return NewItem(f.dep)
}
