package main

import (
	"fmt"
)

func IsValidSubsequence(array []int, sequence []int) bool {
	var position []int
	var counter int
	var hashMap = make(map[int]int)

	for _, value := range sequence {
		hashMap[value]++
	}

	switch {
	case len(array) < len(sequence):
		return false

	case len(sequence) == 1:
		for i := 0; i < len(array); i++ {
			if array[i] == sequence[0] {
				return true
			}
		}

	case len(hashMap) == 1:
		var arrayHash = make(map[int]int)
		for _, value := range array {
			arrayHash[value]++
		}
		for key := range hashMap {
			if arrayHash[key] >= hashMap[key] {
				return true
			}
		}

	default:
		for _, value := range sequence {
			for i := 0; i < len(array); i++ {
				if array[i] == value {
					position = append(position, i)
				}
			}
		}

		for i := 0; i < len(position)-1; i++ {
			if position[i+1]-position[i] > 0 {
				counter++
			}
		}

		if counter == len(sequence)-1 {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(IsValidSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10}))
	fmt.Println("=======")
	fmt.Println(IsValidSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{6, 1, -1, 10}))
	fmt.Println("=======")
	fmt.Println(IsValidSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{6, 6, -1, 10}))
	fmt.Println("=======")
	fmt.Println(IsValidSubsequence([]int{1, 1, 1, 1, 1}, []int{1, 1, 1}))
}
