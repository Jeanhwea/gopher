package main

import (
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// value receiver method 和 pointer receiver method 的调用区别
type Describer interface {
	Describe()
}
type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state, country string
}

func (a *Address) Describe() {
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func test01() {
	////////////////////////////////////////////////////////////////////////////////
	// Value receiver 支持指针和赋值两种调用方式
	var d1 Describer
	p1 := Person{"Sam", 23}
	d1 = p1
	d1.Describe()
	p2 := Person{"Jack", 22}
	d1 = &p2
	d1.Describe()
	////////////////////////////////////////////////////////////////////////////////
	// Pointer receiver 仅仅支持指针调用方式
	var d2 Describer
	a := Address{"Washington", "USA"}
	// d2 = a
	// 如果使用上述赋值，编译会报错
	// cannot use a (type Address) as type Describer in assignment
	d2 = &a // 使用地址传值则是正确的
	d2.Describe()
}

////////////////////////////////////////////////////////////////////////////////
// 实现多个接口
type SalaryCalculator interface {
	DisplaySalary()
}
type LeaveCalculator interface {
	CalculateLeavesLeft() int
}
type Employee struct {
	name        string
	basicPay    int
	performance int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s has salary $%d\n", e.name, e.basicPay+e.performance)
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func test02() {
	emp1 := Employee{
		name:        "Jack",
		basicPay:    5000,
		performance: 3000,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var sc SalaryCalculator = emp1
	var lc LeaveCalculator = emp1
	sc.DisplaySalary()
	fmt.Println("Leaves left =", lc.CalculateLeavesLeft())
}

////////////////////////////////////////////////////////////////////////////////
// 嵌入式接口, 可以实现 mixin 的效果
type EmployeeOperation interface {
	SalaryCalculator
	LeaveCalculator
}

func test03() {
	emp2 := Employee{
		name:        "Tom",
		basicPay:    5500,
		performance: 3100,
		totalLeaves: 32,
		leavesTaken: 8,
	}
	var empOp EmployeeOperation = emp2
	empOp.DisplaySalary()
	fmt.Println("Leaves left =", empOp.CalculateLeavesLeft())
}

func main() {
	// test01()
	// test02()
	test03()
}
