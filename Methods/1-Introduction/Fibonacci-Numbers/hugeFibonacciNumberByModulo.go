package main

import "fmt"

// hugeFibonacciNumberByModulo возвращает
// остаток от деления n-го числа Фибоначчи на m.
func hugeFibonacciNumberByModulo(n, m uint) uint {
	// Матрица оператор
	A := [][]uint{
		{1, 1},
		{1, 0},
	}
	// Вспомогательная матрица
	Q := [][]uint{
		{1, 1},
		{1, 0},
	}
	// Вектор столбец результата
	R := []uint{1, 0}

	for n > 0 {
		if n&1 == 1 {
			R1 := (R[0]*A[0][0] + R[1]*A[1][0]) % m
			R2 := (R[0]*A[0][1] + R[1]*A[1][1]) % m
			R[0] = R1
			R[1] = R2
		}
		Q[0][0] = (A[0][0]*A[0][0] + A[0][1]*A[1][0]) % m
		Q[0][1] = (A[0][0]*A[0][1] + A[0][1]*A[1][1]) % m
		Q[1][0] = (A[1][0]*A[0][0] + A[1][1]*A[1][0]) % m
		Q[1][1] = (A[1][0]*A[0][1] + A[1][1]*A[1][1]) % m

		A[0][0] = Q[0][0]
		A[0][1] = Q[0][1]
		A[1][0] = Q[1][0]
		A[1][1] = Q[1][1]

		n >>= 1

	}
	//возвращаем Fn % m
	return R[1]
}


func main() {
	var n, m uint
	fmt.Scan(&n, &m)
	fmt.Println(hugeFibonacciNumberByModulo(n, m))
	// fmt.Println(fibonacciNumberByModuloWithPisanoPeriod(n, m))

}

func fibonacciNumberByModuloWithPisanoPeriod(n, modulo uint) uint {
	pisano := pisanoPeriod(modulo)
	return pisano[n%uint(len(pisano))]
}

func pisanoPeriod(m uint) []uint {
	// Кладем первые два остатка
	pisano := []uint{0, 1}
	var fn_m1, fn uint = 0, 1
	var n uint = 0
	// Цикл до момента, пока мы не встретим 0, 1 идущие подряд,
	// не на первых двух позициях
	for ; n <= 6*m; n++ {
		fn_m1, fn = fn, (fn_m1+fn)%m
		// Заполняем слайс остатками
		pisano = append(pisano, fn)
		if pisano[len(pisano)-2] == 0 && pisano[len(pisano)-1] == 1 {
			// Обрезаем эти символы, зачем нам начало нового цикла)
			pisano = pisano[:len(pisano)-2]
			break
		}
	}

	return pisano
}
