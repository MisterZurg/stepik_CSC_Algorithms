package main

import (
	"fmt"
	"sort"
)

type Segment struct {
	l int
	r int
}

func CollectingSignatures(segments []Segment) []int {
	n := len(segments)
	/*
		   ----
		---
		 ----
		    --
	*/
	// Отсортировали отрезки
	sortSegmentsByLeftEdge(segments)
	/*
		---
		 ----
		    --
		   ----
	*/
	var points []int
	slow := 0
	fast := 1
	for slow < n && fast < n {
		if segments[slow].r >= segments[fast].l {
			fast++
		} else {
			points = append(points, segments[slow].r)
			slow = fast
			fast++
		}
	}
	// Добавляем последнюю точку
	points = append(points, segments[slow].r)
	return points
}

func sortSegmentsByLeftEdge(segments []Segment) {
	sort.Slice(segments, func(i, j int) bool {
		return segments[i].r < segments[j].r
	})
}

func main() {
	var n int
	fmt.Scan(&n)
	segments := make([]Segment, n)
	for i := range segments {
		fmt.Scan(&segments[i].l, &segments[i].r)
	}

	points := CollectingSignatures(segments)
	fmt.Println(len(points))
	for _, point := range points {
		fmt.Printf("%d ", point)
	}
}

/*
4
4 7
1 3
2 5
5 6

2
3 6
*/
//https://stepik.org/lesson/13238/step/4?unit=3424
