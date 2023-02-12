package main

import "fmt"

func main(){
	n := 1000
	divisor := getDivisor(n)
	mathExpectation := getMathExpectation(n, divisor)
	median := getMedian(n, divisor)

	fmt.Println(median, mathExpectation)
}

func getDivisor(n int) int {
	sum := 0
	k := 1
	proportional := k

	for i:=1; i <= n ; i++{
		sum += k
		proportional -= 1
		if proportional == 0 {
			k+= 1
			proportional = k
		}
	}
	return sum
}

func getMathExpectation(n, divisor int) float64 {
	MathExpectation := 0.0
	k := 1
	proportional := k
	for i:=1; i <= n ; i++{
		MathExpectation += float64(k) * float64(k)/float64(divisor)

		proportional -= 1
		if proportional == 0 {
			k+= 1
			proportional = k
		}
	}
	return MathExpectation
}

func getMedian(n, divisor int) int {
	sum := 0
	k := 1
	proportional := k

	for i := 1; i <= n; i++ {
		sum += k
		if float64(sum)/float64(divisor) >= 0.5 {
			return k
		}
		proportional -= 1
		if proportional == 0 {
			k += 1
			proportional = k
		}
	}
	return 0
}