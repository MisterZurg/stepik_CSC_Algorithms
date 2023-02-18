package main

import (
	"fmt"
	"strconv"
)

type Bracket struct {
	position int
	brType   string
}

type BracketStack []Bracket

// IsEmpty: check if stack is empty
func (bs *BracketStack) IsEmpty() bool {
	return len(*bs) == 0
}

// Push a new value onto the stack
func (bs *BracketStack) Push(br Bracket) {
	*bs = append(*bs, br) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (bs *BracketStack) Pop() (Bracket, bool) {
	if bs.IsEmpty() {
		return Bracket{}, false
	} else {
		index := len(*bs) - 1   // Get the index of the top most element.
		element := (*bs)[index] // Index into the slice and obtain the element.
		*bs = (*bs)[:index]     // Remove it from the stack by slicing it off.
		return element, true
	}
}

func Solution(line string) string {
	bracketsStack := &BracketStack{}

	for i, symbol := range line {
		if !isBracket(string(symbol)) {
			// fmt.Println("SKIP")
			continue
		}
		if isOpeningBracket(string(symbol)) {
			bracketsStack.Push(Bracket{
				position: i + 1,
				brType:   string(symbol),
			})
		} else {
			if bracketsStack.IsEmpty() {
				// Because we have 2 return No of bracket)
				return strconv.Itoa(i + 1)
			}
			top, _ := bracketsStack.Pop()

			if areBracketsMatched(top.brType, string(symbol)) {
				// Because we have 2 return No of bracket
				return strconv.Itoa(i + 1)
			}
		}
	}

	if !bracketsStack.IsEmpty() {
		top, _ := bracketsStack.Pop()
		// fmt.Println(top.position)
		return strconv.Itoa(top.position)
	}
	return "Success"
}

func isBracket(symbol string) bool {
	return symbol == "(" || symbol == ")" ||
		symbol == "[" || symbol == "]" ||
		symbol == "{" || symbol == "}"
}

func isOpeningBracket(symbol string) bool {
	return string(symbol) == "(" || string(symbol) == "[" || string(symbol) == "{"
}

func areBracketsMatched(topStack, currSymbol string) bool {
	return (topStack == "(" && currSymbol != ")") ||
		(topStack == "[" && currSymbol != "]") ||
		(topStack == "{" && currSymbol != "}")
}

func main() {
	// put your code here
	var bracketsLine string
	fmt.Scan(&bracketsLine)
	fmt.Println(Solution(bracketsLine))
	//TestCases()
}

// ([](){([])}) : Success
// ()[]} : 5
// {{[()]] : 7

func TestCases() {
	// I'd better write table tests
	fmt.Println(Solution("[]") == "Success")
	fmt.Println(Solution("{}[]") == "Success")
	fmt.Println(Solution("[()]") == "Success")
	fmt.Println(Solution("(())") == "Success")
	fmt.Println(Solution("{[]}()") == "Success")
	fmt.Println()
	fmt.Println(Solution("{") == "1")
	fmt.Println(Solution("{[}") == "3")
	fmt.Println()
	fmt.Println(Solution("foo(bar);") == "Success")
	fmt.Println(Solution("foo(bar[i);") == "10")
	fmt.Println()
	fmt.Println(Solution("([](){([])})") == "Success")
	fmt.Println(Solution("()[]}") == "5")
	fmt.Println(Solution("{{[()]]") == "7")
	fmt.Println(Solution("{{{[][][]"))
}
