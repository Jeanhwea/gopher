package main

import "fmt"

type Man struct {
	name string
}

func (this *Man) GetName() string {
	return this.name
}

// golang 推荐使用 NewT(...) 方法代替构造器
func NewMan(name string) Man {
	return Man{name: name}
}

type SuperMan struct {
	Man   // 继承
	level int
}

func (this *SuperMan) ShowLevel() {
	fmt.Println(this.name, "'s level is", this.level)
}

func main() {
	var wang = &Man{"王大锤"}
	fmt.Println(wang.GetName())

	var batman = &SuperMan{Man{"蝙蝠侠"}, 10}
	batman.ShowLevel()

	// New 构造方法
	var liu = NewMan("刘晓伟")
	fmt.Println(liu)
}
