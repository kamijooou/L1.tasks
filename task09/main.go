package main

import (
	"fmt"
	"sync"
)

func publish(cancel <-chan struct{}) <-chan int {
	out := make(chan int)
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	go func() {
		defer close(out)
		for _, v := range arr {
			select {
			case out <- v:
			case <-cancel:
				return
			}
		}
	}()
	return out
}

// обработчик вычисляет квадрат
func handle(cancel <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			select {
			case out <- num * num:
			case <-cancel:
				return
			}
		}
	}()
	return out
}

// объединяет c1 и c2 в общий канал
func merge(cancel <-chan struct{}, c1, c2 <-chan int) <-chan int {
	var wg sync.WaitGroup
	wg.Add(2)

	out := make(chan int)
	merge := func(c <-chan int) {
		defer wg.Done()
		for num := range c {
			select {
			case out <- num:
			case <-cancel:
				return
			}
		}
	}
	go merge(c1)
	go merge(c2)

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func print(cancel <-chan struct{}, in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}

// канал отмены передается всем для улучшения контроля над горутинами
func main() {
	cancel := make(chan struct{})
	defer close(cancel)

	c1 := publish(cancel)
	c2_1 := handle(cancel, c1)
	c2_2 := handle(cancel, c1)
	c3 := merge(cancel, c2_1, c2_2)
	print(cancel, c3)
}
