# DI(инъекция зависимостей)

Это паттерн проектирования который является частным случаем IoC(инверсии контроля) и предполагает передачу выполнения
какой-то функции на внешнюю зависимость. Применяется для того что бы поддерживать принцип единой ответственности и
упрощать последующую модификацию компонентов. Допустим у нас есть объект, который умеет делать что полезное(в примере
печатать что-то с интервалом по времени). Что бы его переиспользовать, нам нужно отделить ту часть которая будет
делегироваться во вне(в пример это получения данных). В рамках паттерна DI очень удобно пользовался абстракциями(
интерфейсами), это позволяет потом переиспользовать написанные компоненты с другими зависимостями, допустим в unit
тестировании мы можем прокидывать в качестве зависимостей моки или подменять имплементации DAO объектов. В случае если
объекты внешних зависимостей могут быть изменены в процессе использования то в качестве зависимости можно прокидывать
фабрику, которая будет создавать интересующие нас объекты.

```go
 package factory

import (
	"fmt"
	"time"
)

type ExampleDI struct {
	dep Dependency
}

func NewExampleDI(dep Dependency) *ExampleDI {
	return &ExampleDI{
		dep: dep,
	}
}

func (u *ExampleDI) Print() error {
	for {
		data, err := u.dep.LoadData()
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
		fmt.Println("print - ", data)
	}
}

 ```


В итоге наш объект будет делать только одно дело, а решение о том что именно он будет печатать будет делегированное
наружу.

```go
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

 ```


Если мы просто создаем объекты там где они нам понадобились не прокидывая их в конструктор, то тем самым скрываем факт
наличия зависимости у объекта и если в качестве зависимости используется абстракция, то еще не позволяем изменить ее
имплементацию при создании.

```go
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

 ```


Глобальные переменные не являются правильной реализацией паттерна DI, так как ограничивают использование разных
имплементаций для разных объектов использующих данную зависимость.

##### Дополнительная информация

1. [Wiki DI](https://ru.wikipedia.org/wiki/%D0%92%D0%BD%D0%B5%D0%B4%D1%80%D0%B5%D0%BD%D0%B8%D0%B5_%D0%B7%D0%B0%D0%B2%D0%B8%D1%81%D0%B8%D0%BC%D0%BE%D1%81%D1%82%D0%B8)
2. [Wiki IoC](https://ru.wikipedia.org/wiki/%D0%98%D0%BD%D0%B2%D0%B5%D1%80%D1%81%D0%B8%D1%8F_%D1%83%D0%BF%D1%80%D0%B0%D0%B2%D0%BB%D0%B5%D0%BD%D0%B8%D1%8F)
3. [Wiki Coupling](https://ru.wikipedia.org/wiki/%D0%97%D0%B0%D1%86%D0%B5%D0%BF%D0%BB%D0%B5%D0%BD%D0%B8%D0%B5_(%D0%BF%D1%80%D0%BE%D0%B3%D1%80%D0%B0%D0%BC%D0%BC%D0%B8%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5))