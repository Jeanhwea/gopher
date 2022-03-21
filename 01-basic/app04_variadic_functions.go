package main

import "fmt"

func findElement(num int, nums ...int) {
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Printf("find %d at nums[%d].\n", num, i)
			found = true
		}
	}
	if !found {
		fmt.Printf("Can not find %d.\n", num)
	}
}

func findElement2(num int, nums []int) {
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Printf("find %d at nums[%d].\n", num, i)
			found = true
		}
	}
	if !found {
		fmt.Printf("Can not find %d.\n", num)
	}
}

func main() {
	nums := []int{2, 3, 8, 32, 22}
	findElement(8, 3, 2, 3, 23, 8)
	findElement(8, nums...)
	findElement2(8, nums)
}
