package main

import "fmt"

func LastDigitOfFibonacciNumber(n int32) int8 {
	if n <= 1 {
		return 1
	}
	prev, curr := 1, 1
	for i := int32(3); i <= n; i++ {
		curr, prev = (curr+prev)%10, curr
	}
	return int8(curr)
}

func main() {
	// put your code here
	var n int32
	fmt.Scan(&n)
	fmt.Println(LastDigitOfFibonacciNumber(n))
}
