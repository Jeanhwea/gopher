package main

import (
	"fmt"
)

// 接口
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

// 实现类 Cat
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 实现类 Dog
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println(animal.GetType(), "is", animal.GetColor())
}

func main() {
	var animal AnimalIF
	animal = &Cat{"Green"}
	showAnimal(animal)
	animal = &Dog{"Yellow"}
	showAnimal(animal)
}
