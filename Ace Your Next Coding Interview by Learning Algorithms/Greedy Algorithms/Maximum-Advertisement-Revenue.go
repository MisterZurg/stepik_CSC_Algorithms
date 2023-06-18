package main

import (
	"fmt"
	"sort"
)

func MaximumAdvertisementRevenue(prices, clicks []int) int {
	n := len(prices)
	// O(n log n)
	sort.Ints(prices)
	sort.Ints(clicks)

	product := 1
	product = prices[0] * clicks[0]
	for i := 1; i < n; i++ {
		product += prices[i] * clicks[i]
	}
	return product
}

func main() {
	var n int
	fmt.Scan(&n)

	prices := make([]int, n)
	for i := range prices {
		fmt.Scan(&prices[i])
	}

	clicks := make([]int, n)
	for i := range clicks {
		fmt.Scan(&clicks[i])
	}
	fmt.Println(MaximumAdvertisementRevenue(prices, clicks))
}
