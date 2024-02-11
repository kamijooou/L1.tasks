package main

import "fmt"

/*Используем значения пустых структур, тк они не занимают память. Сложность алгоритма О(n)*/
func set(seq []string) []string {
	seqTab := make(map[string]struct{})
	for _, word := range seq {
		if _, ok := seqTab[word]; ok {
			continue
		}
		seqTab[word] = struct{}{}
	}

	// выделяем сразу нужное количество слотов, чтобы не было аллокаций
	res := make([]string, len(seqTab))
	idx := 0
	for word := range seqTab {
		res[idx] = word
		idx++
	}
	fmt.Println(len(res))
	return res
}

func main() {
	seq := []string{"cat", "cat", "dog", "cat", "tree"}
	res := set(seq)
	fmt.Println(res)
}
