package pkg01

import ("fmt")

const (
	PackageName = "pkg01"
)

func init() {
	fmt.Println("init1 called.")
}

func init() {
	fmt.Println("init2 called.")
}

type greetServiceImpl struct{}

func (g *greetServiceImpl) Greeting() {
	fmt.Printf("Hello world from %s.\n", PackageName)
}
