package main

import "fmt"

func threeNumberSort(array []int, order []int) []int {
	arrayInfo := make(map[int]int)
	for i := 0; i < len(array); i++ {
		arrayInfo[array[i]]++
	}

	array = []int{}
	for i := 0; i <= len(order)-1; i++ {
		for j := 0; j < arrayInfo[order[i]]; j++ {
			array = append(array, order[i])
		}
	}
	return array
}

func main() {
	fmt.Println(threeNumberSort([]int{1, 0, 0, -1, -1, 0, 1, 1}, []int{0, 1, -1}))
}
