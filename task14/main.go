package main

import "fmt"

func typeOf(T any) { // алиас пустого интерфейса
	switch T.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan any:
		fmt.Println("chan")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	meaning := 42
	bullyMcguire := true
	str := "кто прочитал, тот засчитал мне L1"
	chiDa := make(chan any)
	chiNet := make(chan chan chan chan int)
	typeOf(meaning)
	typeOf(bullyMcguire)
	typeOf(str)
	typeOf(chiDa)
	typeOf(chiNet)
}
