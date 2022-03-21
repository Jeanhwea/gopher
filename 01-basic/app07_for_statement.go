package main

import "fmt"

// golang 中只有 for 循环, 没有 while, loop 等其它循环类型
func testFor() {
	for i := 1; i < 10; i++ {
	}

	var i int
	for i < 10 { // 类似 while 循环
	}
	for i < 10 { // 如果只有条件, 可以省略分号
	}
	for { // 死循环
	}
}

// 带标签的 for 循环
func testFor2() {
here:
	for i := 0; i < 2; i++ {
		for j := i + 1; j < 3; j++ {
			if i == 0 {
				continue here // 继续外层循环
			}
			fmt.Println(j)
			if j == 2 {
				break // 跳出内层循环
			}
		}
	}

there:
	for i := 0; i < 2; i++ {
		for j := i + 1; j < 3; j++ {
			if j == 1 {
				continue // 继续内层循环
			}
			fmt.Println(j)
			if j == 2 {
				break there // 跳出外层循环
			}
		}
	}
}

func main() {
	var map01 = make(map[string]int)
	map01["Cat"] = 1
	map01["Dog"] = 8
	for k, v := range map01 {
		fmt.Println(k, v)
	}
}
