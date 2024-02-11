package main

import (
	"fmt"
	"strings"
)

func splitReverse(str string) string {
	slc := strings.Split(str, " ")
	for i, v := range slc {
		j := len(slc) - 1 - i
		if i >= j {
			break
		}
		slc[j], slc[i] = v, slc[j]
	}
	return strings.Join(slc, " ")
}

func main() {
	str := "snow dog sun fuck"
	res := splitReverse(str)
	fmt.Println(res)
}
