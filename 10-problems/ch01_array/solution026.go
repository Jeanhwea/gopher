package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	k := -1
	for i := 0; i < len(nums); i++ {
		if k >= 0 && nums[i] == nums[k] {
			continue
		}
		k++
		nums[k] = nums[i]
	}
	return k + 1
}

func main() {
	nums := []int{1, 2, 2, 3}
	n := removeDuplicates(nums)
	fmt.Println(n)
}
