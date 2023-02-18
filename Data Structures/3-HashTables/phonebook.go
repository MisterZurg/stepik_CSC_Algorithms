package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type PhoneBook map[string]string

func (pb *PhoneBook) Add(number, name string) {
	(*pb)[number] = name
}

func (pb *PhoneBook) Delete(number string) {
	delete(*pb, number)
}

func (pb *PhoneBook) Find(number string) string {
	if val, ok := (*pb)[number]; ok {
		return val
	}
	return "not found"
}

func main() {
	myPB := PhoneBook{}
	var commandsNumber int
	fmt.Scan(&commandsNumber)
	Solution(&myPB, commandsNumber)
}

func Solution(pb *PhoneBook, num int) {
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < num; i++ {
		scanner.Scan()
		line := scanner.Text()
		words := strings.Split(line, " ")
		WhatToDo(pb, words)
	}
}

func WhatToDo(pb *PhoneBook, words []string) {
	switch words[0] {
	case "add":
		pb.Add(words[1], words[2])
	case "del":
		pb.Delete(words[1])
	case "find":
		fmt.Println(pb.Find(words[1]))
	default:
		panic("wtf is a kilometer")
	}
}
