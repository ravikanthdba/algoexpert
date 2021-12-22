package main

import "fmt"

func bubblesorting(array []int) []int {
	if len(array) == 1 {
		return array
	}
	if len(array) >= 2 {
		working := true
		for working {
			working = false
			for i := 0; i < len(array)-1; i++ {
				if array[i] > array[i+1] {
					array[i], array[i+1] = array[i+1], array[i]
					working = true
				}
			}
		}
		return array
	}
	return []int{}
}

func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}
	fmt.Println(bubblesorting(array))
}
