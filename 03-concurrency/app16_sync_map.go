package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建
	smap := sync.Map{}

	// 写入
	smap.Store("apple", 5)
	smap.Store("banana", 6)
	var wg sync.WaitGroup
	for i := 0; i < 99; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			smap.Store("one", 3)
		}()
	}

	// 读取
	length, _ := smap.Load("apple")
	fmt.Println(length)

	// 遍历
	smap.Range(func(key, val interface{}) bool {
		fruit := key.(string)
		length := val.(int)
		fmt.Printf("len(%s)=%d\n", fruit, length)
		return true
	})

	// 删除
	smap.Delete("one")

	// 读取或写入
	smap.LoadOrStore("pear", 4)

	wg.Wait()
	fmt.Println(smap)
}
