package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var channelNums, channelNumsN = make(chan int, len(nums)), make(chan int, len(nums)) //создание каналов для операции

	for _, num := range nums { //передача чисел в канал
		channelNums <- num
	}
	close(channelNums)

	for num := range channelNums { //передача чисел из канала в канал с умножением на 2
		channelNumsN <- num * 2
	}
	close(channelNumsN)

	for numb := range channelNumsN { //вывод цифр из канала
		fmt.Println(numb)
	}
}
