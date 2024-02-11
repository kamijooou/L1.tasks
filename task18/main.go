package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	val int64
}

func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&(c.val))
}

func main() {
	counter := &Counter{}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		atomic.AddInt64(&(counter.val), 1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		atomic.AddInt64(&(counter.val), 1) // безопасно прибавляем
	}()

	wg.Wait()
	fmt.Println(counter.Value()) // безопасно читаем
}
