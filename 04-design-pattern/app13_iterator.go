////////////////////////////////////////////////////////////////////////////////
// 迭代器模式
////////////////////////////////////////////////////////////////////////////////
package main

import "fmt"

type Aggregate interface {
	Iterator() Iterator
}

// 迭代器实现的三个方法
type Iterator interface {
	First()
	hasNext() bool
	Next() interface{}
}

////////////////////////////////////////////////////////////////////////////////
type Sequence struct {
	start, end int
}

func NewSequence(start, end int) *Sequence {
	return &Sequence{start: start, end: end}
}

func (this *Sequence) Iterator() Iterator {
	return &SequenceIterator{seq: this, next: this.start}
}

type SequenceIterator struct {
	seq  *Sequence
	next int
}

func (this *SequenceIterator) First() {
	this.next = this.seq.start
}

func (this *SequenceIterator) hasNext() bool {
	return this.next <= this.seq.end
}

func (this *SequenceIterator) Next() interface{} {
	if this.hasNext() {
		next := this.next
		this.next++
		return next
	}
	return nil
}

func main() {
	iter1 := NewSequence(1, 8).Iterator()
	for iter1.First(); iter1.hasNext(); {
		fmt.Println(iter1.Next())
	}
}
