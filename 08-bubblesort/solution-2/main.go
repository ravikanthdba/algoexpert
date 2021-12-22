package main

import "fmt"

func bubblesort(array []int) []int {
	if len(array) == 1 {
		return array
	}
	if len(array) == 0 {
		return []int{}
	}
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}

func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}
	fmt.Println(bubblesort(array))
}
