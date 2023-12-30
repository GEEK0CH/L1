package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(worker int, in <-chan int) {
	for {
		num := <-in
		fmt.Printf("Воркер %d цифра %d\n", worker, num)
	}
}
func main() {
	var count int
	fmt.Println("Введите количество Воркеров:")
	fmt.Scanf("%d\n", &count) //считывание количества воркеров
	workerVal := make(chan int)

	for i := 0; i < count; i++ { // создаем воркеров
		go worker(i, workerVal)
	}

	for { // постоянная запись рандомного числа от 1 до 1000
		workerVal <- rand.Intn(1000-1) + 1
		time.Sleep(time.Millisecond * 600)
	}
}
