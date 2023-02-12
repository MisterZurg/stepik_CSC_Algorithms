package main

import (
	"fmt"
	"math"
)

func main() {
	// put your code here
	var n, m int // количество уравнений, количество неизвестных
	fmt.Scan(&n, &m)

	// Матрица из коэффициентов при неизвестных
	matrix := make([][]float64, n)
	// Cтолбец свободных членов
	b := make([]float64, n)

	// Инициализация nested-слайса
	for i := range matrix {
		matrix[i] = make([]float64, m)
	}

	// Счёт пользовательского инпута
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&matrix[i][j])
		}
		fmt.Scan(&b[i])
	}

	for p :=0; p < n - 1; p++ {
		max := p
		for i:= p+ 1; i < n;i++ {
			if math.Abs(matrix[i][p]) > math.Abs(matrix[max][p]) {
				max = i
			}
		}
		matrix[p], matrix[max] = matrix[max], matrix[p]
		b[p], b[max] = b[max], b[p]

		if math.Abs(matrix[p][p]) <= 1e-10 {
			fmt.Println("NO")
			return
		}

		for i := p + 1; i < n; i++ {
			alpha := matrix[i][p] / matrix[p][p]
			b[i] -= alpha * b[p]
			for j := p; j < n - 1; j++ {
				fmt.Println(matrix[i][j], "ij=", i,j ,"p=", p)
				matrix[i][j]  -= alpha * matrix[p][j] //shit
			}
		}
	}

	x := make([]float64, n)
	for i := n - 2; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < n- 1; j++ {
			sum += matrix[i][j] * x[j]
		}
		x[i] = (b[i] - sum) / matrix[i][i]
	}

	if n < m {
		fmt.Println("INF")
	} else {
		fmt.Println("YES")
		for i := 0; i < n; i++ {
			fmt.Printf("%f ", x[i])
		}
	}

}