////////////////////////////////////////////////////////////////////////////////
// 装饰模式
////////////////////////////////////////////////////////////////////////////////
package main

import "fmt"

// 装饰模式使用对象组合的方式动态改变或增加对象行为
type Component interface {
	Calc() int
}

type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	return 0
}

type MulDecorator struct {
	Component
	num int
}

func WarpMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

func (this *MulDecorator) Calc() int {
	return this.Component.Calc() * this.num
}

type AddDecorator struct {
	Component
	num int
}

func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

func (this *AddDecorator) Calc() int {
	return this.Component.Calc() + this.num
}

func main() {
	var c Component = &ConcreteComponent{}
	c = WarpAddDecorator(c, 10)
	c = WarpMulDecorator(c, 8)
	res := c.Calc()

	fmt.Printf("res %d\n", res)
}
