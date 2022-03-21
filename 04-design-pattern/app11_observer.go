////////////////////////////////////////////////////////////////////////////////
// 观察者模式
////////////////////////////////////////////////////////////////////////////////
package main

import "fmt"

type Subject struct {
	observers []Observer
	context   string
}

func NewSubject() *Subject {
	return &Subject{observers: make([]Observer, 0)}
}

func (this *Subject) Subscribe(o Observer) {
	this.observers = append(this.observers, o)
}

func (this *Subject) notify() {
	for _, o := range this.observers {
		o.Update(this)
	}
}

func (this *Subject) UpdateContext(context string) {
	this.context = context
	this.notify()
}

type Observer interface {
	Update(*Subject)
}

type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{name: name}
}

func (this *Reader) Update(sub *Subject) {
	fmt.Printf("%s receive %s\n", this.name, sub.context)
}

func main() {
	subject := NewSubject()
	reader1 := NewReader("reader1")
	reader2 := NewReader("reader2")
	reader3 := NewReader("reader3")
	subject.Subscribe(reader1)
	subject.Subscribe(reader2)
	subject.Subscribe(reader3)

	subject.UpdateContext("Context[observer mode]")
}
