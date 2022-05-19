package main

import (
	"fmt"
	"time"
)

// 带有 default 的基本类型 switch

func basicSwitch() {
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Today is Saturday.")
	case time.Sunday:
		fmt.Println("Today is Sunday.")
	default:
		fmt.Println("Today is weekday.")
	}
}

// 没有条件的 switch

func noConditionSwitch() {
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

// case 列表的 switch

func whiteSpace(c rune) bool {
	switch c {
	case ' ', '\t', '\n', '\f', '\r':
		return true
	}
	return false
}

// 带 fallthrough 的 switch

func ftSwitch(in int) {
	switch in {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough // 失败后继续执行下一个 case
	case 3:
		fmt.Println("3")
	}
}

// 带有退出循环的 switch

func exitWithBreak() {
Loop:
	for _, ch := range "a b\nc" {
		switch ch {
		case ' ':
			break
		case '\n':
			break Loop
		default:
			fmt.Printf("%c\n", ch)
		}
	}
}

// 执行顺序控制

func executionOrder() {
	foo := func(n int) int {
		fmt.Println(n)
		return n
	}

	switch foo(2) {
	case foo(1), foo(2), foo(3):
		fmt.Println("First case")
		fallthrough
	case foo(4):
		fmt.Println("Second case")
	}

}

func main() {
	basicSwitch()
	noConditionSwitch()
	ftSwitch(2)
	exitWithBreak()
	executionOrder()
}
