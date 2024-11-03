package main

// Sum calculates the sum of elements in a slice
func Sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
