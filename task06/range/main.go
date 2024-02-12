package main

import (
	"fmt"
	"time"
)

// завершает если канал закрыт
func receiver(ch <-chan string) {
	go func() {
		for msg := range ch {
			fmt.Printf("msg: %v\n", msg)
			return
		}
	}()
}

func main() {
	ch := make(chan string)

	receiver(ch)

	ch <- "3 countdown"
	ch <- "2 countdown"
	ch <- "1 countdown"

	close(ch)

	time.Sleep(time.Millisecond)
}
