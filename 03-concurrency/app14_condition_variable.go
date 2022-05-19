package main

import (
	"fmt"
	"sync"
)

var cond = sync.NewCond(new(sync.Mutex))
var count = 0 // 计算器

// 生产者

func consumer() {
	for {
		cond.L.Lock()
		for count == 0 {
			cond.Wait()
		}
		count--
		fmt.Printf("Consumer: %d\n", count)
		cond.Signal()
		cond.L.Unlock()
	}
}

// 消费者

func producer() {
	for {
		cond.L.Lock()
		for count == 100 {
			cond.Wait()
		}
		count++
		fmt.Printf("Producer: %d\n", count)
		cond.Signal()
		cond.L.Unlock()
	}
}

func main() {
	go consumer()
	go consumer()
	producer()
}
