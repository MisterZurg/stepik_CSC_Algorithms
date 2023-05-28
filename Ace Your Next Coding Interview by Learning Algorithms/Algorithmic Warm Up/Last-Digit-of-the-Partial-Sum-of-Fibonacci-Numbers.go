package main

import "fmt"

func LastDigitofthePartialSumofFibonacciNumbers(Fstart, Fend int) int {
	pisano := getPisanoPeriod()

	FstartMod10 := Fstart % len(pisano)
	FendMod10 := Fend % len(pisano)

	partialSum := 0
	if FstartMod10 > FendMod10 {
		FendMod10 += 60
	}

	for i := FstartMod10; i <= FendMod10; i++ {
		partialSum = (partialSum + pisano[i%60]) % 10
	}

	return partialSum
}

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
	var Fstart, Fend int
	fmt.Scan(&Fstart, &Fend)
	fmt.Println(LastDigitofthePartialSumofFibonacciNumbers(Fstart, Fend))
}
