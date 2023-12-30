package main

import "fmt"

func binarySearch(array []int, numb int) int {
	var left = 0
	var right = len(array) - 1

	for left <= right { //цикл через левую и правую границу до середины
		var mid = (left + right) / 2
		var midNumb = array[mid] //получение центрального элемента

		if midNumb == numb { //если элемент не найден, то в зависимости от числа цикл шагнет влево или вправо
			return mid
		} else if midNumb < numb {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -(left + 1) //возвращение несуществующего элемента если элемент не был найден
}

func main() {
	array := []int{2, 10, 33, 29, 5, 9, 23, 45, 79, 1}
	fmt.Println(array)
	fmt.Println(binarySearch(array, 79))
}
