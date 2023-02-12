package main

import "fmt"

func main() {
	// put your code here
	var n int
	fmt.Scan(&n)

	invertedNumbers := make([]int, n)
	for i := range invertedNumbers {
		fmt.Scan(&invertedNumbers[i])
	}

	left := 0
	right :=  len(invertedNumbers) - 1
	inversionNumber := 0
	_, inv := MergeSort(invertedNumbers, left, right, inversionNumber)
	fmt.Println(inv)
}

func MergeSort(slice []int, left, right, inversionNumber int) ([]int, int) {
	// fmt.Println(slice)
	if len(slice) == 1 {
		return slice, 0
	}
	m := len(slice) / 2
	seqLeft, invA  := MergeSort(slice[:m], left, m, inversionNumber)
	seqRight, invB := MergeSort(slice[m:], m + 1 , right, inversionNumber)
	seq, inv := Merge(seqLeft, seqRight, invA + invB)
	// fmt.Println(invA + invB + inv)
	return seq, inv
}

// Merge сливает два упорядоченных слайса в один
func Merge(leftSlice, rightSlice []int, inversionNumber int) ([]int, int) {
	// Заведем указатели на первые индексы каждого из слайсов
	lp, rp := 0, 0
	i := 0
	// Заведем результирующий слайс равный сумме размеров слайсов
	result := make([]int, len(leftSlice)+len(rightSlice))

	lSize := len(leftSlice)
	rSize := len(rightSlice)
	// В первом цикле будем сравнивать элементы из обоих слайсов
	for lp < lSize && rp < rSize {
		if leftSlice[lp] <= rightSlice[rp] {
			result[i] = leftSlice[lp]
			lp++
			i++
		} else {
			result[i] = rightSlice[rp]
			// Когда берем число справа, в счетчик инверсий добавляем оставшееся число элементов слева.
			inversionNumber += lSize - lp
			rp++
			i++
		}
	}

	for lp < lSize {
		result[i] = leftSlice[lp]
		lp++
		i++
	}

	for rp < rSize {
		result[i] = rightSlice[rp]
		rp++
		i++
	}
	return result, inversionNumber
}

/*
6
10 8 6 2 4 5
Output 12

8
14 8 2 4 3 9 0 11
Output  16


11
1 2 3 4 5 6 7 8 3 4 3
Output  15


44
42 42 70 79 29 85 47 81 96 3 25 52 49 84 95 17 49 7 9 76 77 79 32 45 52 30 92 85 17 32 92 2 14 85 85 21 30 86 89 79 39 49 47 17
Output  466
*/