package factory

type ExampleFactory1 struct {
	dep  *Dependency
	item *Item
}

func NewExampleFactory1(dep *Dependency) *ExampleFactory1 {
	return &ExampleFactory1{
		dep:  dep,
		item: NewItem(dep),
	}
}

func (f *ExampleFactory1) CreateItem() *Item {
	return f.item
}
