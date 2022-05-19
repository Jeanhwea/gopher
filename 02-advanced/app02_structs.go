package main

import (
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// 结构体定义
type Employee struct {
	firstName string
	lastName  string
	age       int
}

func test01() {
	// 创建结构体
	emp1 := Employee{"Tom", "Liu", 23}
	emp2 := Employee{
		firstName: "Dachui",
		lastName:  "Wang",
		age:       10,
	}
	fmt.Println(emp1)
	fmt.Println(emp2)
	// 匿名结构体
	emp3 := struct {
		name string
		age  int
	}{
		name: "Jack",
		age:  34,
	}
	fmt.Println(emp3)
	// 空结构体
	var emp4 Employee
	fmt.Println(emp4)
	pEmp := &emp1
	fmt.Println(pEmp)
	(*pEmp).age = 99
	fmt.Println(emp1)
	pEmp.age = 99 // golang 会推导变量类型, 可以省略 dereference
	fmt.Println(emp1)
}

////////////////////////////////////////////////////////////////////////////////
// 匿名域的结构体
type Manager struct {
	string
	int
}

func test02() {
	mgr1 := Manager{
		string: "Liu",
		int:    9,
	}
	fmt.Println(mgr1)
}

////////////////////////////////////////////////////////////////////////////////
// 嵌套结构体
type Address struct {
	city  string
	state string
}
type Person struct {
	name string
	addr Address
}

func test03() {
	psn1 := Person{
		name: "James",
		addr: Address{
			city:  "Beijing",
			state: "Chaoyang",
		},
	}
	fmt.Println(psn1)
}

////////////////////////////////////////////////////////////////////////////////
// 域提升
type Address02 struct {
	city, state string
}
type Person02 struct {
	name string
	age  int
	Address02
}

func test04() {
	psn2 := Person02{
		name: "James",
		age:  33,
		Address02: Address02{
			city:  "Beijing",
			state: "Chaoyang",
		},
	}
	fmt.Println(psn2)
	fmt.Println(psn2.name)
	fmt.Println(psn2.city) // Address02 结构体中的 city 变量提升到了 Person02 结构体中
}

////////////////////////////////////////////////////////////////////////////////
// 结构体等值比较
type name struct {
	firstName, lastName string
}

func test05() {
	name1 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	name2 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	name3 := name{
		firstName: "Steve",
	}
	// 如果结构体每个域都可以比较，则等值比较会逐个比较
	fmt.Println(name1 == name2) // => true
	fmt.Println(name1 == name3) // => false
	// 如果结构体中包含 slice 或 map 则无法比较
}

func main() {
	test05()
}
