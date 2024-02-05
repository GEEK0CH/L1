package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type cm struct {
	Mutex sync.Mutex
	Data  map[int]int
}

func (m *cm) Add(key int, value int) {
	m.Mutex.Lock()
	m.Data[key] = value //запись данных по ключу
	m.Mutex.Unlock()
}

func main() {
	m := cm{
		Mutex: sync.Mutex{},
		Data:  make(map[int]int),
	}
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ { //запись рандомных значений в массив через функцию Add
		go func(wg *sync.WaitGroup) {
			m.Add(rand.Intn(1000), rand.Intn(1000))
			wg.Done()
		}(&wg)
	}
	wg.Wait()
	fmt.Print(m.Data)
}
