////////////////////////////////////////////////////////////////////////////////
// 原型模式
////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
)

// Cloneable 是原型对象需要实现的接口
type Cloneable interface {
	Clone() Cloneable
}

// 使用一个 Map 记录所有的对象原型
type PrototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{prototypes: make(map[string]Cloneable)}
}

func (this *PrototypeManager) Get(name string) Cloneable {
	return this.prototypes[name].Clone()
}

func (this *PrototypeManager) Set(name string, prototype Cloneable) {
	this.prototypes[name] = prototype
}

////////////////////////////////////////////////////////////////////////////////
type Type1 struct {
	name string
}

func (this *Type1) Clone() Cloneable {
	return this
}

type Type2 struct {
	name string
}

func (this *Type2) Clone() Cloneable {
	return this
}

func main() {
	var manager *PrototypeManager = NewPrototypeManager()
	manager.Set("t1", &Type1{name: "Type1"})
	manager.Set("t2", &Type1{name: "Type2"})
	t1 := manager.Get("t1")
	t2 := t1.Clone()
	fmt.Println(t1 == t2) // => true
	t3 := t2.(*Type1)
	fmt.Println(t3.name) // => Type1
}
