package main

import (
	"fmt"
)

// 接口的定义方法
type VowelsFinder interface {
	Find() []rune
}

////////////////////////////////////////////////////////////////////////////////
// 接口的内部表示形式
type Worker interface {
	Work()
}
type Person struct {
	name string
}

func (p *Person) Work() {
	fmt.Println(p.name, "is working.")
}

// 接口的内部表示形式 (type, value).type

func describe(obj interface{}) {
	// interface{} 表示空接口
	fmt.Printf("Interface type %T value %v\n", obj, obj)
}

// 类型断言 interface.(type)
// interface{} 是万能类型

func display(arg interface{}) {
	// 类型断言
	value, ok := arg.(string)
	if ok {
		fmt.Printf("value =\"%s\"\n", value)
	} else {
		fmt.Println("arg =", arg)
	}
}

// 类型 switch

func findType(obj interface{}) {
	switch obj.(type) {
	case string:
		fmt.Printf("I am a string, value = %s\n", obj.(string))
	case int:
		fmt.Printf("I am a int, value = %d\n", obj.(int))
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	naveen := Person{name: "Naveen"}
	var w Worker = &naveen
	describe(w)
	w.Work()
	findType("Naveen")
	findType(99)
	findType(3.4)
}
