/*
 */
package main

import "fmt"

type Human1 struct {
	a int
	b int
}

func (h Human1) speak() {
	fmt.Println("Omae wa mou shindeiru")
}

type Human2 struct {
	a int
	b float64
}

// Diamond Problem в го не решена, при попытке неявно получить доступ к одноименным полям
// у двух вложенных структур будет ошибка (даже если разные типы), в таком случае нужно
// обращаться к полям явно через вложенные структуры
type Action struct {
	b float64
	// если дать полю имя, то вызывать метод придется явно через это поле, наследования не будет
	Human1
	Human2
}

func main() {
	act := Action{5.0, Human1{1, 2}, Human2{3, 4.0}}
	act.speak() // НАСЛЕДОВАНИЕ
	// будет искать у себя, потом у вложенных структур, если не найдет
	fmt.Println(act.b)        // 5
	fmt.Println(act.Human1.b) // 2
	fmt.Println(act.Human2.b) // 4
	// можно вызвать явно, проблем не будет с этим
	act.Human1.speak()
}
