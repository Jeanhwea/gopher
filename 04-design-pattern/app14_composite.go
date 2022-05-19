////////////////////////////////////////////////////////////////////////////////
// 组合模式
////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
)

// 组合模式统一对象和对象集，使得使用相同接口使用对象和对象集
type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	Print(string)
}

// 组合模式常用于树状结构，用于统一叶子节点和树节点的访问，并且可以用于应用某一操作到所有子节点
const (
	LeafNode = iota
	CompositeNode
)

func NewComponent(kind int, name string) Component {
	var c Component
	switch kind {
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()
	}
	c.SetName(name)
	return c
}

type component struct {
	parent Component
	name   string
}

func (this *component) Parent() Component {
	return this.parent
}

func (this *component) SetParent(parent Component) {
	this.parent = parent
}

func (this *component) Name() string {
	return this.name
}

func (this *component) SetName(name string) {
	this.name = name
}

func (this *component) AddChild(Component) {

}

func (this *component) Print(string) {
}

type Leaf struct {
	component
}

func NewLeaf() *Leaf {
	return &Leaf{}
}

func (this *Leaf) Print(pre string) {
	fmt.Printf("%s-%s\n", pre, this.Name())
}

type Composite struct {
	component
	children []Component
}

func NewComposite() *Composite {
	return &Composite{
		children: make([]Component, 0),
	}
}

func (this *Composite) AddChild(child Component) {
	child.SetParent(this)
	this.children = append(this.children, child)
}

func (this *Composite) Print(pre string) {
	fmt.Printf("%s+%s\n", pre, this.Name())
	pre += " "
	for _, comp := range this.children {
		comp.Print(pre)
	}
}

func main() {
	root := NewComponent(CompositeNode, "root")
	c1 := NewComponent(CompositeNode, "c1")
	c2 := NewComponent(CompositeNode, "c2")
	c3 := NewComponent(CompositeNode, "c3")

	l1 := NewComponent(LeafNode, "l1")
	l2 := NewComponent(LeafNode, "l2")
	l3 := NewComponent(LeafNode, "l3")

	root.AddChild(c1)
	root.AddChild(c2)
	c1.AddChild(c3)
	c1.AddChild(l1)
	c2.AddChild(l2)
	c2.AddChild(l3)

	root.Print("")
}
