package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	key   int
	left  int // usually *Node
	right int // usually *Node
}

type SearchTree struct {
	Nodes  []TreeNode
	Output []int
}

func NewSearchTree(size int) SearchTree {
	nodes := make([]TreeNode, size)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan() && i < size; i++ {
		var key, left, right int
		line := scanner.Text()
		// strings.Split(s, " ")
		txt := strings.Split(line, " ")
		key, _ = strconv.Atoi(txt[0])
		left, _ = strconv.Atoi(txt[1])
		right, _ = strconv.Atoi(txt[2])

		node := TreeNode{key, left, right}
		nodes[i] = node
	}
	return SearchTree{Nodes: nodes}
}

func (st *SearchTree) InOrderTraversal(idx int) {
	if idx == -1 {
		return
	}
	st.InOrderTraversal((*st).Nodes[idx].left)
	st.Output = append(st.Output, st.Nodes[idx].key)
	st.InOrderTraversal((*st).Nodes[idx].right)
}

func (st *SearchTree) CheckProperty() bool {
	st.InOrderTraversal(0)
	prev := st.Output[0]
	for i := 1; i < len(st.Output); i++ {
		curr := st.Output[i]

		if !(curr > prev) {
			return false
		}
		prev = curr
	}
	return true
}

func Solution() {
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("CORRECT")
		return
	}

	tree := NewSearchTree(n)

	switch tree.CheckProperty() {
	case true:
		fmt.Println("CORRECT")
	case false:
		fmt.Println("INCORRECT")
	}
}

func main() {
	Solution()
}
