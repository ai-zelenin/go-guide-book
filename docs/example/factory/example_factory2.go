package factory

type ExampleFactory2 struct {
	dep  *Dependency
	item *Item
}

func NewExampleFactory2(dep *Dependency, item *Item) *ExampleFactory2 {
	return &ExampleFactory2{
		dep:  dep,
		item: item,
	}
}

func (f *ExampleFactory2) CreateItem() *Item {
	return f.item
}
