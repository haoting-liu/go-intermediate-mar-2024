package main

import (
	"fmt"
	"sync"
)

/**
Channel behaviour:
	- A RECEIVE operation is ALWAYS a blocking operation
	- A SEND operation is blocked until a RECEIVE operation is finished

Whenever you have a blocking operation the Go scheduler looks for
other scheduled goroutines and executes them
*/

func main() {
	// v1()
	v2()
}

func v2() { // receiving data in a goroutine
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	// Go scheduler scheduled anonymous function, but does not executes it
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		data := <-ch //blocking, here Go scheduler looks if it has other goroutines scheduled
		fmt.Println("Data: ", data)
	}(wg)

	ch <- 100 //blocking, here Go scheduler runs the anonymous goroutine
	wg.Wait()
}

func v1() { // sending data in a goroutine
	ch := make(chan int)

	// Go scheduler scheduled anonymous function, but does not executes it
	go func() {
		ch <- 100 //blocking, here Go scheduler runs the main routine
	}()
	data := <-ch //blocking, here Go scheduler looks if it has other goroutines scheduled

	fmt.Println("Data: ", data)
}
