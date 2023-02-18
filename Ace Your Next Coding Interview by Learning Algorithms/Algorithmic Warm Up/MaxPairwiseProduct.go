package main

import "fmt"

func MaxPairwiseProductNaive(numbers []int) int {
	n := len(numbers)
	maxProduct := 0

	for first := 0; first < n; first++ {
		for second := first + 1; second < n; second++ {
			maxProduct = max(maxProduct, numbers[first]*numbers[second])
		}
	}
	return maxProduct
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scan(&n)

	inputNumbers := make([]int, n)
	for i := range inputNumbers {
		fmt.Scan(&inputNumbers[i])
	}

	fmt.Println(MaxPairwiseProductNaive(inputNumbers))
}
