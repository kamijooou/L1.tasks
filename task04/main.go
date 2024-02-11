package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

// в отдельной горутине ждем, пока все другие завершатся,
// после этого выходим
func stop(wg *sync.WaitGroup) {
	wg.Wait()
	fmt.Println("\nGOOD BYE")
	os.Exit(0)
}

// каждая горутина пытается читать с канала, а также ждет сигнала о выходе,
// если нигде нельзя читать, то блокируется и ждет
// после сигнала о выходе прожимает wg.Done (для завершения всего процесса)
func startWorker(wg *sync.WaitGroup, done <-chan os.Signal, out <-chan int64) {
	go func() {
		defer wg.Done()
		for {
			select {
			case val := <-out:
				fmt.Println(val)
			case <-done:
				fmt.Println("worker stopped")
				return
			}
		}
	}()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("add number of workers, please")
		return
	}

	workers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	// сразу обозначим, сколько воркеров нам дожидаться,
	// тк количество известно
	wg.Add(workers)
	// сделаем такой буффер, чтобы каждый воркер был занят делом
	in := make(chan int64, workers)

	// создаем воркеры, для каждого свой канал завершения, подписанный
	// на нужный сигнал
	for i := 0; i < workers; i++ {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt)
		startWorker(&wg, done, in)
	}

	go stop(&wg) // пускаем ждуна

	rand.New(rand.NewSource(time.Now().Unix()))
	for {
		in <- rand.Int63() // пишем
	}
}

/*
Вообще для завершения горутин более читабельным вариантом является завершение через контекст,
но у он под капотом также возвращает канал. Просто захотелось попробовать.

А вот для ждуна можно было помимо ВГ использовать пул каналов, с которых пытался бы читать ждун,
но это менее красиво, как мне кажется. Также еще можно завершить через буферизированный канал, откуда бы все
воркеры понабрали бы токенов, а после работы бы их вернули на место, а ждун пытался бы считать N раз токены из буффера.
Все варианты имеют место быть, но мне очень понравилось пользоваться ВГ.
*/
