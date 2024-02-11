package main

import "fmt"

func main() {
	a, b := 1, 2
	a, b = b, a // работает благодаря многозначным присваиваниям
	fmt.Println(a, b)
}
