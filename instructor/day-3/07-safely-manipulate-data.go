package main

import (
	"fmt"
	"sync"
)

// with race detector
// go run --race <filename.go>
// go build --race <filename.go>

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

var counter Counter

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("Result: ", counter.count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Increment()
}
