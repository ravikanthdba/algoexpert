package main

import (
	"fmt"
)

type slice struct {
	array []int
}

func (s *slice) position(n int) int {
	for key := range s.array {
		if s.array[key] == n {
			return key
		}
	}
	return -1
}

func ThreeNumberSort(array []int, order []int) []int {
	if len(array) == 0 || len(order) == 0 {
		return []int{}
	}
	var firstIndex, secondIndex int
	if array[0] == order[0] {
		firstIndex = 1
		secondIndex = 1
	} else {
		firstIndex = 0
		secondIndex = 0
	}
	thirdIndex := len(array) - 1
	sliceOrder := slice{
		array: order,
	}

	for secondIndex <= thirdIndex {
		switch {
		case sliceOrder.position(array[secondIndex]) == 0:
			array[firstIndex], array[secondIndex] = array[secondIndex], array[firstIndex]
			secondIndex++
			firstIndex++
		case sliceOrder.position(array[secondIndex]) == 1:
			if sliceOrder.position(array[firstIndex]) > sliceOrder.position(array[secondIndex]) {
				array[firstIndex], array[secondIndex] = array[secondIndex], array[firstIndex]
				firstIndex++
			}
			secondIndex++
		case sliceOrder.position(array[secondIndex]) == 2:
			array[thirdIndex], array[secondIndex] = array[secondIndex], array[thirdIndex]
			thirdIndex--
		default:
			secondIndex++
		}
	}
	return array
}

func main() {
	fmt.Println(ThreeNumberSort([]int{1, 3, 4, 4, 4, 4, 3, 3, 3, 4, 1, 1, 1, 4, 4, 1, 3, 1, 4, 4}, []int{1, 4, 3}))
}
