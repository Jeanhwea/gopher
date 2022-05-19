package main

import (
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// 方法定义
type Employee struct {
	name     string
	salary   int
	currency string
}

// 定义方法, value receiver, 对 e 的修改不影响原始对象, 每次复制值会有性能问题
func (e Employee) displaySalary() {
	fmt.Printf("%s's Salary is %s%d\n", e.name, e.currency, e.salary)
}

// 定义方法, pointer receiver, 对 e 的修改影响原始对象
func (e *Employee) displaySalary2() {
	fmt.Printf("%s's Salary is %s%d\n", e.name, e.currency, e.salary)
}

// 定义函数, 普通函数没有 receiver, 需要显式传递参数
func displaySalaryFunc(e Employee) {
	fmt.Printf("%s's Salary is %s%d\n", e.name, e.currency, e.salary)
}

func test01() {
	emp1 := Employee{
		name:     "Jack",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary()
	emp1.displaySalary2()
	displaySalaryFunc(emp1)
}

////////////////////////////////////////////////////////////////////////////////
// 不带结构体的 receiver 定义
type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	num1, num2 := myInt(5), myInt(3)
	sum := num1.add(num2)
	fmt.Println(sum)
}
