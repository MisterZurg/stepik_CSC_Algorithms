package main

import (
	"fmt"
	"strconv"
	"strings"
)

func MaximumNumberOfPrizes(target int) {
	k := 1
	for k*(k+1)/2 <= target {
		k++
	}
	// Возвращаемся, на момент когда оно меньше|равно
	k--

	delta := target - k*(k+1)/2

	ks := make([]string, 0)
	for i := 1; i < k; i++ {
		ks = append(ks, strconv.Itoa(i))
	}
	ks = append(ks, strconv.Itoa(k+delta))
	fmt.Println(len(ks))
	fmt.Println(strings.Join(ks, " "))
}

func main() {
	var n int
	fmt.Scan(&n)
	MaximumNumberOfPrizes(n)
}
