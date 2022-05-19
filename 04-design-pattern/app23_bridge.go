////////////////////////////////////////////////////////////////////////////////
// 桥接模式
////////////////////////////////////////////////////////////////////////////////
package main

// 桥接模式分离抽象部分和实现部分。使得两部分独立扩展。
// 桥接模式类似于策略模式，区别在于策略模式封装一系列算法使得算法可以互相替换。
// 策略模式使抽象部分和实现部分分离，可以独立变化。

import ("fmt")

type AbstractMessage interface {
	SendMessage(text, to string)
}

type MessageImplementer interface {
	Send(text, to string)
}

type MessageSMS struct{}

func ViaSMS() MessageImplementer {
	return &MessageSMS{}
}

func (*MessageSMS) Send(text, to string) {
	fmt.Printf("send %s to %s via SMS\n", text, to)
}

type MessageEmail struct{}

func ViaEmail() MessageImplementer {
	return &MessageEmail{}
}

func (*MessageEmail) Send(text, to string) {
	fmt.Printf("send %s to %s via Email\n", text, to)
}

type CommonMessage struct {
	method MessageImplementer
}

func NewCommonMessage(method MessageImplementer) *CommonMessage {
	return &CommonMessage{
		method: method,
	}
}

func (this *CommonMessage) SendMessage(text, to string) {
	this.method.Send(text, to)
}

// 紧急消息可以是: 1.短信, 2. Email
type UrgencyMessage struct {
	method MessageImplementer
}

func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage {
	return &UrgencyMessage{
		method: method,
	}
}

func (this *UrgencyMessage) SendMessage(text, to string) {
	this.method.Send(fmt.Sprintf("[Urgency] %s\n", text), to)
}

func main() {
	var m AbstractMessage
	m = NewCommonMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")

	m = NewUrgencyMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")

	m = NewUrgencyMessage(ViaEmail())
	m.SendMessage("have a drink?", "bob")
}
