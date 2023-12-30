package main

import (
	"fmt"
	"reflect"
)

func typeO(inf []interface{}) {
	for _, values := range inf { //перебор переменных в интерфейсе
		fmt.Println(reflect.TypeOf(values))
	}
}
func main() {
	typeO([]interface{}{make(chan string), make(map[int]int), 0.09, "str"})
}
