package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	Tree MyTree
)

type MyTree struct {
	nodes   []*Node
	prevKey *int
	correct bool
}

func (tr *MyTree) addNode(index int, node *Node) {
	tr.nodes[index] = node
}

func (tr *MyTree) getNode(index int) *Node {
	if index == -1 || index >= len(tr.nodes) {
		return nil
	}

	return tr.nodes[index]
}

func (tr *MyTree) check() {
	Tree.correct = true

	Tree.do(tr.getNode(0))

	if Tree.correct {
		fmt.Println("CORRECT")
	} else {
		fmt.Println("INCORRECT")
	}
}

func (tr *MyTree) do(node *Node) {
	if node == nil {
		return
	}

	left := Tree.getNode(node.leftIndex)
	right := Tree.getNode(node.rightIndex)

	if left != nil {
		if !(tr.max(left) < node.key) {
			tr.correct = false
			return
		}

		Tree.do(left)
	}

	if right != nil {
		if !(node.key <= tr.min(right)) {
			tr.correct = false
			return
		}

		Tree.do(right)
	}
}

func (tr *MyTree) max(node *Node) int {
	right := Tree.getNode(node.rightIndex)
	if right == nil {
		return node.key
	}

	return tr.max(right)
}

func (tr *MyTree) min(node *Node) int {
	left := Tree.getNode(node.leftIndex)
	if left == nil {
		return node.key
	}

	return tr.min(left)
}

type Node struct {
	key        int
	leftIndex  int
	rightIndex int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	nStr, _ := reader.ReadString('\n')
	nStr = strings.Trim(nStr, "\n")
	n, _ := strconv.Atoi(nStr)

	Tree = MyTree{nodes: make([]*Node, n, n)}

	for i := 0; i < n; i++ {
		keyLeftRightStr, _ := reader.ReadString('\n')
		keyLeftRight := strings.Fields(keyLeftRightStr)

		key, _ := strconv.Atoi(keyLeftRight[0])
		leftIndex, _ := strconv.Atoi(keyLeftRight[1])
		rightIndex, _ := strconv.Atoi(keyLeftRight[2])

		Tree.addNode(i, &Node{key, leftIndex, rightIndex})
	}

	Tree.check()
}
