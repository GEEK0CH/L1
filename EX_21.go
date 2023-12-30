package main

import (
	"fmt"
	"strconv"
)

type Subject interface { //интерфейс субъекта
	Request(number string) string
}

type ConcreteSubject struct {
}

func (a *ConcreteSubject) SpecificRequest(number int) int {
	return number * 2
}

type Adapter struct { //адаптер
	ConcreteSubject
}

func (adapter *Adapter) Request(number string) string { //преобразование между двумя интерфейсами
	n, _ := strconv.Atoi(number)
	result := fmt.Sprintf("%d", adapter.SpecificRequest(n))
	return result
}

func clientCode(subject Subject) { //симуляция клиентского блока
	fmt.Println(subject.Request("10"))
}

func main() {
	subject := ConcreteSubject{}
	adapter := Adapter{subject}
	clientCode(&adapter)
}
