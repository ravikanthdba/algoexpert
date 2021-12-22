package main

import (
	"fmt"
	"sort"
)

//ThreeNumberSum - takes array of integers and gives the set of integers where sum == target
func ThreeNumberSum(array []int, target int) [][]int {
	var finalArray [][]int
	sort.Ints(array)
	for i := 0; i < len(array); i++ {
		left := i + 1
		right := len(array) - 1
		for left < right {
			if array[i]+array[left]+array[right] < target {
				left++
			} else if array[i]+array[left]+array[right] > target {
				right--
			} else if array[i]+array[left]+array[right] == target {
				var simpleArray []int
				simpleArray = append(simpleArray, array[i], array[left], array[right])
				finalArray = append(finalArray, simpleArray)
				left++
				right--
			}
		}
	}
	if len(finalArray) == 0 {
		return [][]int{}
	}
	return finalArray
}

func main() {
	// fmt.Println(ThreeNumberSum([]int{12, 3, 1, 2, -6, 5, -8, 6}, 0))
	// fmt.Println(ThreeNumberSum([]int{1, 2, 3}, 6))
	fmt.Println(ThreeNumberSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}, 32))
}
