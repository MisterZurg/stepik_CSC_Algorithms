package main

import (
	"fmt"
	"math"
	"sort"
)

func MiceAndAFox(mice, holes []int) int {
	distance := -1
	sort.Ints(mice)
	sort.Ints(holes)

	for i := range holes {
		// Yeah, go 1.21 but still, no abs for ints
		distance = max(distance, int(math.Abs(float64(holes[i]-mice[i]))))
	}

	return distance
}

func main() {
	mice := []int{13, 4, 2}
	holes := []int{11, 9, 6}
	fmt.Println(MiceAndAFox(mice, holes))
}
