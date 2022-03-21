package main

import "fmt"

func test01() {
	var g = func() { fmt.Println("g1") }
	defer g() // 延迟调用的函数值的估值时刻
	g = func() { fmt.Println("g2") }
}

func test02() {
	var g = func() { fmt.Println("g1") }
	var f = func() { g() }
	defer f() // 延迟调用的函数值的估值时刻
	g = func() { fmt.Println("g2") }
}

////////////////////////////////////////////////////////////////////////////////
type T int

func (t T) M(n int) T {
	fmt.Print(n)
	return t
}

// test03 将打印 1342
func test03() {
	var t T
	defer t.M(1).M(2)
	t.M(3).M(4)
}

func main() {
	test01()
	test02()
	test03()
}
