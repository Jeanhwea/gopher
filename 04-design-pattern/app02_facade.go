////////////////////////////////////////////////////////////////////////////////
// 外观模式
////////////////////////////////////////////////////////////////////////////////
package main

import "fmt"

// 外观的接口
type API interface {
	Test() string
}

// API 为 facade 模块的外观接口，大部分代码使用此接口简化对 facade 类的访问
func NewAPI() API {
	return &apiImpl{a: NewAModuleAPI(), b: NewBModuleAPI()}
}

// 外观模式实现类
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (this *apiImpl) Test() string {
	aRet := this.a.TestA()
	bRet := this.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

////////////////////////////////////////////////////////////////////////////////
// A Module 的实现
type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct{}

func (*aModuleImpl) TestA() string {
	return "A module TestA is running"
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

////////////////////////////////////////////////////////////////////////////////
// B Module 的实现
type BModuleAPI interface {
	TestB() string
}

type bModuleImpl struct{}

func (*bModuleImpl) TestB() string {
	return "B module TestB is running"
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

func main() {
	// API 在使用是只需要创建
	api := NewAPI()

	// 不需要关心 A Module 和 B Module 的实现细节
	ret := api.Test()
	fmt.Println(ret)
}
