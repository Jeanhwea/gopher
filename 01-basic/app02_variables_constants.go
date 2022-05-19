package main

import (
	"fmt"
)

// 全局变量定义
var APP_NAME string
var APP_PORT = 9000

// := 不能用于定义全局变量, 如下定义变量会报错
// bar0 := "hello"
////////////////////////////////////////////////////////////////////////////////
// 基本类型
// bool,
// uint, uint8, uint16, uint32, uint64,
// int, int8, int16, int32, int64,
// float32, float64
func main() {
	// 定义变量的方法有如下几种
	var foo1 int
	var foo2 int = 42
	var foo3, foo4 int = 23, 66
	var foo5 = 42
	// golang 中定义变量不使用会报错, 这样使用一次避免报错
	_, _, _, _, _ = foo1, foo2, foo3, foo4, foo5
	// := 可以在函数里快速定义变量, short definition
	bar1 := 1
	fmt.Println(bar1)
	fmt.Println("Application Port =", APP_PORT)
	////////////////////////////////////////////////////////////////////////////////
	// 定义常量方法如下
	const constant = "This is a constant"
	// 常量中有一个特殊的赋值用法 iota
	const (
		// iota 在 const 括号里面每行累加 1, 第一行默认 0
		BEIJING = iota
		SHANGHAI
		SHENZHEN
	)
	fmt.Println(SHANGHAI)
	fmt.Println(SHENZHEN)
}
