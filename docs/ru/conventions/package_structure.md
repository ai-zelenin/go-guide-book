# Структура пакета

Пакет это строительная единица приложения, он в свою очередь состоит из типов(в основном структур)
/функций/переменных/константы(
глобальные), которые размещаются в файлах.

Пакеты именуются как существительное(желательно в единственном числе) без заглавных букв, допустимы сокращения, аббревиатуры и
просто созвучия. Например, httpcli или iotool или rpcmidware. 
Крайне не желательно давать своим пакетам имена совпадающиее с системными(http,grpc,io).  
[EffectiveGo](https://go.dev/blog/package-names).  

Файлы именуются в стиле snake_case.go  
[Issue](https://github.com/golang/go/issues/36060).  

## Логическая структура

Логику пакета следует разбить на компоненты(типы с методами) с максимальной обособленной ответственностью.  
Каждый компонент должен иметь свой конструктор. Если в проекте используются [Fx](https://github.com/uber-go/fx)
конструкторы для него должны или быть объявлены в отдельном пакете или иметь префикс NewFx.  
Если поведение компонента предполагается конфигурировать снаружи, то нужно описать структуру конфигурации и или
передавать ее в конструктор/рабочий метод или использовать
паттерн [функциональных опций](https://github.com/uber-go/guide/blob/master/style.md#functional-options).    
Компоненты требующее действий не входящих в зону ответственности данного пакета должны требовать в своем конструкторе
реализацию данных функций от внешних зависимостей.

## Размещение по файлам

Размещать описания типов/функций/переменных в файлах стоит следующим образом.

- Полноценный тип, методы которого являются частью полезной нагрузки пакета всегда занимает отдельный файл. В нем
  описана структура данного типа потом его конструктор, потом методы привязанные к данному типу. Желательно, что бы
  последовательность описания методов имела следующую логику:

    1. Метод объекта имеющий зависимость от другого метода(вызывает другой метод этого же объекта)
       идет выше чем тот который вызывают.
    2. Если в названиях методов есть отсылки к процессам,  
       то они располагаются в порядке их логического вызова (Start, work, Stop)

##### Файл с правильно определенным типом

```go
 package factory

import "context"

const DefaultURL = "http://localhost:8080"

type Worker struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewWorker(ctx context.Context) *Worker {
	ctx, cancel := context.WithCancel(ctx)
	return &Worker{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (w *Worker) Start() error {
	for {
		select {
		case <-w.ctx.Done():
			return nil
		default:
			err := w.fetchAndProcess()
			if err != nil {
				return err
			}

		}
	}
}

func (w *Worker) fetchAndProcess() error {
	return nil
}

func (w *Worker) Stop() {
	w.cancel()
}

 ```


- Типы, которые не несут полезной нагрузки, но повсеместно используются в пакете должны быть сгруппированы в файлы
  соответствующими названиями допустим конфигурации в config.go

- Типы, которые являются неотемлемой частью смысловой нагрузки другого типа и не могут быть переиспользованы должны
  размещаться в файле с родительским типом. Допустим DTO объекты должны быть описаны в том же файле, что и http.Handler
  который их использует.

- Функции не привязанные к типам, должны располагаться или в отдельных файлах или файле с названием util.go

- Переменные/константы пакетного уровня размещаются или в файле где они используются или, если они используются в многих
  файлах то файле с описанием структур конфигураций.

### Частые ошибки

##### Описание методов относящихся к другому типу

```go
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

 ```


##### Утилитарные функции/конфигурации в файлах с конкретным типом

```go
 package factory

import "fmt"

type Item2 struct {
	dependency *Dependency
	Data       string
}

func NewItem2(dependency *Dependency) *Item2 {
	return &Item2{
		dependency: dependency,
		Data:       "Item2 data",
	}
}

func (i *Item2) PrintData() error {
	data, err := i.dependency.GetData()
	if err != nil {
		return err
	}
	// Print data to stdout
	fmt.Printf("%s - %s", i.Data, data)
	return nil
}

func GetData2() (string, error) {
	return "some data", nil
}

 ```


Если какая-то часть логики вынесена в отдельную функцию, вероятно это сделано, что бы переиспользовать эту логику,
переиспользуемые компоненты не должны быть размещены(логически сгруппированы) с одним конкретным типом.




