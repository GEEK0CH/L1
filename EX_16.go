package main

import "fmt"

func quickSort(arr []int, left, right int) []int {
	if left < right {
		var p int
		arr, p = partition(arr, left, right)
		//рекурсия для левой и правой стороны
		arr = quickSort(arr, left, p-1) //
		arr = quickSort(arr, p+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) ([]int, int) {
	pivot := arr[right] //берет самый последний элемент для ориентира сортировки
	i := left
	for j := left; j < right; j++ { //все что меньше числа ориентира он помещает в левую часть, а все что больше в правую
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return arr, i
}

func main() {
	array := []int{2, 10, 33, 29, 5, 9, 23, 45, 79, 1}
	fmt.Println(quickSort(array, 0, len(array)-1))
}
