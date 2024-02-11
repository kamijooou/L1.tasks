/* структура синк мапы
type Map struct {
	mu Mutex
	read atomic.Pointer[readOnly]
	dirty map[any]*entry
	misses int
}
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	// под капотом для создания значений испольует мьютекс
	// значения хранятся в виде атомарного unsafe.Pointer'a Entry,
	// значения которого можно считать/изменить атомарно с использованием
	// атомарных инструкций процессора
	var sMap sync.Map
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		sMap.Store("first", 1)
	}()

	go func() {
		defer wg.Done()
		sMap.Store("second", 2)
		val, ok := sMap.Load("first")
		if ok {
			fmt.Println(val)
		} else {
			// вот это выведет, тк мьютекс, хоть и основан на атомике,
			// работает медленнее, чем Entry
			fmt.Println("'first' not found")
		}

	}()

	wg.Wait() // дождались, теперь можно нормально читать
	val, ok := sMap.Load("second")
	if ok {
		fmt.Println(val)
	} else {
		fmt.Printf("'second' not found")
	}
}
