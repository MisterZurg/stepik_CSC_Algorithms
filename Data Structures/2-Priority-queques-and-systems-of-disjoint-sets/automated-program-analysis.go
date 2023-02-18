package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DJS struct {
	parent []int
	rank   []int
}

func (s *DJS) MakeSet(number int) {
	s.parent[number] = number
	s.rank[number] = 0
}

func (s *DJS) Find(number int) int {
	if number != s.parent[number] {
		number = s.Find(s.parent[number])
	}
	return number
}

func (s *DJS) Union(i int, j int) {
	i_id := s.Find(i)
	j_id := s.Find(j)

	if i_id == j_id {
		return
	}

	if s.rank[i_id] > s.rank[j_id] {
		s.parent[j_id] = i_id
	} else {
		s.parent[i_id] = j_id
		if s.rank[i_id] == s.rank[j_id] {
			s.rank[j_id]++
		}
	}
}

func NewDJS(variable int) DJS {
	sets := DJS{
		make([]int, variable+1),
		make([]int, variable+1),
	}

	for i := 1; i <= variable; i++ {
		sets.MakeSet(i)
	}
	return sets
}

func main() {
	result := 1

	reader := bufio.NewReader(os.Stdin)
	nedStr, _ := reader.ReadString('\n')
	ned := strings.Fields(nedStr)
	n, _ := strconv.Atoi(ned[0])
	e, _ := strconv.Atoi(ned[1])
	d, _ := strconv.Atoi(ned[2])

	sets := NewDJS(n)

	for ; e > 0; e-- {
		eStr, _ := reader.ReadString('\n')
		e := strings.Fields(eStr)
		i, _ := strconv.Atoi(e[0])
		j, _ := strconv.Atoi(e[1])

		sets.Union(i, j)
	}

	for ; d > 0; d-- {
		dStr, _ := reader.ReadString('\n')
		d := strings.Fields(dStr)
		i, _ := strconv.Atoi(d[0])
		j, _ := strconv.Atoi(d[1])

		i_id := sets.Find(i)
		j_id := sets.Find(j)

		if i_id == j_id {
			result = 0
			break
		}
	}

	fmt.Println(result)
}
