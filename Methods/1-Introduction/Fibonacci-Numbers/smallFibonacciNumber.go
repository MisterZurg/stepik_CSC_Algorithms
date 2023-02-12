package main

import "fmt"

func smallFibonacciNumber(num int) int {
	fibNumbers := make([]int, num+1)
	// Sequence begins from 1
	fibNumbers[0] = 1
	fibNumbers[1] = 1
	for i := 2; i < num; i++ {
		fibNumbers[i] = fibNumbers[i-1] + fibNumbers[i-2]
	}
	return fibNumbers[num-1]
}

func main() {
	// put your code here
	var n int
	fmt.Scan(&n)
	fmt.Println(smallFibonacciNumber(n))
}