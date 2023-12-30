package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	wg.Add(len(nums)) //добавление количество групп
	for _, num := range nums {
		go func(num int) {
			fmt.Println(num * num)
			wg.Done()
		}(num)
	}
	wg.Wait() //ожидание всех групп
}
