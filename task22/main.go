package main

import (
	"fmt"
	"math/big" // реализует арифметику произваольной точности (большие числа)
)

func main() {
	a := big.NewInt(1 << 21)
	b := big.NewInt(2 << 21)
	res := &big.Int{}

	res.Mul(a, b)
	fmt.Println(res)

	res.Div(a, b)
	fmt.Println(res)

	res.Add(a, b)
	fmt.Println(res)

	res.Sub(a, b)
	fmt.Println(res)
}
