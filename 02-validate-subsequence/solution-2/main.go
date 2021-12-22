package main

import "fmt"

//IsValidSubsequence - takes 2 arrays and checks if the array2 is a subsequence of array1
func IsValidSubsequence(array []int, sequence []int) bool {
	arrayPosition := 0
	sequencePosition := 0

	var finalList []int

	for sequencePosition < len(sequence) && arrayPosition < len(array) {
		if sequence[sequencePosition] == array[arrayPosition] {
			finalList = append(finalList, sequence[sequencePosition])
			sequencePosition++
			arrayPosition++
		} else {
			arrayPosition++
		}
	}

	if len(finalList) == len(sequence) {
		return true
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
