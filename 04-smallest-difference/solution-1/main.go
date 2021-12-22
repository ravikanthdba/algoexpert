package main

import (
	"fmt"
	"math"
	"sort"
)

//SmallestDifference - Takes 2 arrays and returns difference between element in each element between arrays and returns the pair with smallest difference
func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	difference := array1[len(array1)-1] + array2[len(array2)-1]
	var differenceArray []int
	for _, valueFromArray1 := range array1 {
		for _, valueFromArray2 := range array2 {
			if int(math.Abs(float64(valueFromArray1)-float64(valueFromArray2))) < difference {
				difference = int(math.Abs(float64(valueFromArray1) - float64(valueFromArray2)))
			} else {
				continue
			}
			differenceArray = []int{}
			differenceArray = append(differenceArray, valueFromArray1, valueFromArray2)
		}
	}

	if len(differenceArray) != 0 {
		return differenceArray
	}
	return []int{}
}

func main() {
	fmt.Println(SmallestDifference([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 14, 17}))
}
