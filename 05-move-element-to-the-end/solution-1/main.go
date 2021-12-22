package main

import "fmt"

func MoveElementToEnd(array []int, toMove int) []int {
	// Write your code here.
	var finalArray []int
	var count int
	if len(array) > 0 {
		for _, element := range array {
			if element != toMove {
				finalArray = append(finalArray, element)
			} else {
				count++
			}
		}

		for i := 0; i < count; i++ {
			finalArray = append(finalArray, toMove)
		}
		return finalArray
	}
	return []int{}
}

func main() {
	fmt.Println(MoveElementToEnd([]int{2, 1, 2, 2, 2, 3, 4, 2}, 2))
}
