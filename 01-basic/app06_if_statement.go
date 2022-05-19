package main

import (
	"fmt"
)

func testIf() int {
	var x int = 10
	if x > 10 {
		return x
	} else if x == 10 {
		return 10
	} else {
		return -x
	}
}

func testIf2(b, c int) int {
	// if 条件前可以添加一个赋值语句
	if a := b + c; a < 42 {
		return a
	} else {
		return a - 42
	}
}

func testIf3() {
	// 类型断言, 通过 if 来检测
	var val interface{} = "foo"
	if str, ok := val.(string); ok {
		fmt.Println(str)
	}
}

func main() {
	testIf()
}
