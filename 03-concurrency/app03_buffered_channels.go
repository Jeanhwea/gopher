package main

import (
	"fmt"
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

// 带缓存的管道

func test01() {
	ch := make(chan string, 2) // 管道容量为 2

	ch <- "Jack"
	ch <- "Tom"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func test02() {
	chSeq := make(chan int, 2)

	go func(ch chan int) {
		for i := 0; i < 7; i++ {
			ch <- i
			log.Println("Successfully write value", i)
		}
		close(ch)
	}(chSeq)
	time.Sleep(2 * time.Second)

	for val := range chSeq {
		log.Println("Read value", val)
		time.Sleep(2 * time.Second)
	}
}

func test11() {
	fmt.Println("Main is running!")

	// 添加带 3 个缓存的 channel
	ch := make(chan int, 3)

	go func() {
		defer fmt.Println("Goroutine is done!")

		fmt.Println("Goroutine is running!")
		for i := 0; i < 3; i++ {
			ch <- i
			fmt.Printf("Goroutine: len = %d, cap = %d\n", len(ch), cap(ch))
		}
	}()

	// channel 有缓存, 子协程不会阻塞
	for i := 0; i < 3; i++ {
		num := <-ch
		fmt.Println(num)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Main is done!")
}

func main() {
	test02()
}
