package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sortedSlice := makeSlice(n)

	var m int
	fmt.Scan(&m)
	numbersToSearch := makeSlice(m)

	for _, number := range numbersToSearch {
		fmt.Println(BinarySearch(sortedSlice, number))
	}
}

func BinarySearch(sorted []int, target int) int{
	leftIdx := 0
	rightIdx := len(sorted) - 1

	for leftIdx <= rightIdx {
		middleIdx := (leftIdx + rightIdx) / 2

		if sorted[middleIdx] == target {
			// По сути мы возвращаем позицию, поэтому +1
			return middleIdx + 1
		} else if sorted[middleIdx] > target {
			rightIdx = middleIdx - 1
		} else {
			leftIdx = middleIdx + 1
		}

	}
	return -1
}

func makeSlice(size int) []int  {
	slice := make([]int, size)
	for i := range slice {
		fmt.Scan(&slice[i])
	}
	return slice
}
