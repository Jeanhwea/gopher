////////////////////////////////////////////////////////////////////////////////
// 状态模式
////////////////////////////////////////////////////////////////////////////////
package main

import "fmt"

type Thread struct {
	state StateType
}

type StateType int

const (
	Idle StateType = iota
	Running
	Done
)

type ActionType int

const (
	Start ActionType = iota
	Terminate
)

func NewThread() *Thread {
	return &Thread{state: Idle}
}

func (this *Thread) NextState(action ActionType) {
	switch this.state {
	case Idle:
		switch action {
		case Start:
			this.state = Running
		case Terminate:
			this.state = Done
		}
	case Running:
		switch action {
		case Terminate:
			this.state = Done
		}
	}
}

func (this *Thread) ShowStatue() {
	stateMap := map[StateType]string{
		Idle:    "Idle",
		Running: "Running",
		Done:    "Done",
	}
	fmt.Printf("Current state: %s\n", stateMap[this.state])
}

func main() {
	thread := NewThread()
	thread.ShowStatue()

	thread.NextState(Start)
	thread.ShowStatue()
}
