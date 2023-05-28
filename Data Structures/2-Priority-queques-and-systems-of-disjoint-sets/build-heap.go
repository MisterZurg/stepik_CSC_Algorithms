package main

import (
	"fmt"
	"math"
)

func parseInput() []int {
	var n int
	fmt.Scan(&n)
	notHeap := make([]int, n)
	for i := range notHeap {
		fmt.Scan(&notHeap[i])
	}
	return notHeap
}

// Реализация кучи, как по лекциям,
// только методы, нужные для данной задачи
// и с модификацией чтобы получилась min-куча
type MinHeap []int

// Просеивание вниз
func (h *MinHeap) ShiftDown(i int) {
	minIndex := i
	l := h.leftChild(i)
	if l < h.Size() && (*h)[l] < (*h)[minIndex] {
		minIndex = l
	}

	r := h.rightChild(i)
	if r < h.Size() && (*h)[r] < (*h)[minIndex] {
		minIndex = r
	}

	if i != minIndex {
		(*h)[i], (*h)[minIndex] = (*h)[minIndex], (*h)[i]
		h.ShiftDown(minIndex)
	}
}

func (h *MinHeap) getParent(i int) int {
	return int(math.Floor(float64((i - 1) / 2)))
}

func (h *MinHeap) leftChild(i int) int {
	return 2*i + 1
}

func (h *MinHeap) rightChild(i int) int {
	return 2*i + 2
}
func (h *MinHeap) Size() int {
	return len(*h)
}

func main() {
	// читаем данные со стандартного ввода
	array := parseInput()

	heap := MinHeap(array)
	// исправляем массив в куче, чтобы он стал min-кучей,
	// с помощью просеивания вниз (при этом заполняется массив обменов swaps)
	for i := heap.Size()/2 - 1; i >= 0; i-- {
		heap.ShiftDown(i)
	}

	for _, elem := range heap {
		fmt.Printf("%d ", elem)
	}
}
