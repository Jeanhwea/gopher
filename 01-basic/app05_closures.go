package main

import ("fmt")

// 闭包, 返回局部变量 outer_var
func scope() func() int {
	outer_var := 2
	foo := func() int { return outer_var }
	return foo
}

func another_scope() func() int {
	// outer_var 和 foo 未定义, 编译错误
	outer_var = 444
	return foo
}

// 闭包
func outer() (func() int, int) {
	outer_var := 2
	inner := func() int {
		outer_var += 99 // outer_var 被修改了
		return outer_var
	}
	inner()
	return inner, outer_var // 返回值中 outer_var 被修改成了 101
}

func main() {
	// 将函数复制给变量
	add := func(a, b int) int {
		return a + b
	}
	// 调用函数
	fmt.Println(add(3, 4))
}
