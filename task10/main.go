package main

import (
	"fmt"
)

// О(n) * O(1) = O(n)
func grouping(tempSeq []float64) map[int64][]float64 {
	// интуитивно из условия хочется выбрать такую мапу, чтобы комфортнее складывать значения
	groups := map[int64][]float64{}

	// O(n)
	for _, v := range tempSeq {
		// Название группы - это число с нулевым первым разрядом,
		// тк в условии указан шаг 10 с примером.
		// Значит, чтобы его получить, нужно поделить на 10,
		// убрать вещественную часть (там как раз и первый разряд),
		// а потом умножить на 10. Вот мы и получили ключ для
		// подмножества!
		groupName := int64(v/10) * 10
		// проверяем, нет ли такой группы уже. О(1)
		if _, ok := groups[groupName]; ok {
			slc := groups[groupName]
			groups[groupName] = append(slc, v) // обновляем слайс, не опасаясь аллокации
			continue
		}
		groups[groupName] = []float64{v} // если группы не было, создаем ее
	}

	return groups
}

func main() {
	tempSeq := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	res := grouping(tempSeq)
	fmt.Println(res)
}
