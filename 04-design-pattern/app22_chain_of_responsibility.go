////////////////////////////////////////////////////////////////////////////////
// 职责链模式
////////////////////////////////////////////////////////////////////////////////
package main

// 职责链模式用于分离不同职责，并且动态组合相关职责。
//
// Golang 实现职责链模式时候，因为没有继承的支持，使用链对象包涵职责的方式，即：
//
// 1. 链对象包含当前职责对象以及下一个职责链。
// 2. 职责对象提供接口表示是否能处理对应请求。
// 3. 职责对象提供处理函数处理相关职责。
//
// 同时可在职责链类中实现职责接口相关函数，使职责链对象可以当做一般职责对象是用。

import ("fmt")

type Manager interface {
	HaveRight(money int) bool
	HandleFeeRequest(name string, money int) bool
}

type RequestChain struct {
	Manager
	successor *RequestChain
}

func (this *RequestChain) SetSuccessor(m *RequestChain) {
	this.successor = m
}

func (this *RequestChain) HandleFeeRequest(name string, money int) bool {
	if this.Manager.HaveRight(money) {
		return this.Manager.HandleFeeRequest(name, money)
	}
	if this.successor != nil {
		return this.successor.HandleFeeRequest(name, money)
	}
	return false
}

func (this *RequestChain) HaveRight(money int) bool {
	return true
}

type ProjectManager struct{}

func NewProjectManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &ProjectManager{},
	}
}

func (*ProjectManager) HaveRight(money int) bool {
	return money < 500
}

func (*ProjectManager) HandleFeeRequest(name string, money int) bool {
	if name == "bob" {
		fmt.Printf("Project manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("Project manager don't permit %s %d fee request\n", name, money)
	return false
}

type DepManager struct{}

func NewDepManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &DepManager{},
	}
}

func (*DepManager) HaveRight(money int) bool {
	return money < 5000
}

func (*DepManager) HandleFeeRequest(name string, money int) bool {
	if name == "tom" {
		fmt.Printf("Dep manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("Dep manager don't permit %s %d fee request\n", name, money)
	return false
}

type GeneralManager struct{}

func NewGeneralManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &GeneralManager{},
	}
}

func (*GeneralManager) HaveRight(money int) bool {
	return true
}

func (*GeneralManager) HandleFeeRequest(name string, money int) bool {
	if name == "ada" {
		fmt.Printf("General manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("General manager don't permit %s %d fee request\n", name, money)
	return false
}

func main() {
	c1 := NewProjectManagerChain()
	c2 := NewDepManagerChain()
	c3 := NewGeneralManagerChain()

	c1.SetSuccessor(c2)
	c2.SetSuccessor(c3)

	var c Manager = c1

	c.HandleFeeRequest("bob", 400)
	c.HandleFeeRequest("tom", 1400)
	c.HandleFeeRequest("ada", 10000)
	c.HandleFeeRequest("floar", 400)
}
