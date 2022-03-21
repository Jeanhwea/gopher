////////////////////////////////////////////////////////////////////////////////
// 解释器模式
////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 解释器模式定义一套语言文法，并设计该语言解释器，使用户能使用特定文法控制解释器行为。
// 解释器模式的意义在于，它分离多种复杂功能的实现，每个功能只需关注自身的解释。
// 对于调用者不用关心内部的解释器的工作，只需要用简单的方式组合命令就可以。

type Node interface {
	Interpret() int
}

type ValNode struct {
	val int
}

func (n *ValNode) Interpret() int {
	return n.val
}

type AddNode struct {
	left, right Node
}

func (this *AddNode) Interpret() int {
	return this.left.Interpret() + this.right.Interpret()
}

type MinNode struct {
	left, right Node
}

func (this *MinNode) Interpret() int {
	return this.left.Interpret() - this.right.Interpret()
}

type Parser struct {
	exp   []string
	index int
	prev  Node
}

func (this *Parser) Parse(exp string) {
	this.exp = strings.Split(exp, " ")

	for {
		if this.index >= len(this.exp) {
			return
		}
		switch this.exp[this.index] {
		case "+":
			this.prev = this.newAddNode()
		case "-":
			this.prev = this.newMinNode()
		default:
			this.prev = this.newValNode()
		}
	}
}

func (this *Parser) newAddNode() Node {
	this.index++
	return &AddNode{
		left:  this.prev,
		right: this.newValNode(),
	}
}

func (this *Parser) newMinNode() Node {
	this.index++
	return &MinNode{
		left:  this.prev,
		right: this.newValNode(),
	}
}

func (this *Parser) newValNode() Node {
	v, _ := strconv.Atoi(this.exp[this.index])
	this.index++
	return &ValNode{
		val: v,
	}
}

func (this *Parser) Result() Node {
	return this.prev
}

func main() {
	p := &Parser{}
	p.Parse("1 + 2 + 3 - 4 + 5 - 6")
	res := p.Result().Interpret()
	fmt.Println(res)
}
