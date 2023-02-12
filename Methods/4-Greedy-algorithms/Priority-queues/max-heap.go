package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// PriorityQueue is Binary MaxHeap
type PriorityQueue struct {
	MaxHeap []int
	HeapSize int
}

func (pq *PriorityQueue) Insert(value int){
	pq.HeapSize += 1
	pq.MaxHeap = append(pq.MaxHeap, value)
	// После добавления элемнта в кучу, возможно нарушение её главного свойства
	// (каждый элемент больше своих потомков). Поэтому нужно пофиксить эего используя
	// функцию просеивания вверх siftUp.
	pq.siftUp(pq.HeapSize - 1)
}

// O(log n)
func (pq *PriorityQueue) siftUp(i int) {
	// Двигаемся до того момента,
	// пока не станет выполнятся свойство кучи
	// MaxHeap[(i - 1) / 2] < i[i]
	for pq.MaxHeap[i] > pq.MaxHeap[(i - 1) / 2] {
		// Swap them
		pq.Swap(i,(i - 1) / 2)
		i = (i - 1) / 2
	}
}

// ExtractMax returns the maximum element and removes it from the PriorityQueue
func (pq *PriorityQueue) ExtractMax() int {
	if pq.HeapSize == 0 {
		return 0
	}
	max := pq.MaxHeap[0]
	// Put last elem; on poped elem place
	pq.MaxHeap[0] = pq.MaxHeap[pq.HeapSize - 1]
	pq.HeapSize -= 1
	// Crop ast elem;
	pq.MaxHeap = pq.MaxHeap[:pq.HeapSize]

	pq.siftDown(0)
	return max
}

func (pq *PriorityQueue) siftDown(i int) {
	for 0 < pq.maxChild(i) && pq.MaxHeap[i] < pq.MaxHeap[pq.maxChild(i)] {
		maxChild := pq.maxChild(i)
		pq.Swap(i, maxChild)
		i = maxChild
	}


		//     11					11
		//    /  \				   /  \
		//  34    25			  4    25


		//        11
		//     11    10
		//  11   11
		// 11
		// Берем максимальный из двух и меняем их местами
}

func (pq *PriorityQueue) maxChild(i int) int {
	var child int

	if 2*i+1 < pq.HeapSize {
		child = 2*i + 1
	}

	if 2*i+2 < pq.HeapSize && pq.MaxHeap[child] < pq.MaxHeap[2*i+2] {
		child = 2*i + 2
	}

	return child
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.MaxHeap[i], pq.MaxHeap[j] = pq.MaxHeap[j], pq.MaxHeap[i]
}


func main() {
	// put your code here
	var pq PriorityQueue
	ParseInput(&pq)
}

func ParseInput(pq *PriorityQueue) {
	scanner := bufio.NewScanner(os.Stdin)
	var text string

	for scanner.Scan() {
		text = scanner.Text()
		fields := strings.Fields(text)

		if fields[0] == "Insert" {
			number, _ := strconv.Atoi(fields[1])
			pq.Insert(number)
		}
		if fields[0] == "ExtractMax" {
			fmt.Println(pq.ExtractMax())
		}
	}
}

func SlowParseInput(pq *PriorityQueue) {
	var numOfOp int
	fmt.Scan(&numOfOp)
	var operation string
	for i := 0; i < numOfOp; i++ {
		fmt.Scan(&operation)
		switch operation {
		case "Insert":
			var insNumber int
			fmt.Scan(&insNumber)
			pq.Insert(insNumber)
		case "ExtractMax":
			fmt.Println(pq.ExtractMax())
		}
	}
}


/*
Sample Input:

6
Insert 200
Insert 10
ExtractMax
Insert 5
Insert 500
ExtractMax

Sample Output:

200
500



8
Insert 87345
Insert 13562
Insert 63521
Insert 72351
Insert 97235
ExtractMax
ExtractMax
ExtractMax
ExtractMax
ExtractMax


97235
87345
72351
63521
13562

5
Insert 10
Insert 10
Insert 8
ExtractMax
ExtractMax

6
Insert 8
Insert 2
ExtractMax
ExtractMax
ExtractMax
ExtractMax
 */