package main

import "fmt"

func selectionSort(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}

	if len(array) == 1 {
		return array
	}

	for i := range array {
		location := i
		currentMinValue := array[i]
		for j := i + 1; j < len(array); j++ {
			if currentMinValue > array[j] {
				currentMinValue = array[j]
				location = j
			}
		}

		array[i], array[location] = array[location], array[i]
	}
	return array
}

func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}
	fmt.Println(selectionSort(array))
}
