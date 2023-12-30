package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var count int
	fmt.Println("Введите количество секунд:")
	fmt.Scanf("%d\n", &count)                                  //считывание количества секунд работы
	workerVal := make(chan int)                                // можем использовать небуферизированный канал, "синхронно" пишем и читаем
	timer := time.NewTimer(time.Duration(count) * time.Second) // выход из программы по таймеру

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() { // горутина для чтения из канала
		for {
			num, more := <-workerVal
			if more { //если канал открыт выводим значение из канала, в любом другом случае закрываем группу
				fmt.Printf("Значение из канала: %d\n", num)
			} else {
				wg.Done()
				return
			}
		}
	}()

mainLoop:
	for {
		select { // для работы с каналом от таймера используем select

		case <-timer.C: // событие от таймера в канале timer.C
			fmt.Printf("Время вышло!\n")
			close(workerVal) // закрываем канал, уведомляя горутину о завершении работы
			break mainLoop   // завершаем работу внешнего цикла под меткой! а не всего блока mainLoop

		default: // если таймер еще не сработал - посылаем в канал данные
			workerVal <- rand.Intn(1000-1) + 1
			time.Sleep(time.Millisecond * 600)
		}
	}

	wg.Wait()
	fmt.Println("Конец")
}
