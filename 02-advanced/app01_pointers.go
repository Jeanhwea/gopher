package main

import (
	"fmt"
)

// 指针传参
func changeValue(val *int) {
	*val = 666
}

// 与 C/C++ 不同, golang 返回指针的变量会逃逸到堆中
// 即 golang 中返回指针变量是安全的
func varEscape() *int {
	i := 789
	return &i
}

// 数组传参避免使用指针, 尽量考虑切片
func modifyArr(arr []int) {
	arr[0] = 99
}

func main() {
	// 定义指针
	b := 234
	var a *int = &b
	fmt.Printf("Type of a is [%T], address of b is 0x%x\n", a, a)
	// 空指针
	var p1 *int
	if p1 == nil {
		fmt.Println("p1 is zero value pointer")
	}
	p1 = a
	if p1 == nil {
		fmt.Println("p1 is zero value pointer")
	}
	// 修改值
	*a = 666
	fmt.Println("b is changed to", b)
	// 局部参数逃逸
	heapI := varEscape()
	fmt.Println("heapI", *heapI)
	// 切片传指针
	arr01 := [3]int{1, 2, 3}
	modifyArr(arr01[:])
	fmt.Println(arr01) // => [99 2 3]
	// 注意: golang 中的指针不能进行算术运算
	// p1++ 会报错
}
