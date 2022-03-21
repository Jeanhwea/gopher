package main

import (
	"fmt"
	"runtime"
	"time"
)

func print() {
	count := 0
	for {
		fmt.Println(count)
		count++
		time.Sleep(1 * time.Second)
		if count > 10 {
			runtime.Goexit() // 退出协程, 但是主协程不退出
		}
	}
}

func main() {
	// 开启协程, 但是不阻塞
	go print()

	for {
		time.Sleep(1 * time.Second)
	}
}
