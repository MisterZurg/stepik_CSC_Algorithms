package main

import "fmt"

func FibonacciNumber(n int32) int32 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	memo := make([]int32, n+1)
	memo[0], memo[1] = 0, 1
	for i := int32(2); i <= n; i++ {
		memo[i] = memo[i-2] + memo[i-1]
	}
	return memo[n]
}

func main() {
	// put your code here
	var n int32
	fmt.Scan(&n)
	fmt.Println(FibonacciNumber(n))
}
