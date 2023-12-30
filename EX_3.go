package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	wg.Add(len(nums))
	var sum int //добавление переменной для суммы
	for _, num := range nums {
		go func(num int) {
			sum += num * num //прибавление к переменной
			wg.Done()
		}(num)
	}
	wg.Wait()

	fmt.Println(sum)
}
