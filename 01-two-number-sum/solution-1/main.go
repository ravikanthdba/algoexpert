package main

import "fmt"

// TwoNumberSum - takes an array and target sum, and returns the pair of integers which return the sum
func TwoNumberSum(array []int, target int) []int {
	var final []int
	for firstNumber := 0; firstNumber < len(array); firstNumber++ {
		for secondNumber := firstNumber + 1; secondNumber < len(array); secondNumber++ {
			if array[firstNumber]+array[secondNumber] == target {
				final = append(final, array[firstNumber])
				final = append(final, array[secondNumber])
			} else {
				continue
			}
		}
	}

	if len(final) == 0 {
		return []int{}
	}
	return final
}

func main() {
	fmt.Println(TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10))
	fmt.Println(TwoNumberSum([]int{-21, 101, 12, 4, 65, 56, 210, 356, 9, -47}, 164))
	fmt.Println(TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 15))
}
