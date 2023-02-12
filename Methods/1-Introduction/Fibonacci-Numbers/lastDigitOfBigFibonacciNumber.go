package main

import "fmt"

func lastDigitOfBigFibonacciNumber(n int) uint { // *big.Int {
	switch n {
	case 1:
		return 1
	case 2:
		return 1
	default:
		var fn, fn_m1 uint = 1, 1
		for i := 2; i < n; i++ {
			fn_m1, fn = fn, (fn+fn_m1)%10
		}
		return fn
	}
}

func main() {
	// put your code here
	var n int
	fmt.Scan(&n)
	fmt.Println(lastDigitOfBigFibonacciNumber(n))
}