package main

import (
	"fmt"
	"runtime/debug"
	"strconv"
	"time"
)

func test02() {
	test01()
}

func test01() {
	stack := debug.Stack()
	fmt.Printf("stack = %s\n", string(stack))
}

func main() {
	str01 := strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Println(str01)

	str21 := ""
	fmt.Println(str21 < "1")
	fmt.Println("2" < "1")
	fmt.Println("0" < "1")
	fmt.Println("0" > "")

	// a := [5]int{1, 2, 3}
	// s := a[1:2:3]
	// s[0] = 9
	// fmt.Println(a)
	// test02()
}
