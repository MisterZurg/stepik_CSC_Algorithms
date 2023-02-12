package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getFrequency(letters string) map[string]int {
	frequency := make(map[string]int)
	for _, letter := range letters {
		if _, ok := frequency[string(letter)]; !ok {
			frequency[string(letter)] = 0
		}
		frequency[string(letter)]++
	}
	return frequency
}

type Pair struct {
	Letter string
	Frequency int
}

func sortByFrequency(pairs []Pair) {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Frequency < pairs[j].Frequency
	})
}

func getSortedPairs(sequense map[string]int) []Pair {
	pairs := make([]Pair, len(sequense))
	idx := 0
	for k, v := range sequense {
		pairs[idx].Letter = k
		pairs[idx].Frequency = v
		idx++
	}

	// sortByFrequencyDesc(pairs)
	// fmt.Println(pairs)
	return pairs
}

type HuffmanCode struct {
	Letter string
	Code string
}

func Huffman(sentense string, sequense map[string]int) (int, int, map[string]string, string){
	sortedPairs := getSortedPairs(sequense)
	sortByFrequency(sortedPairs)
	// fmt.Println(sortedPairs)
	tree := BuildTree(sortedPairs)
	// fmt.Println(tree)

	// Получаем буквы из строки
	letters := getSortedPairs(sequense)

	codes := make([]HuffmanCode, len(letters))

	for i, letter := range letters {
		codes[i].Letter = letter.Letter
		for j := len(tree) - 1; j>=0; j-- {
			if strings.Contains(tree[j].Letter, letter.Letter) {
				codes[i].Code += strconv.Itoa(tree[j].Frequency)
			}
		}
	}


	answer := make(map[string]string)
	for i := range codes {
		answer[codes[i].Letter] = codes[i].Code
	}

	encodedString := buildEncodedString(sentense, answer)

	return len(letters), len(encodedString), answer, encodedString
}

func BuildTree(pairsSort []Pair) []Pair {
	pairsHuffmanCodes := make([]Pair, 0)
	// Берем из pairsSort два последних по приоритету элемента
	// пока в нем не останется последний элемент.
	for len(pairsSort) > 1 {
		sortByFrequency(pairsSort)
		first1 := pairsSort[0]
		first2 := pairsSort[1]

		f1 := Pair {
			Letter: first1.Letter,
			Frequency:  0,
		}
		f2 := Pair{
			Letter: first2.Letter,
			Frequency:  1,
		}

		pairsHuffmanCodes = append(pairsHuffmanCodes, f1, f2)

		// Уменьшаем размер pairsSort
		pairsSort = pairsSort[2:]
		// Но не забываем занести комбинацию last1 и last2
		comb := Pair{
			Letter: first1.Letter + first2.Letter,
			Frequency: first1.Frequency + first2.Frequency,
		}
		pairsSort = append(pairsSort, comb)
		// fmt.Println(pairsSort)
	}
	// fmt.Println(pairsHuffmanCodes)
	return pairsHuffmanCodes
}

func buildEncodedString(sentense string, codes map[string]string) string {
	answer := ""
	for _, letter := range sentense {
		answer += codes[string(letter)]
	}
	return answer
}

func main() {
	// put your code here
	var letters string
	// fmt.Scan(&letters)
	letters = "abacabad"
	// Получаем частоты для каждой из букв в строке
	Solution(letters)
	// // fmt.Println("~~~~~~~")
	// Solution("a")
	// fmt.Println("~~~~~~~")
	// Solution("aa")
	// fmt.Println("~~~~~~~")
	// Solution("aaaa")
	//fmt.Println("~~~~~~~")
	//Solution("aaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
}

func Solution(sentence string) {
	// put your code here
	// var letters string
	// fmt.Scan(&letters)
	// letters = "abacabad"
	// Получаем частоты для каждой из букв в строке
	if len(sentence) == 1{
		fmt.Println(1, 1)
		fmt.Printf("%s: %s\n",sentence , "0")
		fmt.Println(0)
		return
	}

	if len(sentence) == strings.Count(sentence, string(sentence[0])) {
		fmt.Println(1, len(sentence))
		fmt.Printf("%s: %s\n", string(sentence[0]) , "0")
		fmt.Printf("%0*d", len(sentence), 0)
		return
	}

	toBeHaffmaned := getFrequency(sentence)
	num_letters, lengthOfString, letterCode, encodedString := Huffman(sentence, toBeHaffmaned)

	fmt.Println(num_letters,lengthOfString)
	for k, v := range letterCode {
		fmt.Printf("%s: %s\n",k ,v)
	}
	fmt.Println(encodedString)
}

// accepted
// 11110100011001100010
// p: 110
// a: 111
// c: 10
// t: 011
// d: 010
// e: 00