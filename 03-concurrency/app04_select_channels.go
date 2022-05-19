package main

import (
	"fmt"
	"time"
)

// 通过 select 实现多路复用
func fib(data, quit chan int) {
	x, y := 1, 0
	for {
		select {
		case data <- x:
			x, y = x+y, x
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func test01() {
	data := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-data)
		}
		quit <- 0
	}()

	fib(data, quit)
}

////////////////////////////////////////////////////////////////////////////////
func server1(ch chan string) {
	ch <- "from server1"
	close(ch)
}

func server2(ch chan string) {
	ch <- "from server2"
	close(ch)
}

func test02() {
	out1 := make(chan string)
	out2 := make(chan string)
	go server1(out1)
	go server2(out2)

	time.Sleep(1 * time.Second)
	// 随机从 server1 或 server2 中读取一个字符串
	select {
	case s1 := <-out1:
		fmt.Println(s1)
	case s2 := <-out2:
		fmt.Println(s2)
	}
}

func main() {
	test02()
}
