package main

import (
	"fmt"
	"sync"
	"time"
)

type cc struct {
	Value int
	Mutex sync.Mutex
}

func (c *cc) count() {
	c.Mutex.Lock()
	c.Value++ //счетчик
	c.Mutex.Unlock()
}

func main() {
	c := cc{}
	structInd := make(chan struct{}) //структура по закрытию которой счетчик выведет результат

	go func() {
		for {
			select {
			case _, ok := <-structInd:
				if !ok {
					break
				}
			default:
			}
			c.count()
		}
	}()

	time.Sleep(600 * time.Millisecond)
	close(structInd)
	fmt.Printf("Counter:%d\n", c.Value)
}
