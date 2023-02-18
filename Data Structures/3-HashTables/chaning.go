package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	X = 263
	P = 1_000_000_007
)

type hashTableElem struct {
	data string
	next *hashTableElem
}

type HashTable []*hashTableElem

func (ht HashTable) hash(data string) int {
	strBytes := []byte(data)
	hashSum := 0
	x := 1
	for _, char := range strBytes {
		hashSum += int(char) * x
		hashSum = (hashSum%P + P) % P
		x = (x * X) % P
	}
	return hashSum % len(ht)
}

func (ht HashTable) Find(str string) bool {
	hashSum := ht.hash(str)
	for elem := ht[hashSum]; elem != nil; elem = elem.next {
		if elem.data == str {
			return true
		}
	}
	return false
}

func (ht HashTable) Add(elem string) {
	if ht.Find(elem) {
		return
	}

	hashSum := ht.hash(elem)
	var temp *hashTableElem
	if ht[hashSum] == nil {
		temp = &hashTableElem{
			data: elem,
			next: nil,
		}
	} else {
		temp = &hashTableElem{
			data: elem,
			next: ht[hashSum],
		}
	}
	ht[hashSum] = temp
}

func (ht HashTable) Delete(str string) {
	hashSum := ht.hash(str)
	if ht[hashSum] == nil {
		return
	}
	if ht[hashSum].data == str {
		ht[hashSum] = ht[hashSum].next
		return
	}
	prev := ht[hashSum]
	for elem := ht[hashSum].next; elem != nil; prev, elem = elem, elem.next {
		if elem.data == str {
			prev.next = elem.next
			return
		}
	}
}

func (ht HashTable) Check(i int) *hashTableElem {
	return ht[i]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	hashTable := make(HashTable, m)
	for i := 0; i < n; i++ {
		scanner.Scan()
		command := scanner.Text()
		scanner.Scan()
		switch command {
		case "add":
			hashTable.Add(scanner.Text())
		case "del":
			hashTable.Delete((scanner.Text()))
		case "find":
			if hashTable.Find(scanner.Text()) {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		case "check":
			i, _ := strconv.Atoi(scanner.Text())
			for elem := hashTable.Check(i); elem != nil; elem = elem.next {
				fmt.Print(elem.data, " ")
			}
			fmt.Println()
		}
	}
}
