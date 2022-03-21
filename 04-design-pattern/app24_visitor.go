////////////////////////////////////////////////////////////////////////////////
// 访问者模式
////////////////////////////////////////////////////////////////////////////////
package main

// 访问者模式可以给一系列对象透明的添加功能，并且把相关代码封装到一个类中。
// 对象只要预留访问者接口 Accept 则后期为对象添加功能的时候就不需要改动对象。

import "fmt"

type Customer interface {
	Accept(Visitor)
}

type Visitor interface {
	Visit(Customer)
}

type EnterpriseCustomer struct {
	name string
}

// 将所有的客户都记录到集合中
type CustomerCol struct {
	customers []Customer
}

func (this *CustomerCol) Add(customer Customer) {
	this.customers = append(this.customers, customer)
}

func (this *CustomerCol) Accept(visitor Visitor) {
	for _, customer := range this.customers {
		customer.Accept(visitor)
	}
}

func NewEnterpriseCustomer(name string) *EnterpriseCustomer {
	return &EnterpriseCustomer{
		name: name,
	}
}

func (this *EnterpriseCustomer) Accept(visitor Visitor) {
	visitor.Visit(this)
}

type IndividualCustomer struct {
	name string
}

func NewIndividualCustomer(name string) *IndividualCustomer {
	return &IndividualCustomer{
		name: name,
	}
}

func (this *IndividualCustomer) Accept(visitor Visitor) {
	visitor.Visit(this)
}

type ServiceRequestVisitor struct{}

func (*ServiceRequestVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("serving enterprise customer %s\n", c.name)
	case *IndividualCustomer:
		fmt.Printf("serving individual customer %s\n", c.name)
	}
}

// 只有企业用户才需要做分析
type AnalysisVisitor struct{}

func (*AnalysisVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("analysis enterprise customer %s\n", c.name)
	}
}

func main() {
	c1 := &CustomerCol{}
	c1.Add(NewEnterpriseCustomer("A company"))
	c1.Add(NewEnterpriseCustomer("B company"))
	c1.Add(NewIndividualCustomer("bob"))
	c1.Accept(&ServiceRequestVisitor{})

	c2 := &CustomerCol{}
	c2.Add(NewEnterpriseCustomer("A company"))
	c2.Add(NewIndividualCustomer("bob"))
	c2.Add(NewEnterpriseCustomer("B company"))
	c2.Accept(&AnalysisVisitor{})
}
