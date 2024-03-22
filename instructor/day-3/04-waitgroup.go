package main

import (
	"fmt"
	"sync"
	"time"
)

// Its a semaphone based counter

func main() {
	// var wg sync.WaitGroup
	wg := &sync.WaitGroup{}
	// increment the counter
	wg.Add(1)
	go f1(wg)

	wg.Add(1)
	go f2(wg)

	// waiting here till the counter becomes zero
	wg.Wait()
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("f1 invoked")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
	// decrement the counter
}

func f2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("f2 invoked")
}
