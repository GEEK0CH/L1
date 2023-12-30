package main

import "fmt"

func intersect(slice1 []int, slice2 []int) []int {
	var slice []int
	for _, i := range slice1 { //перебор значений первого слайса
		for _, i2 := range slice2 { //перебор значений второго слайса
			if i == i2 { //при совпадении элементов добавляется в слайс
				slice = append(slice, i)
			}
		}
	}
	return slice
}

func main() {
	s1 := []int{10, 5, 8, 9, 3}
	s2 := []int{6, 2, 7, 5, 10}

	fmt.Println(intersect(s1, s2))
}
