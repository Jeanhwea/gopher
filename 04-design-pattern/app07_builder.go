////////////////////////////////////////////////////////////////////////////////
// 创建者模式
////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"strconv"
)

type EmployeeBuilder interface {
	Code(code string)
	Name(name string)
	Age(age int)
}

type Manager struct {
	builder EmployeeBuilder
}

func NewManager(builder EmployeeBuilder) *Manager {
	return &Manager{builder: builder}
}

////////////////////////////////////////////////////////////////////////////////
func (this *Manager) Construct(code, name string, age int) {
	this.builder.Code(code)
	this.builder.Name(name)
	this.builder.Age(age)
}

type StringBuilder struct {
	Data string
}

func (this *StringBuilder) Code(code string) {
	this.Data += code
}

func (this *StringBuilder) Name(name string) {
	this.Data += name
}

func (this *StringBuilder) Age(age int) {
	this.Data += strconv.Itoa(age)
}

type ByteBuilder struct {
	Data []byte
}

func (this *ByteBuilder) Code(code string) {
	this.Data = append(this.Data, []byte(code)...)
}

func (this *ByteBuilder) Name(name string) {
	this.Data = append(this.Data, []byte(name)...)
}

func (this *ByteBuilder) Age(age int) {
	this.Data = append(this.Data, byte(age))
}

func main() {
	stringBuilder := StringBuilder{}
	manager1 := NewManager(&stringBuilder)
	manager1.Construct("001", "Liu", 23)
	fmt.Println(manager1.builder)

	byteBuilder := ByteBuilder{}
	manager2 := NewManager(&byteBuilder)
	manager2.Construct("001", "Liu", 23)
	fmt.Println(manager2.builder)
}
