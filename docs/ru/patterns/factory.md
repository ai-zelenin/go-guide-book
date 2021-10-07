# Фабрика

Фабрика паттерн порождения который передусматривает объект создающий другие объекты.

```go
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

 ```


```go
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

 ```


Это удобно в случаях если создание объекта является сложным процессом требующим внешних зависимостей. Объект, который
будет использовать Item не должен знать что-либо о его внутреннем устройстве(допустим что Item для работы нужен
Dependency). Соответственно мы вместо того что бы вызывать конструктор Item напрямую и тем самым заставить Usage
зависеть от Dependancy мы ему передали фабрику.

### Частые ошибки

```go
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

 ```


Хранить в объекте фабрики то что она должна создавать неправильно. Исходя из названия предполагается что фабрика
возвращает новые объекты, а не ссылки на уже существующие. Это особенно критично если получаемый из фабрики объект может
быть модифицирован/конфигурирован перед использованием, или если содержит лимитированные ресурсы(допустим пул соединений
или буфер)

```go
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

 ```


Для создания фабрики требовать те объекты которые потом она сама же будет производить так же неправильно.
