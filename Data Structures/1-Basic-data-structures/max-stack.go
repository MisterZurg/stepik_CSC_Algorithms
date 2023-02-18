package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	parseInputForMaxStack()
}

func parseInputForMaxStack() {
	var n int
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	maxSt := NewMaxStack()
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()

		words := strings.Split(line, " ")
		switch words[0] {
		case "push":
			val, err := strconv.Atoi(words[1])
			if err != nil {
				panic("value is broken")
			}
			maxSt.Push(val)
		case "pop":
			maxSt.Pop()
		case "max":
			fmt.Println(maxSt.getCurrentMax())
		}
	}
}

type Stack []int

type MaxStack struct {
	elems Stack
	maxs  Stack
}

func NewMaxStack() *MaxStack {
	return new(MaxStack)
}

func (s *Stack) Push(value int) {
	*s = append(*s, value)
}

func (s *Stack) Pop() int {
	poped := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return poped
}

func (mx *MaxStack) Push(value int) {
	if len(mx.elems) == 0 {
		mx.elems.Push(value)
		mx.maxs.Push(value)
		return
	}

	if value > mx.getCurrentMax() {
		mx.maxs.Push(value)
	} else {
		mx.maxs.Push(mx.getCurrentMax())
	}
	mx.elems.Push(value)
}

func (mx *MaxStack) Pop() (int, int) {
	return mx.elems.Pop(), mx.maxs.Pop()
}

func (mx *MaxStack) getCurrentMax() int {
	return (*mx).maxs[len((*mx).maxs)-1]
}
