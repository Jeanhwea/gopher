package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

func main() {
	var p unsafe.Pointer
	newP := 42
	// CAS 操作
	atomic.CompareAndSwapPointer(&p, nil, unsafe.Pointer(&newP))

	v := (*int)(p)
	fmt.Println(*v)
}
