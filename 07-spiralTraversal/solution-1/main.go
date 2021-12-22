package main

import "fmt"

func spiralTraversal(array [][]int) {
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}

func main() {
	matrix := [][]int{}
	matrix.append(matrix, []int{1, 2, 3, 4})
	matrix.append(matrix, []int{12, 13, 14, 5})
	matrix.append(matrix, []int{11, 16, 15, 6})
	matrix.append(matrix, []int{10, 9, 8, 7})
	spiralTraversal(matrix)
}
