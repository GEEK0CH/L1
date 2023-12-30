package main

import (
	"fmt"
	"strings"
)

func uniqueSimbol(str string) bool {
	m := make(map[string]int)
	str = strings.ToLower(str)  //перемещение в нижний регистр
	for _, value := range str { // цикл для перебора строки
		if m[string(value)] == 0 { //если символа нет в мапе мы добовляем его
			m[string(value)] = 1
		} else { //в любом другой случае возвращаем false
			return false
		}
	}
	return true
}

func main() {
	str := "abcd"
	str2 := "abCdefAaf"

	fmt.Println(uniqueSimbol(str))
	fmt.Println(uniqueSimbol(str2))
}
