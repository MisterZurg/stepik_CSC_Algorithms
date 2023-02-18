package main

import "fmt"

func main() {
	ajList := parseInput()
	root := ajList[-1][0]
	fmt.Println(treeHeight(root, ajList))
}

func parseInput() map[int][]int {
	ajList := make(map[int][]int)

	var nodes, parent int
	fmt.Scan(&nodes)

	for node := 0; node < nodes; node++ {
		fmt.Scan(&parent)
		ajList[parent] = append(ajList[parent], node)
	}
	return ajList
}

func treeHeight(root int, aj map[int][]int) int {
	height := 1
	if aj[root] == nil {
		return height
	}
	for _, child := range aj[root] {
		height = max(height, 1+treeHeight(child, aj))
	}
	return height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
