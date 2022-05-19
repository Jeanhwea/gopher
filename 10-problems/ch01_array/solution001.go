package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	opposite := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := opposite[nums[i]]; ok {
			return []int{v, i}
		} else {
			opposite[target-nums[i]] = i
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
