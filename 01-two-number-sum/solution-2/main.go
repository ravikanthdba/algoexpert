package main

import "fmt"

// TwoNumberSum - takes an array and target sum, and returns the pair of integers which return the sum
func TwoNumberSum(array []int, target int) []int {
	var hashMap = make(map[int]bool)
	for _, value := range array {
		if !hashMap[value] {
			hashMap[target-value] = true
		} else {
			return []int{value, target - value}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10))
	fmt.Println(TwoNumberSum([]int{-21, 101, 12, 4, 65, 56, 210, 356, 9, -47}, 164))
	fmt.Println(TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 15))
}
