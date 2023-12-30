package main

import (
	"fmt"
	"time"
)

func sleep(numb time.Duration) {
	<-time.After(numb) //функция для ожидания истечения заданного времени
}

func main() {
	fmt.Println("wait...")
	sleep(5 * time.Second)
	fmt.Println("ready!!!")
}
