package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 给随机数生成器添加种子
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		// 生成小于 10 的随机数
		fmt.Println(rand.Intn(10))
	}
}
