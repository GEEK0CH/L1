package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//1 способ:
	ch1 := make(chan string)
	run := true
	go func(run *bool) {
		for {
			select { //дождаться закрытия канала
			case <-ch1:
				fmt.Println("Горутина 1 остановлена!")
				return
			default:
				fmt.Println("Горутина 1 работает...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}(&run)

	time.Sleep(300 * time.Millisecond)
	run = false
	close(ch1) // остановка горутины по закрытию канала
	time.Sleep(300 * time.Millisecond)
	//2 способ:
	ch2 := make(chan int)
	go func() {
		for {
			num, more := <-ch2 //используем второе возвращаемое каналом значение, more - bool переменная, равная false, если канал закрыт
			if !more {
				fmt.Println("Горутина 2 остановлена!")
				return
			}
			fmt.Printf("Горутина 2 работает передает цифру %d\n", num)
		}
	}()
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	close(ch2) // остановка горутины по закрытию канала
	time.Sleep(350 * time.Millisecond)
	//3 способ:
	timer := time.NewTimer(300 * time.Millisecond)
	go func() {
		for {
			select {
			case <-timer.C: //канал таймера time.NewTimer
				fmt.Println("Горутина 3 остановлена!")
				return
			default:
				fmt.Println("Горутина 3 работает...")
				time.Sleep(100 * time.Millisecond)
			}

		}
	}()
	time.Sleep(650 * time.Millisecond)
	//4 способ:
	forever := make(chan int)
	ctx, cancel := context.WithCancel(context.Background()) // контекст с отменой для завершения горутины

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				forever <- rand.Intn(100)
				return
			default:
				fmt.Println("Горутина 4 работает...")
			}

			time.Sleep(350 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(1 * time.Second)
		cancel() //завершение горутины
	}()

	<-forever
	fmt.Println("Горутина 4 остановлена!")
}

