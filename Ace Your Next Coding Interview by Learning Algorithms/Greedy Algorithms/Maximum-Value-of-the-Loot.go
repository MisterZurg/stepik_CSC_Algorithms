package main

import (
	"fmt"
	"sort"
)

func MaximumValueOfTheLoot(bpWeight int, compounds []Compound) float64 {
	// We are allowed to split Compounds
	for i := range compounds {
		compounds[i].partitionCost = float64(compounds[i].cost) / float64(compounds[i].weight)
	}
	// Sort them by Cost / Weight DESC 4 3 2
	SortCompounds(compounds)

	var value float64
	for _, compound := range compounds {
		if bpWeight >= compound.weight {
			value += float64(compound.cost)
			bpWeight -= compound.weight
		} else {
			value += float64(bpWeight) * compound.partitionCost
			break
		}
	}
	return value
}

func SortCompounds(compounds []Compound) {
	sort.Slice(compounds, func(i, j int) bool {
		return compounds[i].partitionCost > compounds[j].partitionCost
	})
}

type Compound struct {
	cost          int
	weight        int
	partitionCost float64
}

func main() {
	// put your code here
	var n, w int // compounds and backpack capacity
	fmt.Scan(&n, &w)
	compounds := make([]Compound, n)

	for i := range compounds {
		fmt.Scan(&compounds[i].cost, &compounds[i].weight)
	}

	fmt.Println(MaximumValueOfTheLoot(w, compounds))
}
