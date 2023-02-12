package main

import "fmt"

func main() {
	// put your code here
	W, weights := ParseInput()
	fmt.Println(MaxGoldWeight(W, weights))
}

func MaxGoldWeight(W int, weights []int) int {
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, W+1)
	}

	for j := 0; j <= W; j++ {
		if weights[0] <= j {
			m[0][j] = weights[0]
		}
	}

	// Во внешнем цикле перебираем веса предметов
	for i := 1; i < n; i++ {
		// Во внутреннем цикле перебираем capacity рюкзака
		for j := 1; j <= W; j++ {
			// Наполняем таблицу реешения, "ответами"
			m[i][j] = m[i-1][j]

			dw := j - weights[i]
			// Или максимум из предидущего значения
			if dw >= 0 && m[i-1][dw]+weights[i] > m[i-1][j] {
				m[i][j] = m[i-1][dw] + weights[i]
			}
		}
	}
	return m[n-1][W]
}

func ParseInput() (int, []int) {
	var W, n int // вместимость рюкзака, число золотых слитков соответственно
	fmt.Scan(&W, &n)

	weights := make([]int, n)
	for i := range weights {
		fmt.Scan(&weights[i])
	}

	return W, weights
}

func PrintTable(table [][]int) {
	for _, row := range table {
		fmt.Println(row)
	}

}

/*
Sample Input:
10 3
1 4 8

Sample Output:
9
 */