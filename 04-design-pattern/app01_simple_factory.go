////////////////////////////////////////////////////////////////////////////////
// 简单工厂模式
////////////////////////////////////////////////////////////////////////////////
package main

import "fmt"

// 包中只有 API 接口和 NewAPI 函数为包外可见，封装了实现细节
type API interface {
	Say(name string) string
}

const (
	API_TYPE_HI = iota
	API_TYPE_HELLO
)

// go 语言没有构造函数一说，所以一般会定义 NewXXX 函数来初始化相关类
// NewXXX 函数返回接口时就是简单工厂模式，也就是说 Golang 的一般推荐做法就是简单工厂
func NewAPI(whichType int) API {
	if whichType == API_TYPE_HI {
		return &hiAPI{}
	} else if whichType == API_TYPE_HELLO {
		return &helloAPI{}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
type hiAPI struct{}

func (this *hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

type helloAPI struct{}

func (this *helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func main() {
	api1 := NewAPI(API_TYPE_HI)
	fmt.Println(api1.Say("Wang"))

	api2 := NewAPI(API_TYPE_HELLO)
	fmt.Println(api2.Say("Liu"))
}
