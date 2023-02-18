package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
	root *Node

	prevKey    *rune
	correct    bool
	inOrderSum int
}

type Node struct {
	key    rune
	parent *Node
	left   *Node
	right  *Node
	count  int
}

func (t *Tree) Merge(v2 *Tree) {
	if v2.root == nil {
		return
	}

	t.splayMax()

	if t.root == nil {
		t.root = v2.root
		v2.root = nil
		return
	}

	t.root.right = v2.root
	v2.root.parent = t.root
	t.updateCount(t.root)
	v2.root = nil
}

func (t *Tree) Cut(l, r int) *Tree {
	if t.root == nil {
		return &Tree{}
	}

	leftTree, rightTree := t.split(l - 1)

	if rightTree.root == nil {
		t.Merge(leftTree)
		t.Merge(rightTree)
		return &Tree{}
	}
	cutedCount := 0
	if leftTree.root != nil {
		cutedCount = leftTree.root.count
	}

	leftOfRightTree, rightOfRightTree := rightTree.split(r - cutedCount)

	if leftOfRightTree.root == nil {
		rightTree.Merge(leftOfRightTree)
		rightTree.Merge(rightOfRightTree)

		t.Merge(leftTree)
		t.Merge(rightTree)
		return &Tree{}
	}

	t.Merge(leftTree)
	t.Merge(rightOfRightTree)

	return leftOfRightTree
}

func (t *Tree) Paste(cuted *Tree, k int) {
	if t.root == nil {
		t.Merge(cuted)
		return
	}

	i := k - 1
	leftTree, rightTree := t.split(i)
	leftTree.Merge(cuted)
	leftTree.Merge(rightTree)
	t.Merge(leftTree)
}

func (tr *Tree) Print() {
	tr.printInOrder(tr.root)
}

func (t *Tree) split(i int) (*Tree, *Tree) {
	if t.root == nil {
		return &Tree{}, &Tree{}
	}

	var left *Node
	var right *Node

	t.search(i, t.root)

	currentIndex := t.currentIndex(t.root)

	if currentIndex <= i {
		left = t.root
		right = t.root.right
		if right != nil {
			left.right = nil
			right.parent = nil
		}
		t.updateCount(left)
		t.root = nil
	} else {
		left = t.root.left
		right = t.root
		if left != nil {
			left.parent = nil
			right.left = nil
		}
		t.updateCount(right)
		t.root = nil
	}

	return &Tree{root: left}, &Tree{root: right}
}

func (t *Tree) search(i int, n *Node) bool {
	if n == nil {
		return false
	}

	return t.doSearch(i, n)
}

func (t *Tree) doSearch(i int, n *Node) bool {
	currentIndex := t.currentIndex(n)

	if currentIndex == i {
		t.splay(n)
		return true
	}

	if i < currentIndex {
		if n.left == nil {
			t.splay(n)
			return false
		}

		return t.search(i, n.left)
	}

	if n.right == nil {
		t.splay(n)
		return false
	}

	return t.search(i-currentIndex-1, n.right)
}

func (t *Tree) currentIndex(n *Node) int {
	currentIndex := 0
	if n.left != nil {
		currentIndex = n.left.count
	}
	return currentIndex
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
	t.updateCount(a)

	b.left = u.right
	if u.right != nil {
		u.right.parent = b
	}
	t.updateCount(b)

	u.left = a
	a.parent = u

	u.right = b
	b.parent = u

	t.updateCount(u)
}

func (t *Tree) zagZig(u *Node) {
	b := u.parent
	a := b.parent

	t.linkXWithParentOfY(u, a)

	b.right = u.left
	if u.left != nil {
		u.left.parent = b
	}
	t.updateCount(b)

	a.left = u.right
	if u.right != nil {
		u.right.parent = a
	}
	t.updateCount(a)

	u.left = b
	b.parent = u

	u.right = a
	a.parent = u

	t.updateCount(u)
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
	t.updateCount(y)

	x.right = y
	y.parent = x
	t.updateCount(x)
}

func (t *Tree) smallLeftTurn(x, y *Node) {
	y.right = x.left
	if x.left != nil {
		x.left.parent = y
	}
	t.updateCount(y)

	x.left = y
	y.parent = x
	t.updateCount(x)
}

func (t *Tree) updateCount(n *Node) {
	if n == nil {
		return
	}

	count := 1
	if n.left != nil {
		count += n.left.count
	}

	if n.right != nil {
		count += n.right.count
	}

	n.count = count
}

func (tr *Tree) printInOrder(node *Node) {
	if node == nil {
		return
	}
	tr.printInOrder(node.left)
	fmt.Printf("%c", node.key)
	tr.printInOrder(node.right)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')
	s = strings.Trim(s, "\n")

	sTree := &Tree{}
	for _, v := range s {
		sTree.Merge(&Tree{root: &Node{key: v, count: 1}})
	}

	nStr, _ := reader.ReadString('\n')
	nStr = strings.Trim(nStr, "\n")
	n, _ := strconv.Atoi(nStr)

	for i := 0; i < n; i++ {
		opStr, _ := reader.ReadString('\n')
		op := strings.Fields(opStr)

		l, _ := strconv.Atoi(op[0])
		r, _ := strconv.Atoi(op[1])
		k, _ := strconv.Atoi(op[2])

		cuted := sTree.Cut(l, r)

		sTree.Paste(cuted, k)

	}

	sTree.Print()
}
