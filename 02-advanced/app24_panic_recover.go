package main

import (
	"fmt"
	"runtime/debug"
)

////////////////////////////////////////////////////////////////////////////////
// 使用 panic

func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("normally returned")
}

func test01() {
	firstName := "Elon"
	fullName(&firstName, nil)
}

////////////////////////////////////////////////////////////////////////////////
// 数组下标越界

func slicePanic() {
	n := []int{2, 3, 2}
	fmt.Println(n[4])
	fmt.Println("normally returned")
}

func test02() {
	slicePanic()
}

////////////////////////////////////////////////////////////////////////////////
// defer 在 panic 发生后仍然调用

func slicePanic2() {
	defer fmt.Println("slicePanic2 defered, defer will called!")
	n := []int{2, 3, 2}
	fmt.Println(n[4])
	fmt.Println("normally returned")
}

////////////////////////////////////////////////////////////////////////////////
// recover, 使用

func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Printf("recovered from [%s]", r)
		debug.PrintStack()
	}
}

func fullName2(firstName *string, lastName *string) {
	defer recoverFullName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("normally returned")
}

func test04() {
	lastName := "Elon"
	fullName2(nil, &lastName)
}

func test05() {
	defer func() {
		recover()
	}()
	panic(nil)
}

func SafeRun() {
	defer func() {
		if r := recover(); r != nil {
			stack := debug.Stack()
			fmt.Printf("message = [%v], stack = %s", r, stack)
		}
	}()
	panic("I'm paniced!")
}

func main() {
	// test01()
	// slicePanic()
	// slicePanic2()
	// test04()
	SafeRun()
	fmt.Println("end of main")
}
