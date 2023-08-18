package main

import (
	"fmt"
	"math"
	"strconv"
	//"math"
)

func LargestConcatenate(numbers []int) string {
	var yourSalary string
	n := len(numbers)
	for n > 0 {
		maxNumber := math.MinInt
		var maxNumberIdx int
		for i := 0; i < n; i++ {
			number := numbers[i]
			if IsBetter(number, maxNumber) {
				maxNumber = number
				maxNumberIdx = i
			}
		}
		yourSalary = fmt.Sprintf("%s%s", yourSalary, strconv.Itoa(maxNumber))
		// remove maxNumber from Numbers
		numbers[maxNumberIdx], numbers[n-1] = numbers[n-1], numbers[maxNumberIdx]
		n--
	}

	return yourSalary
}

func IsBetter(a, b int) bool {
	// Convertion TC: O(n)
	ab, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	ba, _ := strconv.Atoi(fmt.Sprintf("%d%d", b, a))

	return ab > ba
}

func main() {
	var n int
	fmt.Scan(&n)
	numbers := make([]int, n)
	for i := range numbers {
		fmt.Scan(&numbers[i])
	}
	fmt.Println(LargestConcatenate(numbers))
}
