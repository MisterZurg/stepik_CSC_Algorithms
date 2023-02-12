package main

type vector struct {
	begin int
	end int
}

func (v vector)normalize() {

}

func (v vector)add() {

}

func (v vector)show(){

}

type equationSystem struct {

}



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
//	matrix, b := makeAValueMatrices(n, m)
//
//
//	Solution(n, m, matrix, b)
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
//	for p := 0; p < N; p++ {
//		max := p
//
//		for i := p + 1; i < N; i++ {
//			if math.Abs(matrix[i][p]) > math.Abs(matrix[max][p]) {
//				max = i
//			}
//		}
//
//
//		matrix[p], matrix[max] = matrix[max], matrix[p]
//		b[p], b[max] = b[max], b[p]
//
//		if math.Abs(matrix[p][p]) <= 1e-14 {
//			return NoSolutions, nil
//		}
//
//		// Pivot within matrix and b
//		for i := p + 1; i < N; i++ {
//			alpha := matrix[i][p] / matrix[p][p]
//			b[i] -= alpha * b[p]
//			for j := p; j < N; j++ {
//				matrix[i][j] -= alpha * matrix[p][j] // сделать элементы ниже поворотных элементов равными нулю или устранить переменные
//			}
//		}
//	}
//	return reverseGaussMethod(N, M, matrix, b)
//}
//
//func reverseGaussMethod(n, m int, matrix [][]float64, b []float64) (Solutions, []float64) {
//	if n < m {
//		return InfinitySolutions, nil
//	}
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
//	return OneSolution, x
//}
//
//func makeAValueMatrices(n, m int) ([][]float64, []float64) {
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
//		for j := 0; j < m; j++ {
//			fmt.Scan(&matrix[i][j])
//		}
//		fmt.Scan(&b[i])
//	}
//	return matrix, b
//}
//
