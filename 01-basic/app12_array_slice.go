package main

import "fmt"

func arrayTutor() {
	var arr1 [3]int
	arr1[0] = 4
	fmt.Println(arr1)
	fmt.Println(len(arr1)) // 获取数组元素个数

	// 遍历数组
	arr2 := [3]int{1, 2, 3}
	for k, v := range arr2 {
		fmt.Println(k, v)
	}

	// ... 表示编译器自动推断数组长度
	arr3 := [...]int{1, 2}
	fmt.Println(arr3)

	fmt.Println("hello")

	// 数组是传值的, 并且数组作为函数传参也是 pass by value
	cities := [...]string{"ShangHai", "BeiJing"}
	citiesCopy := cities
	citiesCopy[0] = "WuHu"
	fmt.Println(cities)     // => [ShangHai BeiJing]
	fmt.Println(citiesCopy) // => [WuHu BeiJing]

	// 多维数组, 注意下标
	mat1 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(mat1)
	fmt.Println(mat1[1][0]) // => 4
}

func sliceTutor() {
	// 切片定义
	arr1 := [5]int{3, 4, 5, 6}
	var slc1 []int = arr1[1:3]
	// slice[start:end] 返回 [start,end) 的数据
	fmt.Println(slc1) // => [4, 5]

	// 使用 make 创建切片
	slc2 := make([]int, 5, 5) // => make([]T, len, cap)
	fmt.Println(slc2)

	// 切片是传引用的, pass by reference
	slc3 := slc2
	slc3[3] = 8
	fmt.Println(slc2) // => [0 0 0 8 0]
	fmt.Println(slc3) // => [0 0 0 8 0]

	// append 追加元素
	slc2 = append(slc2, 9)
	fmt.Println(slc2)

	// 合并两个切片, 使用 append 和 ... 操作符
	slc4 := append(slc2, slc3...)
	fmt.Println(slc4)

	// 使用 copy 函数来克隆切片
	slc5 := make([]int, len(slc2))
	copy(slc5, slc2)
	slc5[0] = 6
	fmt.Println(slc2) // => [0 0 0 8 0 9]
	fmt.Println(slc5) // => [6 0 0 8 0 9]
}

func main() {
	arrayTutor()
	sliceTutor()
}
