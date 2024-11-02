package main

import "fmt"

// Sum calculates the sum of elements in a slice
func Sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(nums))
}
