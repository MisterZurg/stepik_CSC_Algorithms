package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var p = 1000000007
var x = 263
var reader = bufio.NewReader(os.Stdin)

func readLine() string {
	line, _ := reader.ReadString('\n')
	return strings.TrimRight(line, "\n")
}

func calculateHash(val string, powersOfX []int) int {
	sum := 0
	for idx, char := range val {
		sum = ((sum + (int(char) * powersOfX[idx])) % p) % p
	}
	return sum
}

func main() {
	pattern := readLine()
	text := readLine()
	patternLength := len(pattern)
	textLength := len(text)
	var powersOfX = make([]int, patternLength)
	powersOfX[0] = 1
	for i := 1; i < patternLength; i++ {
		powersOfX[i] = powersOfX[i-1] * x % p
	}
	indexes := make([]int, 0)
	patternHash := calculateHash(pattern, powersOfX)
	hash := calculateHash(text[textLength-patternLength:textLength], powersOfX)
	if hash == patternHash {
		indexes = append(indexes, textLength-patternLength)
	}
	for i := textLength - patternLength - 1; i >= 0; i-- {
		hash = ((hash-int(text[i+patternLength])*powersOfX[patternLength-1]%p+p)%p*x%p + int(text[i])) % p
		if hash == patternHash {
			indexes = append(indexes, i)
		}
	}
	for i := len(indexes) - 1; i >= 0; i-- {
		fmt.Print(indexes[i], " ")
	}
}
