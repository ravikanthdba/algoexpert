package main

import "fmt"

//IsValidSubsequence - takes 2 arrays and checks if the array2 is a subsequence of array1
func IsValidSubsequence(array []int, sequence []int) bool {
	sequencePosition := 0
	for i := 0; i < len(array); i++ {
		if sequencePosition == len(sequence) {
			break
		}

		if sequence[sequencePosition] == array[i] {
			sequencePosition++
		}
	}
	return sequencePosition == len(sequence)
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
