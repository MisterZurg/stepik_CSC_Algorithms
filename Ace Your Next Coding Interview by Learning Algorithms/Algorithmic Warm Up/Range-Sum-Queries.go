package main

import "fmt"

type Query struct {
	left  int
	right int
}

func RangeSumQueries(arr []int, queries []Query) {
	prefixSum := getPrefixSums(arr)

	for _, query := range queries {
		if query.left > 0 {
			fmt.Println(prefixSum[query.right] - prefixSum[query.left-1])
		} else {
			fmt.Println(prefixSum[query.right])
		}
	}
}

func getPrefixSums(arr []int) []int {
	n := len(arr)
	prefixSum := make([]int, n)
	prefixSum[0] = arr[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + arr[i]
	}
	return prefixSum
}

func main() {
	arr := []int{2, -1, 7, 2, -3, -2, 4}
	q := []Query{
		{1, 3},
		{2, 5},
	}
	RangeSumQueries(arr, q)
}
