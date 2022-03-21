package main

import (
	"fmt"
	"unsafe"
)

func setTutor() {
	// golang 不支持内置集合（set）类型。但是，集合类型可以用轻松使用映射类型来模拟
	set1 := make(map[string]struct{})

	// 类型struct{} 的尺寸为零，所以此映射类型的值中的元素不消耗内存
	fmt.Println(unsafe.Sizeof(set1))
	_ = set1
}

func main() {
	// 使用 make(...) 直接声明并创建一个 map 对象
	emplSalary := make(map[string]int)
	emplSalary["Jack"] = 5000
	emplSalary["Tom"] = 5500
	fmt.Println(emplSalary) // => map[Jack:5000 Tom:5500]

	delete(emplSalary, "Tom")
	fmt.Println(emplSalary)

	emplSalary["Liu"] = 9000
	if value, ok := emplSalary["Liu"]; ok {
		fmt.Println("found Liu", value)
	}

	for k, v := range emplSalary {
		fmt.Printf("empl = %s, salary = %d\n", k, v)
	}

	// 检测是否包括 "Liu" 作为键, 如果存在执行相应逻辑
	if value, ok := emplSalary["Liu"]; ok {
		fmt.Println(value)
	}

	setTutor()
}
