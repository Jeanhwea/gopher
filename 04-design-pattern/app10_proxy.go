////////////////////////////////////////////////////////////////////////////////
// 代理模式
////////////////////////////////////////////////////////////////////////////////
package main

import ("fmt")

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (this *RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (this *Proxy) Do() string {
	var result string

	// 代理前置处理
	result += "pre:"

	// 实际代理调用
	result += this.real.Do()

	// 代理后置处理
	result += ":post"

	return result
}

func main() {
	var sub Subject
	sub = &Proxy{}

	fmt.Println(sub.Do()) // => pre:real:post
}
