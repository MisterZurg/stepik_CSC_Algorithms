//package main
//
//// import numpy as np
//import (
//	"fmt"
//	"math"
//)
//
//type Solutions int
//
//const (
//	NoSolutions Solutions = iota -1
//	InfinitySolutions
//	OneSolution
//)
//
//func main() {
//	// put your code here
//	var n, m int // количество уравнений, количество неизвестных
//	fmt.Scan(&n, &m)
//
//	// matrix := makeAugmentedMatrix(n, m)
//	matrix, b := makeAValueMatrices(n, m)
//
//
//	Solution(n, m, matrix, b)
//}
//
//func PrintMatrix(matrix [][]float64)  {
//	fmt.Println()
//	for i := range matrix {
//		for j:= range (matrix)[i]{
//			fmt.Printf("%f ", (matrix)[i][j])
//		}
//		fmt.Println()
//	}
//}
//
//func Solution(n, m int, matrix [][]float64, b []float64) {
//	// YES, если решение существует и единственно,
//	// слово NO в случае, если решение не существует,
//	// и слово INF в случае, когда решений существует бесконечно много
//	var result Solutions
//	var params []float64
//
//
//	result, params = forwardGaussMethod(n, m, matrix, b)
//
//	switch result {
//	case NoSolutions:
//		fmt.Println("NO")
//	case InfinitySolutions:
//		fmt.Println("INF")
//	case OneSolution:
//		fmt.Println("YES")
//		for i := range params{
//			fmt.Printf("%f ", params[i])
//		}
//	}
//}
//
//// forwardGaussMethod returns trapezoidal view or triangular view
//func forwardGaussMethod(N, M int, matrix [][]float64, b []float64) (Solutions, []float64){
//
//	for i := 0; i < N; i++ {
//		max := i
//
//		for j := i + 1; j < N; j++ {
//			if math.Abs(matrix[j][i]) > math.Abs(matrix[max][i]) {
//				max = j
//			}
//		}
//
//		matrix[i], matrix[max] = matrix[max], matrix[i]
//		b[i], b[max] = b[max], b[i]
//
//		if math.Abs(matrix[i][i]) <= 1e-14 {
//			return NoSolutions, nil
//		}
//
//		for k := i + 1; k < N; k++ {
//			alpha := matrix[k][i] / matrix[i][i]
//			b[k] -= alpha * b[i]
//			for m := i; m < N; m++ {
//				matrix[k][m] -= alpha * matrix[i][m]
//			}
//		}
//	}
//	return reverseGaussMethod(N, M, matrix, b)
//}
//
//func reverseGaussMethod(n, m int, matrix [][]float64, b []float64) (Solutions, []float64) {
//	// end of WP algorithm.
//	// now back substitute to get result.
//	N := n
//	x := make([]float64, N)
//	for i := N - 1; i >= 0; i-- {
//		var sum float64
//		for j := i + 1; j < N; j++ {
//			sum += matrix[i][j] * x[j]
//		}
//		x[i] = (b[i] - sum) / matrix[i][i]
//	}
//
//	if n < m {
//		return InfinitySolutions, nil
//	}
//	return OneSolution, x
//}
//
//// makeAndValueMatrices
//func makeAValueMatrices(n, m int) ([][]float64, []float64)  {
//	// Матрица из коэффициентов при неизвестных
//	matrix := make([][]float64, n)
//	// Cтолбец свободных членов
//	b := make([]float64, n)
//
//	// Инициализация nested-слайса
//	for i := range matrix {
//		matrix[i] = make([]float64, m)
//	}
//
//	// Счёт пользовательского инпута
//	for i := 0; i < n; i++ {
//		for j:=0; j < m; j++ {
//			fmt.Scan(&matrix[i][j])
//		}
//		fmt.Scan(&b[i])
//	}
//	return matrix, b
//}
//// makeAugmentedMatrix возвращает расширенную матрицу полученную с пользовательского ввода
//func makeAugmentedMatrix(n, m int) [][]float64 {
//	matrix := make([][]float64, n)
//	for i := range matrix {
//		matrix[i] = make([]float64, m + 1)
//	}
//	// var param float64
//	for i := 0; i < n; i++ {
//		for j:=0; j < m + 1; j++ {
//			fmt.Scan(&matrix[i][j])
//		}
//	}
//	return matrix
//}