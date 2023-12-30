package main

import "fmt"

func bunch(mass []string) []string {
	var new []string
	for i := 0; i <= len(mass)-1; i++ { //цикл для перебора индекса массива mass
		var check = true
		for _, a := range new { //перебор массива new
			if a == mass[i] { // при нахождении в массиве new схожей строки check приравнивается false
				check = false
			}
		}
		if check == true { // если check = false строка записывается в массив new
			new = append(new, mass[i])
		}
	}
	return new
}

func main() {
	var s = []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(bunch(s))
}
