package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var bufferCreated int32

func createBuffer() interface{} {
	atomic.AddInt32(&bufferCreated, 1)
	buffer := make([]byte, 1024)
	return &buffer
}

func main() {
	// Pool.New 可能会被并发调用
	bufferPool := &sync.Pool{
		New: createBuffer,
	}

	workerCount := 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()
			buffer := bufferPool.Get() // 申请一个 buffer 实例
			// buffer := createBuffer()
			_ = buffer.(*[]byte)
			defer bufferPool.Put(buffer) // 释放一个 buffer 实例
		}()
	}

	wg.Wait()
	fmt.Printf("%d buffer objects were created\n", bufferCreated)
}
