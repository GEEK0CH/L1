package main

import "fmt"

type Human struct {
	Gender string
	Name   string
	Age    int
}

type Action struct { // структура Human
	Human
}

func (h Human) Show() string { //вывод структуры Action от родительской структуры Human
	return fmt.Sprintf("The %s's name is %s,he is %d years old", h.Gender, h.Name, h.Age)
}

func main() {
	a := Action{}
	a.Gender = "Man"
	a.Name = "Mark"
	a.Age = 24
	fmt.Println(a.Show())
}
