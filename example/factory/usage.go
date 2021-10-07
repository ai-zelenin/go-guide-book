package factory

type Usage struct {
	factory *ExampleFactory
	item    *Item
}

func NewUsage(factory *ExampleFactory) *Usage {
	return &Usage{
		factory: factory,
		item:    factory.CreateItem(),
	}
}

func (u *Usage) CreateAndUseItem() error {
	item := u.factory.CreateItem()
	item.Data = "Usage data"
	return item.PrintData()
}

func (u *Usage) UseItem() error {
	return u.item.PrintData()
}
