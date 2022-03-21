////////////////////////////////////////////////////////////////////////////////
// 中介者模式
////////////////////////////////////////////////////////////////////////////////
package main

import (
	"log"
	"strings"
)

// 中介者模式封装对象之间互交，使依赖变的简单，并且使复杂互交简单化，封装在中介者中。
var mediator *Mediator

type Mediator struct {
	CD    *CDDriver
	CPU   *CPU
	Video *VideoCard
	Sound *SoundCard
}

func GetMediatorInstance() *Mediator {
	if mediator == nil {
		mediator = &Mediator{}
	}
	return mediator
}

func (this *Mediator) changed(o interface{}) {
	switch obj := o.(type) {
	case *CDDriver:
		this.CPU.Process(obj.Data)
	case *CPU:
		this.Sound.Play(obj.Sound)
		this.Video.Display(obj.Video)
	}
}

////////////////////////////////////////////////////////////////////////////////
type CDDriver struct {
	Data string
}

func (this *CDDriver) ReadData() {
	this.Data = "music,image"
	log.Printf("CDDriver: reading data %s\n", this.Data)
	GetMediatorInstance().changed(this)
}

type CPU struct {
	Video string
	Sound string
}

func (this *CPU) Process(data string) {
	sp := strings.Split(data, ",")
	this.Sound = sp[0]
	this.Video = sp[1]
	log.Printf("CPU: split data with Sound %s, Vides %s\n", this.Sound, this.Video)
	GetMediatorInstance().changed(this)
}

type VideoCard struct {
	Data string
}

func (this *VideoCard) Display(data string) {
	this.Data = data
	log.Printf("VideoCard: display %s\n", this.Data)
	GetMediatorInstance().changed(this)
}

type SoundCard struct {
	Data string
}

func (this *SoundCard) Play(data string) {
	this.Data = data
	log.Printf("SoundCard: play %s\n", this.Data)
	GetMediatorInstance().changed(this)
}

func main() {
	mediator := GetMediatorInstance()
	mediator.CD = &CDDriver{}
	mediator.CPU = &CPU{}
	mediator.Video = &VideoCard{}
	mediator.Sound = &SoundCard{}

	mediator.CD.ReadData()
}
