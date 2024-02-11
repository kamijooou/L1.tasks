package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func publisher(cancel <-chan struct{}) <-chan int {
	in := make(chan int)
	go func() {
		defer close(in)
		for i := 0; ; i++ {
			select {
			case <-cancel:
				fmt.Println("I'm done!")
				return
			case in <- i:
			}
		}
	}()
	return in
}

func consumer(out <-chan int, cancel <-chan struct{}) {
	go func() {
		for {
			select {
			case <-cancel:
				fmt.Println("I'm done!")
				return
			case val, ok := <-out:
				if ok {
					fmt.Println(val)
				}
			}
		}
	}()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("add timeout, please")
	}
	n, err := strconv.Atoi(os.Args[1])
	fmt.Println(n)
	if err != nil {
		fmt.Println(err)
		return
	}

	// используем каналы отмены, тк горутины будут подстраиваться под наше время работы
	pubCancel := make(chan struct{})
	conCancel := make(chan struct{})

	out := publisher(pubCancel)
	consumer(out, conCancel)

	<-time.After(time.Duration(n) * time.Second)
	pubCancel <- struct{}{}
	conCancel <- struct{}{}
}
