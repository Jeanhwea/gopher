package main

import (
	"fmt"
	"runtime"

	"github.com/jeanhwea/gopher/06-packages/pkg01"
)

var (
	countGlobal = 1
)

func test() {
	count := 1
	defer func(count int) {
		fmt.Printf("count = %d, countGlobal = %d\n", count, countGlobal)
	}(count)
	countGlobal++
	count++
	fmt.Printf("count = %d, countGlobal = %d\n", count, countGlobal)
}

func main() {
	service := pkg01.NewGreetService()
	fmt.Printf("Number of Goroutine = %d\n", runtime.NumGoroutine())
	service.Greeting()
	test()
}
