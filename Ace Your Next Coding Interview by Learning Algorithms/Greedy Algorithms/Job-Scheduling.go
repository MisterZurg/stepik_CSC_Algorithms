package main

import (
	"fmt"
	"sort"
)

type JobsWithDeadlines struct {
	name     string // field for ordering
	profit   int
	deadline int
}

type Slot struct {
	jobName string
	present bool
}

func JobScheduling(jobs []JobsWithDeadlines, trueDeadLine int) ([]Slot, int) {
	maxProfit := 0

	slots := make([]Slot, trueDeadLine)

	// O(n log n)
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].profit > jobs[j].profit
	})

	for i := range jobs {
		for j := range slots {
			if slots[j].present == false && jobs[i].deadline-1 <= j {
				slots[j].jobName = jobs[i].name
				slots[j].present = true
				maxProfit += jobs[i].profit
				break
			}
		}
	}
	return slots, maxProfit
}

func main() {
	jobs := []JobsWithDeadlines{
		{"A", 50, 2},
		{"B", 20, 1},
		{"C", 30, 2},
		{"D", 25, 1},
		{"E", 15, 3},
	}

	// Otherwise we just go through the jobs slice and find the lat one
	trueDeadLine := 3

	seq, profit := JobScheduling(jobs, trueDeadLine)
	for _, jb := range seq {
		fmt.Printf("%s", jb.jobName)
	}
	fmt.Printf(" = %d", profit)
}
