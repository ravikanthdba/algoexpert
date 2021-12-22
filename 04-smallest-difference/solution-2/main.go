package main

import (
	"fmt"
	"math"
	"sort"
)

func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	idxOne := 0
	idxTwo := 0
	smallest := math.Inf(+10)
	current := math.Inf(+10)
	smallestPair := []int{}
	for idxOne < len(array1) && idxTwo < len(array2) {
		firstNum := array1[idxOne]
		secondNum := array2[idxTwo]
		switch {
		case firstNum < secondNum:
			current = float64(secondNum) - float64(firstNum)
			idxOne++
		case firstNum > secondNum:
			current = float64(firstNum) - float64(secondNum)
			idxTwoc++
		case firstNum == secondNum:
			return []int{firstNum, secondNum}
		}
		if smallest > current {
			smallest = current
			smallestPair = []int{firstNum, secondNum}
		}
	}
	return smallestPair
}

func main() {
	fmt.Println(SmallestDifference([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 14, 17}))
}
