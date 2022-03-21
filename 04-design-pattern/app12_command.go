////////////////////////////////////////////////////////////////////////////////
// 命令模式
////////////////////////////////////////////////////////////////////////////////
package main

import "fmt"

// 命令模式本质是把某个对象的方法调用封装到对象中，方便传递、存储、调用
type Command interface {
	Execute()
}

////////////////////////////////////////////////////////////////////////////////

// 定义开机命令
type StartCommand struct {
	mb *MotherBoard
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{mb: mb}
}

func (this *StartCommand) Execute() {
	this.mb.Start()
}

// 定义关机命令
type RebootCommand struct {
	mb *MotherBoard
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{mb: mb}
}

func (this *RebootCommand) Execute() {
	this.mb.Reboot()
}

type MotherBoard struct{}

func (this *MotherBoard) Start() {
	fmt.Println("System starting")
}

func (this *MotherBoard) Reboot() {
	fmt.Println("System rebooting")
}

// 主机对象中只包含对应的按钮, 并不关心按钮所绑定的命令
type Box struct {
	button1 Command
	button2 Command
}

func NewBox(button1, button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}

func (this *Box) PressButton1() {
	this.button1.Execute()
}

func (this *Box) PressButton2() {
	this.button2.Execute()
}

func main() {
	motherBoard := &MotherBoard{}
	btn1 := NewStartCommand(motherBoard)
	btn2 := NewRebootCommand(motherBoard)

	box1 := NewBox(btn1, btn2)
	box1.PressButton1()
	box1.PressButton2()

	// 配置灵活
	box2 := NewBox(btn2, btn1)
	box2.PressButton1()
	box2.PressButton2()
}
