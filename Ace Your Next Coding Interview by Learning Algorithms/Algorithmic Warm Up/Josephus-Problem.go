package main

import "fmt"

// Josephus() Slow algorithm for solving the Josephus Problem
// O(n log n)
func Josephus(n, k int) int {
	aliveOrDead := make([]bool, n)
	for position := 0; position < n; position++ {
		aliveOrDead[position] = true
	}

	numberOfSurvivors := n
	currentPosition := 0
	idx := 0

	for {
		if aliveOrDead[currentPosition] {
			idx++
		}
		if idx == k {
			if numberOfSurvivors == 1 {
				return currentPosition
			} else {
				aliveOrDead[currentPosition] = false
				idx = 0
				// fmt.Println(aliveOrDead)
			}
			numberOfSurvivors--
		}
		currentPosition = (currentPosition + 1) % n
	}
}

func FastJosephus(n, k int) int {
	// (old number − k) mod n = new number .
	old_number := n
	var new_number int
	if old_number-k > 0 {
		new_number = (old_number - k) % n
	} else {
		new_number = old_number
	}
	return FastJosephus(new_number-1, k)
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	//fmt.Println(Josephus(n, k))
	fmt.Println(FastJosephus(n, k))
}
