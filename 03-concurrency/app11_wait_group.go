package main

import (
	"log"
	"sync"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

////////////////////////////////////////////////////////////////////////////////
// WaitGroup

func test01() {
	process := func(rid int, wg *sync.WaitGroup) {
		log.Printf("Start Goroutine(#%d).\n", rid)
		startTime := time.Now()
		time.Sleep(time.Duration(rid) * time.Second)
		duration := time.Now().Sub(startTime).Milliseconds()
		wg.Done()
		log.Printf("Goroutine(#%d) finished in %d milliseconds.\n", rid, duration)
	}

	numGoruntine := 5

	var wg sync.WaitGroup
	for i := 0; i < numGoruntine; i++ {
		wg.Add(1) // 添加一个等待的协程
		go process(i, &wg)
	}

	// 主线程阻塞等待
	wg.Wait()
	log.Println("All GOROUTINE finished.")
}

func main() {
	test01()
}
