// Встроенного метода нет, пишем сами.
// Первый способ стандартный, второй быстрее и дешевле по памяти,
// но лишь заменяет значение копией другого элемента
package main

import "fmt"

func remove1(slc []int, i int) []int {
	res := make([]int, 0)
	switch true {
	case len(slc) == 0:
	case len(slc) == 1:
		if i != 0 {
			fmt.Println("index out of range")
			res = slc
		}
	case i > len(slc) || i < 0:
		fmt.Println("index out of range")
		res = slc
	case i == len(slc)-1:
		res = slc[:i]
	default:
		res = append(slc[:i], slc[i+1:]...)
	}
	return res
}

func remove2(slc []int, i int) {
	switch true {
	case len(slc) == 0:
	case len(slc) == 1:
		if i != 0 {
			fmt.Println("index out of range")
		}
	case i > len(slc) || i < 0:
		fmt.Println("index out of range")
	case i == len(slc)-1:
		slc[i] = slc[i-1]
	default:
		slc[i] = slc[len(slc)-1]
	}
}

func main() {
	// помним, что передается копия ссылки на базовый массив,
	// когда будем передавать слайсы в функции.
	slc1 := []int{1, 2, 3, 4}
	slc2 := []int{1, 2, 3, 4}

	res := remove1(slc1, 2)
	fmt.Println(res)

	remove2(slc2, 1)
	fmt.Println(slc2)
}
