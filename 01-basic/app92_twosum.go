package main

import (
	"fmt"
)

func twosum(arr []int, target int) [2]int {
	var cache = make(map[int]int)
	for i := 0; i < len(arr); i++ {
		j, ok := cache[target-arr[i]]
		if ok {
			ans := [2]int{j, i}
			return ans
		} else {
			cache[arr[i]] = i
		}
	}
	return [2]int{-1, -1}
}

func main() {
	var arr = []int{1, 3, 3, 4, 8, 9}
	ans := twosum(arr, 14)
	fmt.Println(ans)
}
