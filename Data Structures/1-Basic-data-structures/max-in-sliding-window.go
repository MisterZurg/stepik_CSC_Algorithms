package main

import (
	"fmt"
)

func main() {
	// put your code here
	elems, window := parseInputMaxInSlidingWindow()
	SolutionMaxSlidingWindow(elems, window)
}

type Deque []int

func (dq *Deque) Front() int {
	return (*dq)[0]
}

func (dq *Deque) Back() int {
	return (*dq)[len(*dq)-1]
}

func (dq *Deque) PushBack(v int) {
	*dq = append(*dq, v)
}

func (dq *Deque) PopFront() int {
	p := (*dq)[0]
	*dq = (*dq)[1:]
	return p
}

func (dq *Deque) PopBack() int {
	p := (*dq)[len(*dq)-1]
	*dq = (*dq)[:len(*dq)-1]
	return p
}

func parseInputMaxInSlidingWindow() ([]int, int) {
	var n int
	fmt.Scan(&n)

	elems := make([]int, n)
	for i := range elems {
		fmt.Scan(&elems[i])
	}

	var window int
	fmt.Scan(&window)

	return elems, window
}

func SolutionMaxSlidingWindow(elems []int, window int) {
	res := maxSlidingWindow(elems, window)
	for _, num := range res {
		fmt.Printf("%d ", num)
	}
}

//Initialise the first K elements of the array into the deque.
//Iterate over the input array and for each step:
//	- Consider only the indices of the elements in the current window.
//	- Pop-out all the indices of elements smaller than the current element, since their value will be less than the current element.
//	- Push the current element into the deque.
//	- Push the first element of the deque i.e. deque[0] into the output array.
func maxSlidingWindow(elems []int, window int) []int {
	res := make([]int, 0, len(elems)-window+1)
	dq := make(Deque, 0, window+1)

	for i, v := range elems {
		if i > window-1 && dq.Front() <= i-window {
			dq.PopFront()
		}
		for len(dq) > 0 && v > elems[dq.Back()] {
			dq.PopBack()
		}
		dq.PushBack(i)
		if i >= window-1 {
			res = append(res, elems[dq.Front()])
		}
	}
	return res
}
