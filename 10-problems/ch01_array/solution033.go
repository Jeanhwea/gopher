package main

import "fmt"

func search(nums []int, target int) int {
	beg, end := 0, len(nums)-1
	// [beg, end]
	for beg <= end {
		// 1. 将数组一分为二
		mid := beg + (end-beg)/2
		if nums[mid] == target {
			return mid
		}
		// 2. 二分后的两个区间的左右区间必定有一个是有序的
		if nums[beg] <= nums[mid] {
			if nums[beg] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				beg = mid + 1
			}
		} else if nums[beg] > nums[mid] {
			if nums[mid] < target && target <= nums[end] {
				beg = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 5
	// nums := []int{3,1}
	// target := 1
	k := search(nums, target)
	fmt.Println(k)
}
