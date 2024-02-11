package main

import "fmt"

// 3*O(n) = O(n)
func intersec(seq1, seq2 []int) []int {
	seqTab := make(map[int]int)
	res := make([]int, 0)
	if len(seq1) > len(seq2) {
		seq1, seq2 = seq2, seq1 // для пересечения будем заполнять мапу меньшим множеством
	}

	for _, v := range seq1 { // O(n)
		seqTab[v] = 1
	}

	for _, v := range seq2 { // O(n) * O(1) = O(n)
		if _, ok := seqTab[v]; ok {
			seqTab[v]++
			continue
		}
	}

	for k, v := range seqTab { // O(n)
		if v > 1 {
			res = append(res, k)
		}
	}

	return res
}

func main() {
	seq1 := []int{1, 2, 3, 5}
	seq2 := []int{5, 2, 7, 8, 9, 10}
	res := intersec(seq1, seq2)
	fmt.Println(res)
}
