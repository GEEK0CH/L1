package main

import (
	"fmt"
)

//var justString string // лучше не использовать глобальную переменную, они нарушают инкапсуляцию и еще множество других правил написания кода

//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100] //будет вызвана утечка памяти вызванная подстраками
//}
//
//func main() {
//	someFunc()
//}

func createHugeString() (str string) {
	for i := 0; i <= 100; i++ {
		str += "+"
	}
	return
}

func someCorrectFunc() (newSlice []byte) {
	v := createHugeString()
	hugeStringByte := []byte(v)
	justStringByte := hugeStringByte[:100]       // оставил срез, для того, чтобы указать диапазон данных
	newSlice = make([]byte, len(justStringByte)) // создаем пустой срез с заданной длиной
	copy(newSlice, hugeStringByte)
	return
}

func main() {
	fmt.Println(string(someCorrectFunc()))
}
