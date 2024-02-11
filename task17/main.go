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

// изначально думал рекурсивно вызывать функцию, передавая в нее
// слайс от слайса, но исходники мне понравились больше
func search(slc []int, num int) int {
	i, j := 0, len(slc)
	for i < j {
		h := int(uint(i+j) >> 1)
		if !(slc[h] >= num) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

// Загадка джокера. Может быть, встроенные пакеты по-особенному компилируются?...
func run1() {
	slc := []int{1, 2, 3, 3, 4, 5, 6, 7, 7, 8}
	defer benchmark("search")() // ~10400ns :((
	idx := search(slc, 2)
	fmt.Println(idx)
}

func run2() {
	slc := []int{1, 2, 3, 3, 4, 5, 6, 7, 7, 8, 9}
	defer benchmark("sort.SearchInts")() // ~800ns
	idx := sort.SearchInts(slc, 2)
	fmt.Println(idx)
}

func main() {
	run1()
	run2()
}
