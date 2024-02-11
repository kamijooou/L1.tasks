package main

import (
	"fmt"
	"sort"
	"time"
)

func benchmark(name string) func() {
	startTime := time.Now()
	return func() {
		fmt.Printf("%s() worked for %v \n", name, time.Since(startTime).Nanoseconds())
	}
}

type QSlc []int

func (slc QSlc) Less(i, j int) bool {
	return slc[i] < slc[j]
}

func (slc QSlc) Swap(i, j int) {
	slc[i], slc[j] = slc[j], slc[i]
}

func (slc QSlc) Len() int {
	return len(slc)
}

func quickSort(slc []int, pivot int) []int {
	switch len(slc) {
	case 0:
		return []int{}
	case 1:
		return slc
	case 2:
		if slc[0] > slc[1] {
			return []int{slc[1], slc[0]}
		}
		return []int{slc[0], slc[1]}
	}

	left := make([]int, 0, len(slc)/2+1)
	right := make([]int, 0, len(slc)/2+1)
	for i := 1; i < len(slc); i++ {
		if slc[i] >= slc[pivot] && i != pivot {
			right = append(right, slc[i])
		} else if slc[i] < slc[pivot] && i != pivot {
			left = append(left, slc[i])
		}
	}

	left = quickSort(left, 0)
	right = quickSort(right, 0)

	result := append(left, slc[pivot])
	return append(result, right...)
}

func run1() {
	slc := []int{2, 6, 4, 8, 3, 9, 1, 0, 1, 6, 7}
	defer benchmark("quickSort")() // ~49200ns :((
	res := quickSort(slc, 0)
	fmt.Println(res)
}

func run2() {
	qslc := QSlc([]int{2, 6, 4, 8, 3, 9, 1, 0, 1, 6, 7})
	defer benchmark("sort.Sort")() // ~3200ns
	sort.Sort(qslc)
	fmt.Println(qslc)
}

/*
ВЫВОД -> реализуем интерфейс и используем встроенный метод :)
Так происходит потому, что разрабы выбрали гибридную версию сортировки
под названием pattern-defeating quick sort.
Она представляет из себя быструю сортровку, сортировку вставками и резервную
(heapsort в гошке, обычно приходится использовать редко). Маленькие данные
сортируются вставками, остальные разделяются и рекурсивно сортируются быстрой.
Если глубина рекурсии становится слишком большой, начинают использовать резервную.
*/
func main() {
	run1()
	run2()
}
