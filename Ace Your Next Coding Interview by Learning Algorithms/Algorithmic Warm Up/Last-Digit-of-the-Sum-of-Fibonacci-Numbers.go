package main

import "fmt"

func LastDigitoftheSumofFibonacciNumbers(n int) int {
	pisano := getPisanoPeriod()

	n_periodic := n % len(pisano)

	sum := 0
	// <= Cause my F0 = 0
	for i := 0; i <= n_periodic; i++ {
		sum = (sum + pisano[i]) % 10
	}
	return sum
}

// getPisanoPeriod returns PP by modulo 10
// The Fibonacci sequence by mod 10 is periodic
// — the last digits repeat themselves with the Pisano period of length 60
func getPisanoPeriod() []int {
	var pisano []int
	pisano = append(pisano, 0, 1, 1)
	for i := 2; ; i++ {
		if i > 2 && pisano[i-1] == 0 && pisano[i] == 1 {
			break
		}
		fib := pisano[i-1] + pisano[i]
		pisano = append(pisano, fib%10)
	}
	return pisano[:len(pisano)-2] // crop last two digits
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(LastDigitoftheSumofFibonacciNumbers(n))
}
