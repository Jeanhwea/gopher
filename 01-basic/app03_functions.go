package main

import (
	"fmt"
)

// 定义函数

func functionName() {}

// 定义带参数的函数

func functionName2(param1 string, param2 int) {}

// 合并参数

func functionName3(param1, param2 int) {}

// 返回函数

func functionName4() int {
	return 42
}

// 一次返回多个函数

func returnMulti() (int, string) {
	return 42, "foobar"
}

var x, str = returnMulti()

// 命名函数的返回值

func returnMulti2() (n int, s string) {
	n = 42
	s = "foobar"
	// n and s will be returned
	return
}

// 主函数

func main() {
	var x2, str2 = returnMulti2()
	fmt.Println("x2", x2)
	fmt.Println("str2", str2)
}
