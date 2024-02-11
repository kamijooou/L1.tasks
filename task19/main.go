package main

import "fmt"

func reverse(str string) string {
	// []byte не подойдет, тк range не будет учитывать префикc символа
	res := make([]rune, len(str))
	// определит, есть ли руны. в отличие от for с условием
	// передает в i индекс первого байта символа
	for i, v := range str {
		res[len(str)-1-i] = v
	}
	return string(res)
}

func main() {
	str := "главрыба"
	fmt.Println(str)
	res := reverse(str)
	fmt.Println(res)
}
