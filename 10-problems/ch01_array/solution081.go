package main

import (
	"fmt"
)

func search(nums []int, target int) bool {
	beg, end := 0, len(nums)-1
	// [beg, end]
	for beg <= end {
		mid := beg + (end-beg)/2
		if nums[mid] == target {
			return true
		}
		if nums[beg] < nums[mid] {
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
		} else {
			beg++
		}
	}
	return false
}

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 2
	k := search(nums, target)
	fmt.Println(k)
}
