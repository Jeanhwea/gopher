////////////////////////////////////////////////////////////////////////////////
// 工厂方法模式
////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
)

// 被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// 工厂接口
type OperatorFactory interface {
	Create() Operator
}

////////////////////////////////////////////////////////////////////////////////
type OperatorBase struct {
	a, b int
}

func (this *OperatorBase) SetA(a int) {
	this.a = a
}

func (this *OperatorBase) SetB(b int) {
	this.b = b
}

type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type PlusOperator struct {
	*OperatorBase
}

func (this *PlusOperator) Result() int {
	return this.a + this.b
}

type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type MinusOperator struct {
	*OperatorBase
}

func (this *MinusOperator) Result() int {
	return this.a - this.b
}

func main() {
	plusFactory := PlusOperatorFactory{}
	plusOp := plusFactory.Create()
	plusOp.SetA(4)
	plusOp.SetB(1)
	fmt.Println(plusOp.Result())
	minusFactory := MinusOperatorFactory{}
	minusOp := minusFactory.Create()
	minusOp.SetA(4)
	minusOp.SetB(1)
	fmt.Println(minusOp.Result())
}
