package main

import ("fmt")

func removeDuplicates(nums []int) int {
	k := -1
	for i := 0; i < len(nums); i++ {
		if k >= 1 && nums[i] == nums[k] && nums[i] == nums[k-1] {
			continue
		}
		k++
		nums[k] = nums[i]
	}
	return k + 1
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 3}
	n := removeDuplicates(nums)
	fmt.Println(n)
}
