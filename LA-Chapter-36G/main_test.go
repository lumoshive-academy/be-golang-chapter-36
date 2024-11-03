package main

import (
	"testing"
)

// BenchmarkSum benchmarks the Sum function with various slice sizes
func BenchmarkSum(b *testing.B) {
	b.Run("SmallSlice", func(b *testing.B) {
		numbers := []int{1, 2, 3, 4, 5}
		for i := 0; i < b.N; i++ {
			Sum(numbers)
		}
	})

	b.Run("MediumSlice", func(b *testing.B) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for i := 0; i < b.N; i++ {
			Sum(numbers)
		}
	})

	b.Run("LargeSlice", func(b *testing.B) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
		for i := 0; i < b.N; i++ {
			Sum(numbers)
		}
	})
}
