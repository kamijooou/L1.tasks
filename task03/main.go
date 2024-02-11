package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func atomicSum(nums []int64) int64 {
	var res int64

	wg := sync.WaitGroup{}
	wg.Add(5)

	for i := len(nums) - 1; i >= 0; i-- {
		go func(i int) {
			defer wg.Done()
			// атомарно прибавляем
			atomic.AddInt64(&res, nums[i]*nums[i])
		}(i)
	}

	wg.Wait()
	return res
}

func bufferSum(nums []int64) int64 {
	var res int64
	buffer := make(chan int64, len(nums))

	for _, num := range nums {
		go func(num int64) {
			buffer <- num * num
		}(num)
	}

	for i := 0; i < len(nums); i++ {
		res += <-buffer
	}
	close(buffer)
	return res
}

func muxSum(nums []int64) int64 {
	var res int64
	mux := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(nums))

	for _, num := range nums {
		go func(num int64) {
			defer wg.Done()
			mux.Lock()
			res += num * num
			mux.Unlock()
		}(num)
	}

	wg.Wait()
	return res
}

func main() {
	nums := []int64{2, 4, 6, 8, 10}
	fmt.Println(muxSum(nums))
	fmt.Println(bufferSum(nums))
	fmt.Println(atomicSum(nums))
}
