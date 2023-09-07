package main

import "fmt"

func PartyPlanningAtWork(root string, tree map[string][]string) []string {
	var answer []string
	bt := BuildBackwardsTree(tree)
	visited := make(map[string]bool)
	// DFS
	var traverse func(string)
	traverse = func(node string) {
		if _, ok := visited[node]; ok {
			return
		}
		visited[node] = true

		if len(tree[node]) == 0 {
			answer = append(answer, node)
			// Remove current leaf
			delete(tree, node)
			// Check children of parent
			for _, adjNode := range tree[bt[node]] {
				traverse(adjNode)
			}
			// Remove parent
			delete(tree, bt[node])
		}

		for _, adjNode := range tree[node] {
			traverse(adjNode)
		}
	}
	for _, adjNode := range tree[root] {
		traverse(adjNode)
	}
	if len(tree) == 1 {
		answer = append(answer, root)
	}

	return answer
}

func BuildBackwardsTree(tree map[string][]string) map[string]string {
	backwardsTree := make(map[string]string)
	for k := range tree {
		for _, node := range tree[k] {
			backwardsTree[node] = k
		}
	}
	return backwardsTree
}

func main() {
	company := map[string][]string{
		"K": {"J", "I", "H"},
		"J": {"G"},
		"I": {"F"},
		"H": {"D", "E"},
		"D": {"A", "B", "C"},
	}

	fmt.Println(PartyPlanningAtWork("K", company))
}
