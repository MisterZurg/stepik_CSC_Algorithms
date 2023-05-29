package main

import "fmt"

func LastDigitoftheSumofSquaresofFibonacciNumbers(n int) int {
	squaredPisano := getSquaredPisanoPeriod()

	n_mod_60 := n % len(squaredPisano)
	squaredSum := 0
	for i := 0; i <= n_mod_60; i++ {
		squaredSum = (squaredSum + squaredPisano[i]) % 10
	}
	return squaredSum
}

func getSquaredPisanoPeriod() []int {
	var pisano []int
	pisano = append(pisano, 0, 1, 1)
	for i := 2; ; i++ {
		if i > 2 && pisano[i-1] == 0 && pisano[i] == 1 {
			break
		}
		fib := pisano[i-1] + pisano[i]
		pisano = append(pisano, fib%10)
	}
	// crop last two digits
	pisano = pisano[:len(pisano)-2]
	//Square Pisano Loop
	for i := range pisano {
		pisano[i] = (pisano[i] * pisano[i]) % 10
	}
	return pisano
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(LastDigitoftheSumofSquaresofFibonacciNumbers(n))
}
