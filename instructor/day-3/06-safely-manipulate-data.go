package main

import (
	"fmt"
	"sync"
)

// with race detector
// go run --race <filename.go>
// go build --race <filename.go>

var count int

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("Result: ", count)
}

// mutually exclusive lock
var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	{
		count++
	}
	mutex.Unlock()
}
