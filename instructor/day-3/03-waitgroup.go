package main

import (
	"fmt"
	"sync"
	"time"
)

// Its a semaphone based counter
var wg sync.WaitGroup

func main() {
	// increment the counter
	wg.Add(1)
	go f1()

	wg.Add(1)
	go f2()

	// waiting here till the counter becomes zero
	wg.Wait()
}

func f1() {
	defer wg.Done()
	fmt.Println("f1 invoked")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
	// decrement the counter
}

func f2() {
	defer wg.Done()
	fmt.Println("f2 invoked")
}
