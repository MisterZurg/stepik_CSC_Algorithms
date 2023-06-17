package main

import "fmt"

func CarFueling(distance, fuel int, stops []int) int {
	refills := 0
	currDist := 0
	currTank := fuel

	for i := 0; i < len(stops); i++ {
		if stops[i]-currDist <= currTank {
			currTank -= stops[i] - currDist
			currDist = stops[i]
			//fmt.Println("Passed stop")
		} else if stops[i]-currDist <= fuel {
			currTank = fuel - (stops[i] - currDist)
			currDist = stops[i]
			refills++
			//fmt.Println("Passed stop and refield")
		} else {
			return -1
		}
	}

	if distance-currDist <= currTank {
		return refills
	} else if distance-currDist <= fuel {
		return refills + 1
	}
	return -1
}

func main() {
	var d, m, n int
	fmt.Scan(&d, &m, &n)

	stops := make([]int, n)
	for i := range stops {
		fmt.Scan(&stops[i])
	}

	fmt.Println(CarFueling(d, m, stops))
}
