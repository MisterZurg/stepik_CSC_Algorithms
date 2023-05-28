package main

import "fmt"

func HugeFibonacciNumber(n int, m int) int {
	if n <= 1 {
		return 1
	}
	pisano := []int{0, 1, 1}
	for i := 3; !(pisano[i-2] == 0 && pisano[i-1] == 1); i++ {
		/* Вместо сомнительного условия в for
		//if pisano[i - 2] == 0 && pisano[i - 1] == 1 {
		//	break
		}
		*/
		pisano = append(pisano, (pisano[i-2]+pisano[i-1])%m)
	}

	pisano = pisano[:len(pisano)-2]
	fmt.Println(pisano)
	return pisano[n%len(pisano)]
}

func main() {
	// put your code here
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(HugeFibonacciNumber(n, m))
	// fmt.Println(HugeFibonacciNumber(2816213588,239))
}
