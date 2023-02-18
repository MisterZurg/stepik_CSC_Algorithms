package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val   int
	left  int // usually *Node
	right int // usually *Node
}

type Tree []Node

// FastReadTree builds tree from the input
// Max time used: 0.38/1.50
func NewTree() *Tree {
	var n int
	fmt.Scan(&n)
	newTree := make(Tree, n)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan() && i < n; i++ {
		var key, left, right int
		line := scanner.Text()
		// strings.Split(s, " ")
		txt := strings.Split(line, " ")
		key, _ = strconv.Atoi(txt[0])
		left, _ = strconv.Atoi(txt[1])
		right, _ = strconv.Atoi(txt[2])

		node := Node{
			key, left, right,
		}
		newTree[i] = node
	}
	return &newTree
}

func (tr *Tree) InOrderTraversal(idx int) {
	if idx == -1 {
		return
	}
	tr.InOrderTraversal((*tr)[idx].left)
	fmt.Print((*tr)[idx].val, " ")
	tr.InOrderTraversal((*tr)[idx].right)
}

func (tr *Tree) PreOrderTraversal(idx int) {
	if idx == -1 {
		return
	}
	fmt.Print((*tr)[idx].val, " ")
	tr.InOrderTraversal((*tr)[idx].left)
	tr.InOrderTraversal((*tr)[idx].right)
}

func (tr *Tree) PostOrderTraversal(idx int) {
	if idx == -1 {
		return
	}
	tr.PostOrderTraversal((*tr)[idx].left)
	tr.PostOrderTraversal((*tr)[idx].right)
	fmt.Print((*tr)[idx].val, " ")
}

func main() {
	tree := NewTree()
	tree.InOrderTraversal(0)
	fmt.Println()
	tree.PreOrderTraversal(0)
	fmt.Println()
	tree.PostOrderTraversal(0)
}
