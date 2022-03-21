package main

import (
	"fmt"
	"sync"
)

var x = 0

func inc(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	x = x + 1
	mu.Unlock()
	wg.Done()
}

func test01() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 9999; i++ {
		wg.Add(1)
		go inc(&wg, &mu)
	}
	wg.Wait()
	fmt.Println("x =", x)
}

func main() {
	test01()
}
