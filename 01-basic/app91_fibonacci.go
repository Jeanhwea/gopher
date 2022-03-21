package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return 0
	} else if n <= 2 {
		return 1
	}

	var a, b = 1, 0
	for i := 2; i < n; i++ {
		a, b = a+b, a
	}
	return a
}

func main() {
	for i := 1; i < 20; i++ {
		fmt.Println(fib(i))
	}
}
