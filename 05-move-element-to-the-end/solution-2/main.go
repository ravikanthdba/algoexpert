package main

import "fmt"

func MoveElementToTheEnd(array []int, toMove int) []int {
	startPointer := 0
	endPointer := len(array) - 1
	for startPointer < endPointer {
		for startPointer < endPointer && array[endPointer] == toMove {
			endPointer--
		}

		if array[startPointer] == toMove {
			array[startPointer], array[endPointer] = array[endPointer], array[startPointer]
		}

		startPointer++
	}
	return array
}

func main() {
	fmt.Println(MoveElementToTheEnd([]int{2, 1, 2, 2, 2, 3, 4, 2}, 2))
}
