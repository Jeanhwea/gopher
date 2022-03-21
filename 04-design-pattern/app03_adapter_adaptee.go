////////////////////////////////////////////////////////////////////////////////
// 适配器模式
////////////////////////////////////////////////////////////////////////////////
package main

import "fmt"

// 适配器模式用于转换一种接口适配另一种接口, Target 是适配目标, 适配目标的 Request 方法
type Target interface {
	Request() string
}

type Adaptee interface {
	SpecialRequest() string
}

////////////////////////////////////////////////////////////////////////////////
type adapteeImpl struct{}

func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

func (this *adapteeImpl) SpecialRequest() string {
	return "adaptee method"
}

type adaptor struct {
	Adaptee
}

func NewAdaptor(adaptee Adaptee) Target {
	return &adaptor{
		Adaptee: adaptee,
	}
}

func (this *adaptor) Request() string {
	return this.SpecialRequest()
}

func main() {
	adaptee := NewAdaptee()       // 被适配的对象
	target := NewAdaptor(adaptee) // 适配的目标

	fmt.Println(target.Request()) // 调用目标方法
}
