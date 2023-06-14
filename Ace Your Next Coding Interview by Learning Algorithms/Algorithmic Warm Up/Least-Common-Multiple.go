package main

import "fmt"

func LCM(a, b int) int {
	return (a * b) / GDC(a, b)
}

func GDC(a, b int) int {
	if b == 0 {
		return a
	}
	return GDC(b, a%b)
}

func main() {
	// put your code here
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(LCM(a, b))
}
