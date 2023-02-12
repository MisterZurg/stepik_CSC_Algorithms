package main

import (
	"fmt"
	"sort"
)

type pointType int

const (
	LEFT pointType= iota - 1
	POINT
	RIGHT
)

type Point struct {
	value int
	pType pointType
	index int
}

//func isSort(a, b Point) bool {
//	if a.value == b.value {
//		if a.pType == LEFT {
//			return true
//		} else if b.pType== LEFT {
//			return false
//		} else if a.pType == POINT {
//			return true
//		} 	else if b.pType== POINT {
//			return false
//		} else if a.pType == RIGHT {
//			return true
//		} else if b.pType == RIGHT {
//			return false
//		}
//	}
//
//	return a.value < b.value
//}

func gavnoSort(points []Point) {
	sort.Slice(points, func(left, right int) bool {
		if points[left].value != points[right].value {
			return points[left].value < points[right].value
		} else { // Определяется типом
			return points[left].pType < points[right].pType
		}
	})
}


func main() {
	points, m := VerySlowParseInput()
	fmt.Println(points, m)
	fmt.Print("Start sol")
	Solution(points, m)
	fmt.Print("End sol")
}

func VerySlowParseInput() ([]Point, int) {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(n, m)
	// Рассматриваем отрезки
	points := make([]Point, n * 2 + m)
	for i := range points {
		var a , b int
		fmt.Scan(&a, &b)

		min := Min(a, b)
		points[i].value = min
		points[i].pType = LEFT
		points[i].index = -1

		max := Max(a, b)
		points[i+1].value = max
		points[i+1].pType = RIGHT
		points[i+1].index = -1
		fmt.Println(min, max)

		i++
	}

	for i := n * 2 ; i < n * 2 + m; i++ {
		var point int
		fmt.Scan(&point)

		points[i].value = point
		points[i].pType = POINT
		points[i].index = i - n * 2 // <= i-я по счёту точка
	}

	return points, m
}

func Solution(points []Point, m int) {
	pointCount := make([]int, m)
	// Сортируем
	gavnoSort(points)
	cnt := 0
	for _, p := range points {
		if p.pType == LEFT {
			cnt++
		} else if p.pType == RIGHT{
			cnt--
		} else {
			pointCount[p.index] = cnt
		}
	}
	// Выводим ответ
	for _, num := range pointCount {
		fmt.Printf("%d ", num)
	}
}

func Max(a, b int) int{
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int{
	if a < b {
		return a
	}
	return b
}

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//type Segment struct {
//	first   int
//	last    int
//}
//
//func main() {
//	// put your code here
//	segments, points := ParseInput()
//	Solution(segments, points)
//	// Quick3(segments, 0, len(Quick3) - 1)
//}
//
//func Solution(segments []Segment, points []int) {
//	// Шаг 1: Сортируем концы отрезков по возрастанию
//	Quick3ForStructure(segments, 0, len(segments) - 1)
//	// Шаг 2: Заведем мапу, которая будет хранить ключ [точку] значение кол-во отрезков
//	memorization := make(map[int]int)
//	// Шаг 3: Пройдемся точками по отрезкам
//	for point := range points{
//		for seg := range segments{
//			// fmt.Println(points[point], segments[seg])
//			// Точка считается принадлежащей отрезку, если она находится внутри него или на границе.
//			if segments[seg].first <= points[point] && points[point] <= segments[seg].last {
//				memorization[points[point]]++
//			}
//		}
//	}
//
//	// Шаг 4: Пройдемся точками для вывода ответа
//	for point := range points{
//		fmt.Printf("%d ", memorization[points[point]])
//	}
//}
//
//func VerySlowParseInput() ([]Segment, []int) {
//	var n, m int
//	fmt.Scan(&n, &m)
//
//	segments := make([]Segment, n)
//	for i := range segments {
//		fmt.Scan(&segments[i].first, &segments[i].last)
//	}
//
//	points := make([]int, m)
//
//	for i := range points {
//		fmt.Scan(&points[i])
//	}
//	// fmt.Println(segments, points)
//	return segments, points
//}
//
//func SlowParseInput() ([]Segment, []int) {
//	scanner := bufio.NewScanner(os.Stdin)
//	scanner.Scan()
//	splited := strings.Split(scanner.Text(), " ")
//
//	var n, m int
//	n, _ = strconv.Atoi(splited[0])
//	m, _ = strconv.Atoi(splited[1])
//
//	segments := make([]Segment, n)
//	for i := range segments {
//		scanner.Scan()
//		seg := strings.Split(scanner.Text(), " ")
//		segments[i].first, _ = strconv.Atoi(seg[0])
//		segments[i].last, _ = strconv.Atoi(seg[1])
//	}
//
//	points := make([]int, m)
//	for i := range points {
//		scanner.Scan()
//		points[i], _ = strconv.Atoi(scanner.Text())
//	}
//
//	return segments, points
//}
//
//func ParseInput() ([]Segment, []int) {
//	// rand.Seed(time.Now().UnixNano())
//	reader := bufio.NewReader(os.Stdin)
//
//	s, _ := reader.ReadString(' ') // nr of lines
//	n, _ := strconv.Atoi(strings.TrimSuffix(s, " "))
//	s, _ = reader.ReadString('\n') // nr of points
//	m, _ := strconv.Atoi(strings.TrimSuffix(s, "\n"))
//
//	segments := make([]Segment, n)
//	for idx := 0; idx < n; idx++ {
//		segmentsString, _ := reader.ReadString('\n')
//		// segmentsStringArr := strings.Split(strings.TrimSuffix(segmentsString, "\n"), " ")
//		segmentsStringArr := strings.Split(segmentsString, " ")
//		segments[idx].first, _ = strconv.Atoi(segmentsStringArr[0])
//		segments[idx].last, _ = strconv.Atoi(strings.TrimSuffix(segmentsStringArr[1], "\n"))
//	}
//
//	ss, _ := reader.ReadString('\n')
//	ss = strings.TrimSuffix(ss, "\n")
//	spoints := strings.Split(ss, " ")
//	// fmt.Println("spoints",spoints)
//	points := make([]int, m)
//	for i := range spoints {
//		points[i], _ = strconv.Atoi(spoints[i])
//	}
//	// fmt.Println(segments, points)
//	return segments, points
//}
//
//func Quick3ForStructure(seg []Segment, left, right int) {
//	if left >= right {
//		return
//	}
//	m1, m2 := Partition3ForStructure(seg, left, right)
//
//	Quick3ForStructure(seg, left, m1 - 1)
//	Quick3ForStructure(seg, m2 + 1, right)
//}
//
//
//func Partition3ForStructure(seg []Segment, left, right int) (int, int){
//	// Опорный элемент
//	pivot := seg[left]
//	m1 := left    // Указатель на часть меньше опорного элемента
//	m2 := right   // Указатель на часть больше опорного элемента
//
//	i := left // Просто итератор, изначально ставим в левое "начальное" положение
//
//	for i <= m2 {
//		if seg[i].last < pivot.last {
//			seg[m1], seg[i] = seg[i], seg[m1]
//			m1++
//			i++
//		} else if seg[i].last > pivot.last {
//			seg[m2], seg[i] = seg[i], seg[m2]
//			m2--
//		} else {
//			i++
//		}
//	}
//	return m1, m2
//}
//
//
//// ImprovingQuickSort returns given sequence sorted in non-decreasing order.
//func ImprovingQuickSort(sequence []int, l, r int) {
//	// Here we have to implement 3-way partition
//	if l >= r {
//		return
//	}
//	m1, m2 := Partition3(sequence, l, r)
//
//	ImprovingQuickSort(sequence, l, m1-1)
//	ImprovingQuickSort(sequence, m2+1, r)
//}
//
//// Partition3 is a helper function for ImprovingQuickSort function
//// realizes 3-way partition to handle few equal elements in slice
//func Partition3(sequence []int, l, r int) (int, int) {
//	pivot := sequence[l]
//	m1 := l // We initiate m1 to be the part that is less than the pivot
//	m2 := r // The part that is greater than the pivot
//	i := l
//	for i <= m2 {
//		if sequence[i] < pivot {
//			sequence[m1], sequence[i] = sequence[i], sequence[m1]
//			m1++
//			i++
//		} else if sequence[i] > pivot {
//			sequence[m2], sequence[i] = sequence[i], sequence[m2]
//			m2--
//		} else {
//			i++
//		}
//	}
//
//	return m1, m2
//}
//
/*
3 2
0 5
-3 2
7 10
1 6
//
//2 0
// */