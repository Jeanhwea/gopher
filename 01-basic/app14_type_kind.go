package main

import (
	"fmt"
	"unsafe"
)

type (
	MyInt int64 // => int64
	Age   MyInt // => MyInt => int64
	// 上面的是一种底层类型, 都为 int64
)

type (
	IntSlice   []int   // => []int
	MyIntSlice []MyInt // => []MyInt
	AgeSlice   []Age   // => []Age
	// golang 的类型推导到底层类型结束, 所有上面是三种底层类型
)

func main() {
	var int01 MyInt = 2
	var int02 Age = 9
	fmt.Printf("sizeof(int01) = %d, sizeof(int02) = %d\n", unsafe.Sizeof(int01), unsafe.Sizeof(int02))

	var slice01 IntSlice
	fmt.Println(unsafe.Sizeof(slice01))

	var array01 [100]int
	fmt.Println(unsafe.Sizeof(array01))
}
