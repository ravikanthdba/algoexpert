// Move all the first elements to first place and last elements to last place, we maintain all middle elements at middle place here

package main

import (
	"fmt"
)

func ThreeNumberSort(array []int, order []int) []int {
	if len(array) == 0 || len(order) == 0 {
		return []int{}
	}

	var firstIndex, lastIndex int

	if array[0] == order[0] {
		firstIndex = 1
	} else {
		firstIndex = 0
	}
	lastIndex = len(array) - 1
	
	for i := 1; i < len(array); i ++ {
		if array[i] == order[0] {
			array[firstIndex], array[i] = array[i], array[firstIndex]
			firstIndex++
		} else {
			continue
		}
	}


	for i := len(array)-1; i >= 0; i--{
		if array[i] == order[len(order) - 1] {
			array[lastIndex], array[i] = array[i], array[lastIndex]
			lastIndex--
		} else {
			continue
		}
	}
	return array
}

func main() {
	fmt.Println(ThreeNumberSort([]int{1,0,0,-1,-1,0,1,1}, []int{0,1,-1}))
	fmt.Println(ThreeNumberSort([]int{1, 3, 4, 4, 4, 4, 3, 3, 3, 4, 1, 1, 1, 4, 4, 1, 3, 1, 4, 4}, []int{1,4,3}))
}