package main

import (
	"context"
	"fmt"
	"time"
)

func producer(ctx context.Context, ch chan<- string) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Print("producer [cancel]\n")
				return

			default:
				select {
				case ch <- "countdown":
					time.Sleep(100 * time.Millisecond)
				case <-ctx.Done():
					return
				}
			}
		}
	}()
}

func receiver(ctx context.Context, ch <-chan string) {
	go func() {
		for {
			select {
			case msg := <-ch:
				fmt.Printf("msg: %v\n", msg)

			case <-ctx.Done():
				fmt.Print("receiver [cancel]\n")
				return
			}
		}
	}()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan string)

	producer(ctx, ch)
	producer(ctx, ch)

	receiver(ctx, ch)
	receiver(ctx, ch)

	time.Sleep(1000 * time.Millisecond)

	cancel()

	time.Sleep(4 * time.Millisecond)
}
