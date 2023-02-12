package main

import (
	"fmt"
	"strings"
)

func main(){
	_, length, codes, sequence := parseInput()
	fmt.Println(length, codes, sequence)
	if length == 1 {
		fmt.Println(codes[sequence])
		return
	}

	i, j := 0, 1
	var answer string
    for i < length && j < length + 1 {
		// fmt.Println(sequence[j])
		if _, ok := codes[string(sequence[i:j])]; ok {
			answer += codes[string(sequence[i:j])]
			i = j
		}
		j++
	}
	fmt.Println(answer)
}

func parseInput() (int, int, map[string]string, string){
	var numberOfLetters, length int
	fmt.Scan(&numberOfLetters, &length)


	letters := make(map[string]string)
	for i:= 0; i < numberOfLetters;i++ {
		var letter, code string
		fmt.Scan(&letter, &code)
		letter = strings.TrimRight(letter, ":")
		letters[code] = letter
	}
	var seq string
	fmt.Scan(&seq)
	return numberOfLetters, length, letters, seq
}

/*
Sample Input 1:

1 1
a: 0
0
Sample Output 1:

a
Sample Input 2:

4 14
a: 0
b: 10
c: 110
d: 111
01001100100111
Sample Output 2:

abacabad
*/