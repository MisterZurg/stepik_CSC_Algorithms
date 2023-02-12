package main

import (
	"bufio"
	"fmt"
	"os"
)

type Card struct {
	val  int
	prev int
	idx  int
}

type CardPile []Card

var n int
var activePiles int
var piles []*CardPile
var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)
	n = ReadInt()
	// create piles
	piles = make([]*CardPile, n)

	for i := range piles {
		piles[i] = new(CardPile)
		p := make(CardPile, 0)
		piles[i] = &p
	}

	for i := 0; i < n; i++ {
		c := new(Card)
		c.val = ReadInt()
		c.idx = i
		// find appropriate pile
		var pidx int
		pidx = findPile2(c.val)
		if pidx == 0 {
			c.prev = -1
		} else {
			c.prev = len(*piles[pidx-1]) - 1
		}
		*piles[pidx] = append(*piles[pidx], *c)
	}
	fmt.Println(activePiles)

	answer := make([]int, activePiles)
	currentPile := activePiles - 1
	ccard := (*piles[currentPile])[(len(*piles[currentPile]) - 1)]
	for {
		answer[currentPile] = ccard.idx + 1
		if ccard.prev == -1 {
			break
		}
		currentPile--
		ccard = (*piles[currentPile])[ccard.prev]
	}
	for _, ans := range answer {
		fmt.Printf("%d ", ans)
	}
	fmt.Println()
}

func findPile2(c int) (index int) {

	if activePiles == 0 {
		activePiles++
		return activePiles - 1
	}
	l := 0
	h := activePiles
	var mid int
	for {
		if l >= h {
			break
		}
		mid = l + (h-l)/2
		val := (*piles[mid])[len(*piles[mid])-1].val
		if c > val {
			h = mid// - 1
		} else {
			l = mid + 1
		}
	}
	if l == activePiles {
		activePiles++
	}
	return l
}

/*------------ fast stdin read routines ------------------------- */
func ReadInt() int {
	minus := false
	result := 0
	ch := getchar()
	for {
		if ch == '-' {
			break
		}
		if ch >= '0' && ch <= '9' {
			break
		}
		ch = getchar()
	}
	if ch == '-' {
		minus = true
	} else {
		result = ch - '0'
	}
	for {
		ch = getchar()
		if ch < '0' || ch > '9' {
			break
		}
		result = result*10 + (ch - '0')
	}
	if minus {
		return -result
	} else {
		return result
	}
}

func getchar() int {
	for {
		b, err := reader.ReadByte()
		if err != nil {
			//break
			panic("getchar error")
		}
		return int(b)
	}
}