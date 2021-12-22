package main

import "fmt"

type sliceArray struct {
	intArray []int
}

func (s *sliceArray) Partition(left, right int) int {
	var pivot int
	pivot = left

	// Code to be executed until left becomes greater than right
	for left < right {
		// increment left pointer until the value at left pointer is greater than value at pivot pointer
		// We might not see an element where value of left pointer is greater than the value of pivot pointer, in that case, we iterate until left becomes == right
		for s.intArray[left] <= s.intArray[pivot] && left < right {
			left++
		}

		// decrement right pointer until the value at right pointer is less than value at pivot pointer
		for s.intArray[right] > s.intArray[pivot] {
			right--
		}

		// if we find an instance where value at left pointer is greater than pivot and value at right pointer is less than pivot, then swap both the values at left and right pointers
		if left < right {
			s.intArray[left], s.intArray[right] = s.intArray[right], s.intArray[left]
		}

		// there might be instances where pivot becomes the right element, in that case, swap left and pivot elements
		if right == pivot {
			s.intArray[left], s.intArray[pivot] = s.intArray[pivot], s.intArray[left]
		}
	}
	// when left > right, then swap pivot and right
	s.intArray[pivot], s.intArray[right] = s.intArray[right], s.intArray[pivot]
	return right
}

func QuickSort(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}

	if len(array) == 1 {
		return array
	}

	if len(array) == 2 {
		if array[0] > array[1] {
			array[0], array[1] = array[1], array[0]
		}
		return array
	}

	var s sliceArray
	s.intArray = array
	low := 0
	high := len(array) - 1
	var pivot int
	if low < high {
		pivot = s.Partition(low, high)
		left := sliceArray{
			intArray: s.intArray[0:pivot],
		}
		right := sliceArray{
			intArray: s.intArray[pivot:],
		}
		QuickSort(left.intArray)
		QuickSort(right.intArray)
	}

	return s.intArray
}

func main() {
	var s sliceArray
	s.intArray = []int{8, 5, 2, 9, 5, 6, 3}
	fmt.Println(QuickSort(s.intArray))
	s.intArray = []int{10, 90, 80, 70, 5, 75, 40, 65}
	fmt.Println(QuickSort(s.intArray))
	s.intArray = []int{3, 1, 2}
	fmt.Println(QuickSort(s.intArray))
	s.intArray = []int{-4, 5, 10, 8, -10, -6, -4, -2, -5, 3, 5, -4, -5, -1, 1, 6, -7, -6, -7, 8}
	fmt.Println(QuickSort(s.intArray))
	s.intArray = []int{-4, -7, -6, -7, -10, -6, -4, -5, -5, -4}
	fmt.Println(QuickSort(s.intArray))
}
