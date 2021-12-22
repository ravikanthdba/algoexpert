package main

import "fmt"

func selectionSort(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}

	if len(array) == 1 {
		return array
	}

	// Steps for a selection sort
	// Step1: traverse through the array once, pick up each and every element.
	// Step2: Identify a number which is smallest than the number we've picked up in step 1
	// Step3: Identify, if the number picked up in step 2 is actually the smallest number in the remaining array.
	// Step4: If yes, then swap the numbers
	// Step5: If no, then update the current smallest number to the latest number
	// Step6: Repeat steps 1 through 5 untl all elements in the list are covered.
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
