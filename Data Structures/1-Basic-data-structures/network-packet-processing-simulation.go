package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	arrival  int
	duration int
	end      int
}

func (n *Node) updateEnd(node Node) {
	if n.arrival < node.end {
		n.end = node.end + n.duration
	} else {
		n.end = n.arrival + n.duration
	}
	fmt.Println(n.end - n.duration)
}

type Stack struct {
	stack []Node
}

func (s *Stack) push(node Node) {
	s.stack = append(s.stack, node)
}

func (s *Stack) pop() Node {
	elem := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return elem
}

type Queue struct {
	pushStack Stack
	popStack  Stack
}

func (q *Queue) isEmptyPopStack() bool {
	if len(q.popStack.stack) == 0 {
		return true
	}
	return false
}

func (q *Queue) getFirstElement() Node {
	if q.isEmptyPopStack() {
		return q.pushStack.stack[0]
	}
	return q.popStack.stack[len(q.popStack.stack)-1]
}

func (q *Queue) push(node Node) {
	q.pushStack.push(node)
}

func (q *Queue) pop() Node {
	if q.isEmptyPopStack() {
		for i := len(q.pushStack.stack) - 1; i >= 0; i-- {
			node := q.pushStack.stack[i]
			q.popStack.push(node)
		}
		q.pushStack.stack = []Node{}
	}

	return q.popStack.pop()
}

func main() {
	var size, count int
	var arr, dur int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	split := strings.Split(scanner.Text(), " ")
	size, _ = strconv.Atoi(split[0])
	count, _ = strconv.Atoi(split[1])

	var queue = Queue{pushStack: Stack{stack: []Node{}}, popStack: Stack{stack: []Node{}}}

	for i := 0; i < count; i++ {
		scanner.Scan()
		split := strings.Split(scanner.Text(), " ")
		arr, _ = strconv.Atoi(split[0])
		dur, _ = strconv.Atoi(split[1])
		node := Node{arrival: arr, duration: dur}

		if len(queue.popStack.stack)+len(queue.pushStack.stack) == 0 {
			node.updateEnd(Node{})
			queue.push(node)
		} else {
			prevNode := queue.pushStack.stack[len(queue.pushStack.stack)-1]
			if len(queue.popStack.stack)+len(queue.pushStack.stack) < size {
				node.updateEnd(prevNode)
				queue.push(node)
			} else {
				firtsNode := queue.getFirstElement()
				if node.arrival < firtsNode.end {
					fmt.Println(-1)
				} else {
					node.updateEnd(prevNode)
					queue.pop()
					queue.push(node)
				}
			}
		}
	}
}
