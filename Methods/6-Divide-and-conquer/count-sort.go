package main

import "fmt"

func main() {
	// put your code here
	var n int
	fmt.Scan(&n)
	toBeCountSorted := make([]int, n)

	for i := range toBeCountSorted {
		fmt.Scan(&toBeCountSorted[i])
	}

	CountSortYandex(toBeCountSorted)
	for i := range toBeCountSorted {
		fmt.Printf("%d ", toBeCountSorted[i])
	}
}

func CountSortYandex(seq []int){
	max := seq[0]
	min := seq[0]
	//
	for _, elem := range seq {
		if elem < min {
			min = elem
		}
		if elem > max {
			max = elem
		}
	}
	// Количество возможных значений
	k := max - min + 1
	// Создадим массив из k нулей
	count := make([]int, k)
	// Перебираем все элеметы
	for _, elem := range seq {
		count[elem - min] +=1
	}

	nowPosition := 0
	// Перебираем все элеметы
	for val := 0; val < k; val++ {
		for i := 0; i < count[val]; i++ {
			// Заменяем то что было в последовательности
			seq[nowPosition] = val + min
			nowPosition++
		}
	}
}
