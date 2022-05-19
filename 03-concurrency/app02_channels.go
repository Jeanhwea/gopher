package main

import (
	"fmt"
	"time"
)

// 管道的创建
func test01() {
	ch := make(chan int)
	// 从管道 ch 中读取数据到 data 中
	data := <-ch
	// 将 data 写入管道
	ch <- data
}

// 单向管道
func test02() {
	// send-only 类型的管道
	unidirectChan := make(chan<- int)
	go func(sendChan chan<- int) {
		sendChan <- 10
	}(unidirectChan)
	// 读取 send-only 管道会报错
	// fmt.Println(<-unidirectChan)

	bidirectChan := make(chan int)
	go func(sendChan chan<- int) {
		sendChan <- 10
	}(bidirectChan) // 调用时，隐含将管道类型转换
	fmt.Println(<-bidirectChan)
}

// 使用 range 循环来读取管道
func test03() {
	chSeq := make(chan int)

	// 生产者
	go func(ch chan int) {
		for i := 0; i < 6; i++ {
			ch <- i
		}
		close(ch)
	}(chSeq)

	// 消费者
	for {
		if val, ok := <-chSeq; ok {
			fmt.Println("Received", val, ok)
		} else {
			break
		}
	}
}

func test04() {
	chSeq := make(chan int)

	// 生产者
	go func(ch chan int) {
		for i := 0; i < 6; i++ {
			ch <- i
		}
		close(ch)
	}(chSeq)

	// 消费者, 直接使用 for-range 来读取, 如果管道关闭退出循环
	for data := range chSeq {
		fmt.Println("Received", data)
	}
}

// 使用管道编排协程
func test11() {
	fmt.Println("Main is running!")
	cha := make(chan int)

	go func() {
		defer fmt.Println("Goroutine is done!")

		fmt.Println("Goroutine is running!")
		time.Sleep(1 * time.Second)

		// 往管道中写入数据
		cha <- 666
	}()

	// channel 读取数据时会阻塞, 可以同步协程
	num := <-cha
	fmt.Println(num)

	fmt.Println("Main is done!")
}

func main() {
	test04()
}
