package main

import (
	"fmt"
	"sync"
	"time"
)

func benchmark(name string) func() {
	startTime := time.Now()
	return func() {
		fmt.Printf("%s() worked for %v", name, time.Since(startTime).Nanoseconds())
	}
}

func syncSquare(nums [5]int) {
	defer benchmark("syncSquare")()
	for _, num := range nums {
		fmt.Println(num * num)
	}
}

func square(nums [5]int) {
	defer benchmark("square")()
	wg := sync.WaitGroup{}
	// будет дэдлок, если не выполнятся все wg.Done()
	// или паника, если wg.Done() будет больше необходимого
	wg.Add(5)
	for _, num := range nums {
		/* нужно либо в теле цикла делать копию переменной num,
		либо передавать копию значения в аргумент функции,
		иначе некоторые горутины будут читать текущее значение num,
		которое после завершения цикла будет равно последнему значению (10) */
		go func(num int) {
			defer wg.Done()
			fmt.Println(num * num)
		}(num)
	}
	wg.Done()
	wg.Wait()
}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	square(arr)     // ~5млн наносекунд
	syncSquare(arr) // ~500k наносекунд

	/* это происходит, потому что планировщик тратит много времени на
	на переключение контекста (доп операции), в данном случае выигрыш во времени
	не реализуется, но если бы обе функции обращалась к массиву и изменяли его, то
	тогда асинхронная будет быстрее */
}
