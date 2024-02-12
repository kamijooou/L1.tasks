package main

import (
	"fmt"
	"time"
)

func producer(stop <-chan struct{}, ch chan<- string) {
	go func() {
		for {
			select {
			case <-stop:
				fmt.Print("producer [stop]\n")
				return

			default:
				select {
				case ch <- "countdown":
					time.Sleep(100 * time.Millisecond)
				case <-stop:
					return
				}
			}
		}
	}()
}

func receiver(stop <-chan struct{}, ch <-chan string) {
	go func() {
		for {
			select {
			case msg := <-ch:
				fmt.Printf("msg: %v\n", msg)

			case <-stop:
				fmt.Print("receiver [stop]\n")
				return
			}
		}
	}()
}

func main() {
	stop := make(chan struct{})
	ch := make(chan string)

	producer(stop, ch)
	producer(stop, ch)

	receiver(stop, ch)
	receiver(stop, ch)

	time.Sleep(1000 * time.Millisecond)

	close(stop)

	time.Sleep(4 * time.Millisecond)
}
