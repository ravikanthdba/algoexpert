package main

import "fmt"

func IsMonotonic(array []int) bool {
	if len(array) <= 2 {
		return true
	}
	incCounter := 0
	decCounter := 0
	index := 0
	working := true
	for working && index < len(array)-1 {
		working = false
		if array[index] == array[index+1] && index < len(array)-1 {
			working = true
			index++
			continue
		} else if array[index] > array[index+1] {
			if incCounter > 0 {
				return false
			}
			decCounter++
		} else if array[index] < array[index+1] {
			if decCounter > 0 {
				return false
			}
			incCounter++
		}
		working = true
		index++
	}
	return true
}

func main() {
	fmt.Println(IsMonotonic([]int{-1, -1, -2, -3, -4, -5, -5, -5, -6, -7, -8, -8, -9, -10, -11}))
}
