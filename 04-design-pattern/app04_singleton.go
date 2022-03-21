////////////////////////////////////////////////////////////////////////////////
// 单例模式
////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"sync"
)

// 获取单例的方法
func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

// 单例模式的接口，通过该接口导出可以避免 GetInstance 函数返回包私有类型的指针
type Singleton interface {
	foo()
}

////////////////////////////////////////////////////////////////////////////////
// singleton 私有的单例模式类
type singleton struct{}

func (this *singleton) foo() {}

var (
	instance *singleton
	once     sync.Once
)

func main() {
	ins01 := GetInstance()
	ins02 := GetInstance()
	fmt.Println(ins01 == ins02)
}
