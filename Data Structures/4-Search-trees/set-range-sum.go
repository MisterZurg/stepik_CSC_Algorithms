package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
	root    *Node
	lastSum int

	prevKey    *int
	correct    bool
	inOrderSum int
}

type Node struct {
	key    int
	parent *Node
	left   *Node
	right  *Node
	sum    int
}

func (t *Tree) Add(key int) {
	if t.root == nil {
		t.root = &Node{key: key}
		return
	}

	t.add(key, t.root)
}

func (t *Tree) add(key int, n *Node) {
	if n.key == key {
		return
	}

	if key < n.key {
		if n.left == nil {
			n.left = &Node{key: key, parent: n, sum: key}
			t.splay(n.left)
			return
		}

		t.add(key, n.left)
		return
	}

	if n.right == nil {
		n.right = &Node{key: key, parent: n, sum: key}
		t.splay(n.right)
		return
	}

	t.add(key, n.right)
}

func (t *Tree) Search(key int) bool {
	if t.root == nil {
		return false
	}

	return t.search(key, t.root)
}

func (t *Tree) search(key int, n *Node) bool {
	if n.key == key {
		t.splay(n)
		return true
	}

	if key < n.key {
		if n.left == nil {
			t.splay(n)
			return false
		}

		return t.search(key, n.left)
	}

	if n.right == nil {
		t.splay(n)
		return false
	}

	return t.search(key, n.right)
}

func (t *Tree) Remove(key int) {
	isFound := t.Search(key)

	if !isFound {
		return
	}

	// delete finded node
	finded := t.root
	t.root = nil
	if finded.left != nil {
		finded.left.parent = nil
	}
	if finded.right != nil {
		finded.right.parent = nil
	}

	left := &Tree{root: finded.left}
	right := &Tree{root: finded.right}

	t.merge(left, right)
}

func (t *Tree) Sum(l, r int) int {
	if t.root == nil {
		t.lastSum = 0
		return 0
	}

	leftTree, rightTree := t.split(l - 1)

	if rightTree.root == nil {
		t.merge(leftTree, rightTree)
		t.lastSum = 0
		return 0
	}

	leftOfRightTree, rightOfRightTree := rightTree.split(r)

	if leftOfRightTree.root == nil {
		rightTree.merge(leftOfRightTree, rightOfRightTree)
		t.merge(leftTree, rightTree)
		t.lastSum = 0
		return 0
	}

	res := leftOfRightTree.root.sum
	rightTree.merge(leftOfRightTree, rightOfRightTree)
	t.merge(leftTree, rightTree)
	t.lastSum = res
	return res
}

func (t *Tree) split(key int) (*Tree, *Tree) {

	var left *Node
	var right *Node

	t.search(key, t.root)
	if t.root.key <= key {
		left = t.root
		right = t.root.right
		if right != nil {
			left.right = nil
			right.parent = nil
		}
		t.updateSum(left)
		t.root = nil
	} else {
		left = t.root.left
		right = t.root
		if left != nil {
			left.parent = nil
			right.left = nil
		}
		t.updateSum(right)
		t.root = nil
	}

	return &Tree{root: left}, &Tree{root: right}
}

func (t *Tree) merge(v1, v2 *Tree) {
	v1.splayMax()

	if v1.root == nil {
		t.root = v2.root
		return
	}

	if v2.root == nil {
		t.root = v1.root
		return
	}

	v1.root.right = v2.root
	v2.root.parent = v1.root
	t.updateSum(v1.root)
	t.root = v1.root
}

func (t *Tree) splayMax() {
	if t.root == nil {
		return
	}

	node := t.root
	for {
		if node.right == nil {
			t.splay(node)
			return
		}

		node = node.right
	}
}

func (t *Tree) splay(u *Node) {
	if u == nil {
		return
	}

	if u.parent == nil {
		return
	}

	if u.parent.parent == nil {
		if u.parent.left == u {
			t.zig(u)
			return
		}

		t.zag(u)
		return
	}

	if u.parent.left == u && u.parent.parent.left == u.parent {
		t.zigZig(u)
		t.splay(u)
		return
	}

	if u.parent.left == u && u.parent.parent.right == u.parent {
		t.zigZag(u)
		t.splay(u)
		return
	}

	if u.parent.right == u && u.parent.parent.right == u.parent {
		t.zagZag(u)
		t.splay(u)
		return
	}

	if u.parent.right == u && u.parent.parent.left == u.parent {
		t.zagZig(u)
		t.splay(u)
		return
	}

	panic("panic")
}

func (t *Tree) zigZig(u *Node) {
	a := u.parent
	b := a.parent

	t.linkXWithParentOfY(u, b)
	t.smallRightTurn(a, b)
	t.smallRightTurn(u, a)
}

func (t *Tree) zagZag(u *Node) {
	a := u.parent
	b := a.parent

	t.linkXWithParentOfY(u, b)
	t.smallLeftTurn(a, b)
	t.smallLeftTurn(u, a)
}

func (t *Tree) zig(u *Node) {
	a := u.parent
	t.smallRightTurn(u, a)
	u.parent = nil
	t.root = u
}

func (t *Tree) zag(u *Node) {
	a := u.parent
	t.smallLeftTurn(u, a)
	u.parent = nil
	t.root = u
}

func (t *Tree) zigZag(u *Node) {
	b := u.parent
	a := b.parent

	t.linkXWithParentOfY(u, a)

	a.right = u.left
	if u.left != nil {
		u.left.parent = a
	}
	t.updateSum(a)

	b.left = u.right
	if u.right != nil {
		u.right.parent = b
	}
	t.updateSum(b)

	u.left = a
	a.parent = u

	u.right = b
	b.parent = u

	t.updateSum(u)
}

func (t *Tree) zagZig(u *Node) {
	b := u.parent
	a := b.parent

	t.linkXWithParentOfY(u, a)

	b.right = u.left
	if u.left != nil {
		u.left.parent = b
	}
	t.updateSum(b)

	a.left = u.right
	if u.right != nil {
		u.right.parent = a
	}
	t.updateSum(a)

	u.left = b
	b.parent = u

	u.right = a
	a.parent = u

	t.updateSum(u)
}

func (t *Tree) linkXWithParentOfY(x, y *Node) {
	if y.parent == nil {
		x.parent = nil
		t.root = x
	} else {
		if y.parent.left == y {
			y.parent.left = x
			x.parent = y.parent
		} else {
			y.parent.right = x
			x.parent = y.parent
		}
	}
}

func (t *Tree) smallRightTurn(x, y *Node) {
	y.left = x.right
	if x.right != nil {
		x.right.parent = y
	}
	t.updateSum(y)

	x.right = y
	y.parent = x
	t.updateSum(x)
}

func (t *Tree) smallLeftTurn(x, y *Node) {
	y.right = x.left
	if x.left != nil {
		x.left.parent = y
	}
	t.updateSum(y)

	x.left = y
	y.parent = x
	t.updateSum(x)
}

func (t *Tree) updateSum(n *Node) {
	if n == nil {
		return
	}

	sum := n.key
	if n.left != nil {
		sum += n.left.sum
	}

	if n.right != nil {
		sum += n.right.sum
	}

	n.sum = sum
}

func (tr *Tree) Check() {
	tr.correct = true
	tr.inOrderSum = 0

	tr.inOrder(tr.root)

	if tr.correct {
		fmt.Println("CORRECT", tr.inOrderSum)
	} else {
		fmt.Println("INCORRECT")
	}
}

func (tr *Tree) Print() {
	tr.printInOrder(tr.root)
}

func (tr *Tree) printInOrder(node *Node) {
	if node == nil {
		return
	}

	tr.printInOrder(node.left)

	fmt.Println(node.key, " ")

	tr.printInOrder(node.right)
}

func (tr *Tree) inOrder(node *Node) {
	if node == nil {
		return
	}

	tr.inOrder(node.left)

	if tr.prevKey != nil && node.key < *tr.prevKey {
		tr.correct = false
		return
	}
	tr.prevKey = &node.key
	tr.inOrderSum += node.key

	tr.inOrder(node.right)
}

func main() {
	tree := &Tree{}

	reader := bufio.NewReader(os.Stdin)
	nStr, _ := reader.ReadString('\n')
	nStr = strings.Trim(nStr, "\n")
	n, _ := strconv.Atoi(nStr)

	for i := 0; i < n; i++ {
		opStr, _ := reader.ReadString('\n')
		op := strings.Fields(opStr)

		if op[0] == "+" {
			i, _ := strconv.Atoi(op[1])
			tree.Add(getI(i, tree.lastSum))
			continue
		}

		if op[0] == "-" {
			i, _ := strconv.Atoi(op[1])
			tree.Remove(getI(i, tree.lastSum))
			continue
		}

		if op[0] == "?" {
			i, _ := strconv.Atoi(op[1])
			res := tree.Search(getI(i, tree.lastSum))
			if res == true {
				fmt.Println("Found")
			} else {
				fmt.Println("Not found")
			}
			continue
		}

		if op[0] == "s" {
			l, _ := strconv.Atoi(op[1])
			r, _ := strconv.Atoi(op[2])
			res := tree.Sum(getI(l, tree.lastSum), getI(r, tree.lastSum))
			fmt.Println(res)
			continue
		}
	}
}

func getI(x, s int) int {
	return (x + s) % 1000000001
}
