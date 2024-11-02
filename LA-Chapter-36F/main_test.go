package main

import (
	"testing"
)

// BenchmarkSum benchmarks the Sum function
func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}
