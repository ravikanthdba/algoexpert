package main

import (
	"math/rand"
	"testing"
)

func generateRandom(n int) []int {
	arrayOfIntegers := []int{}
	for i := 0; i < n; i++ {
		number := rand.Intn(100)
		arrayOfIntegers = append(arrayOfIntegers, number)
	}
	return arrayOfIntegers
}

func Benchmark_IsMonotonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		number := rand.Intn(10)
		numberList := generateRandom(number)
		IsMonotonic(numberList)
	}
}
