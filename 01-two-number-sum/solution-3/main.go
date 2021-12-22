package main

import (
	"fmt"
	"sort"
)

// TwoNumberSum - takes an array and target sum, and returns the pair of integers which return the sum
func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array)
	fmt.Println(array)
	left := 0
	right := len(array) - 1
	for left < right {
		if array[left]+array[right] == target {
			return []int{array[left], array[right]}
		} else if array[left]+array[right] < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}

func main() {
	fmt.Println(TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10))
	fmt.Println(TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 13))
	fmt.Println(TwoNumberSum([]int{-21, 101, 12, 4, 65, 56, 210, 356, 9, -47}, 164))
	fmt.Println(TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 15))
}
