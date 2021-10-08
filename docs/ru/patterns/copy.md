# Копирование объектов

Как копировать объекты правильно

```go
 package cp

import "fmt"

func ExampleCopy() {
	val := 1
	item := &Item{
		Ptr:     &val,
		Number:  1,
		Text:    "OK",
		Numbers: []int{1, 2, 3},
	}

	newItem := item.Copy()
	newItem.Number = 1
	newItem.Numbers[0] = 99
	fmt.Printf("item - %#v", item)
	fmt.Printf("new item - %#v", newItem)
}

// out
// item - &main.Item{Text:"OK", Number:1, Numbers:[]int{1, 2, 3}}
// new item - main.Item{Text:"OK", Number:1, Numbers:[]int{99, 2, 3}}

 ```


Как копировать объекты неправильно

```go
 package cp

import "fmt"

func ExampleCopy1() {
	item := &Item{
		Number:  1,
		Text:    "OK",
		Numbers: []int{1, 2, 3},
	}

	newItem := *item
	newItem.Number = 1
	newItem.Numbers[0] = 99
	fmt.Printf("item - %#v", item)
	fmt.Printf("new item - %#v", newItem)
}

// out
// item - &main.Item{Text:"OK", Number:1, Numbers:[]int{99, 2, 3}}
// new item - main.Item{Text:"OK", Number:1, Numbers:[]int{99, 2, 3}}

 ```
