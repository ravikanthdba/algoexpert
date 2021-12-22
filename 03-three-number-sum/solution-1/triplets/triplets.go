package triplets

import (
	"sort"
)

//ThreeNumberSum - takes array of integers and gives the set of integers where sum == target
func ThreeNumberSum(array []int, target int) [][]int {
	var finalResult [][]int
	sort.Ints(array)
	for i := 0; i < len(array)-2; i++ {
		for j := i + 1; j < len(array)-1; j++ {
			for k := j + 1; k < len(array); k++ {
				if array[i]+array[j]+array[k] == target {
					var simpleArray []int
					simpleArray = append(simpleArray, array[i], array[j], array[k])
					finalResult = append(finalResult, simpleArray)
				} else {
					continue
				}
			}
		}
	}
	if len(finalResult) != 0 {
		return finalResult
	}
	return [][]int{}
}
