package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

//...
var (
	Reader = bufio.NewReader(os.Stdin)
	Writer = bufio.NewWriter(os.Stdout)
)

//Printf is for buffered output
func Printf(format string, vals ...interface{}) {
	fmt.Fprintf(Writer, format, vals...)
}

//Scanf is for buffered input
func Scanf(format string, vals ...interface{}) {
	fmt.Fscanf(Reader, format, vals...)
}

type proc struct {
	number      int
	releaseTime int
}

type ProcHeap []proc

func (h ProcHeap) Len() int {
	return len(h)
}

func (h ProcHeap) Less(i, j int) bool {
	return h[i].releaseTime < h[j].releaseTime ||
		h[i].releaseTime == h[j].releaseTime &&
			h[i].number < h[j].number
}

func (h ProcHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ProcHeap) Push(x interface{}) {
	*h = append(*h, x.(proc))
}

func (h *ProcHeap) Pop() interface{} {
	old := *h
	n := len(old)
	*h = old[:n-1]
	val := old[n-1]
	return val
}

func main() {
	defer Writer.Flush()

	var n, m int
	Scanf("%d %d\n", &n, &m)
	t := make([]int, m)
	for i := range t {
		Scanf("%d", &t[i])
	}

	ph := ProcHeap{}
	for i := 0; i < n; i++ {
		heap.Push(&ph, proc{i, 0})
	}
	for _, v := range t {
		worker := ph[0]
		heap.Pop(&ph)
		Printf("%d %d\n", worker.number, worker.releaseTime)
		worker.releaseTime += v
		heap.Push(&ph, worker)
	}
}
