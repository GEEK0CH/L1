package main

import (
	"fmt"
	"strings"
)

func reverseWord(str string) (strN string) {
	var idx = len(str)                   //поинтер на конец слова
	for i := len(str) - 1; i >= 0; i-- { //цикл для поиска пробела
		if string(str[i]) == " " { //при нахождении пробела записывается слово в новую строку с пробелом
			strN += str[i+1:idx] + " "
			idx = i
		} else if i == 0 { //если слово последние слово записывается без пробела
			strN += str[i:idx]
		}
	}
	return
}

func reverseWord2(str string) (strN string) {
	wordsArr := strings.Split(str, " ")       //отделение слов от пробелов и запись в массив
	for i := len(wordsArr) - 1; i >= 1; i-- { //цикл для переворачивания порядка слов и записи в новую строку
		strN += wordsArr[i] + " "
	}
	strN += wordsArr[0]
	return
}

func main() {
	str := "snow dog sun"

	fmt.Println(reverseWord(str))
	fmt.Println(reverseWord2(str))
}
