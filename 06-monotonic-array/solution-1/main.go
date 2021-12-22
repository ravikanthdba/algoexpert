package main

import "fmt"

func IsMonotonic(array []int) bool {
	var finalArray = make(map[string]int)
	if len(array) == 0 || len(array) == 1 {
		return true
	}

	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			finalArray["Increase"]++
		} else if array[i] < array[i+1] {
			finalArray["Decrease"]++
		} else if array[i] == array[i+1] {
			finalArray["Equal"]++
		}
	}

	if finalArray["Increase"] > 0 && finalArray["Decrease"] > 0 {
		return false
	}

	return true
}

func main() {
	fmt.Println(IsMonotonic([]int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}))
}
