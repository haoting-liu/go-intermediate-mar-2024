package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var cnt int64

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("Result: ", cnt)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&cnt, 1)
}
