package main

import "fmt"

func Merge(left, right []int) []int {
	var finalArray []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] > right[0] {
			finalArray = append(finalArray, right[0])
			right = right[1:]
		} else if left[0] < right[0] {
			finalArray = append(finalArray, left[0])
			left = left[1:]
		} else {
			finalArray = append(finalArray, left[0])
			finalArray = append(finalArray, right[0])
			left = left[1:]
			right = right[1:]
		}
	}

	if len(left) == 0 && len(right) != 0 {
		for len(right) != 0 {
			finalArray = append(finalArray, right[0])
			right = right[1:]
		}
	}

	if len(left) != 0 && len(right) == 0 {
		for len(left) != 0 {
			finalArray = append(finalArray, left[0])
			left = left[1:]
		}
	}
	return finalArray
}

func MergeSort(array []int) []int {
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

	low := 0
	high := len(array) - 1
	mid := (low + high) / 2
	return Merge(MergeSort(array[:mid+1]), MergeSort(array[mid+1:]))
}

func main() {
	fmt.Println(MergeSort([]int{10, 90, 80, 70, 5, 35, 75, 40, 65}))
	// fmt.Println(MergeSort([]int{-4, 5, 10, 8, -10, -6, -4, -2, -5, 3, 5, -4, -5, -1, 1, 6, -7, -6, -7, 8}))
}
