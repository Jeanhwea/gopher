package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := *new([]int) // => a = nil
	fmt.Printf("%p:%v\n", a, a)

	x := []byte{4: 'x'}
	fmt.Println(len(x))

	hello := "hello"
	fmt.Println(hello[len(hello):])

	emptyStruct := struct{}{}
	fmt.Println(unsafe.Sizeof(emptyStruct)) // struct{} 的长度为零，不占用内存空间

	// 一个空的管道用来进行协程同步
	done := make(chan struct{})
	go func() {
		fmt.Println("In inner function")
		done <- struct{}{}
	}()

	<-done
}
