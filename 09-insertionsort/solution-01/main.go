package main

import "fmt"

func insertionSort(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}
	if len(array) == 1 {
		return array
	}

	working := true
	i := 0
	for working {
		working = false
		for i < len(array) {
			for j := i; j > 0; j-- {
				if array[j] < array[j-1] {
					array[j], array[j-1] = array[j-1], array[j]
					working = true
				}
			}
			i++
		}
	}
	return array
}

func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}
	fmt.Println(insertionSort(array))
}
