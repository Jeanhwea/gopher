////////////////////////////////////////////////////////////////////////////////
// 备忘录模式
////////////////////////////////////////////////////////////////////////////////
package main

import ("fmt")

// 备忘录模式用于保存程序内部状态到外部，又不希望暴露内部状态的情形
type Memento interface{}

type Game struct {
	hp, mp int
}

type gameMemento struct {
	hp, mp int
}

func (this *Game) Play(mpDelta, hpDelta int) {
	this.mp += mpDelta
	this.hp += hpDelta
}

func (this *Game) Save() Memento {
	return &gameMemento{
		hp: this.hp,
		mp: this.mp,
	}
}

func (this *Game) Load(m Memento) {
	gm := m.(*gameMemento)
	this.mp = gm.mp
	this.hp = gm.hp
}

func (this *Game) Status() {
	fmt.Printf("Current HP:%d, MP:%d\n", this.hp, this.mp)
}

func main() {
	game := &Game{
		hp: 10,
		mp: 10,
	}

	game.Status()
	progress := game.Save()

	game.Play(-2, -3)
	game.Status()

	game.Load(progress)
	game.Status()
}

////////////////////////////////////////////////////////////////////////////////
// Current HP:10, MP:10
// Current HP:7, MP:8
// Current HP:10, MP:10
