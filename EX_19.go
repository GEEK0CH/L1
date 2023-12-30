package main

import "fmt"

func reverse(str string) (strN string) {
	for _, values := range str { // перебор каждого символа строки str
		strN = string(values) + strN // добавление символа к строке strN с конца
	}
	return
}

func main() {
	str := "главрыба"
	str2 := "🙈🙉🙊"

	fmt.Println(reverse(str))
	fmt.Println(reverse(str2))
}
