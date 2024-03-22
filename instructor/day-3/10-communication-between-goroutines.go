package main

import (
	"fmt"
)

/**
Channels: data type just like int, float, string, etc.
	- Its not an API

	Channel: data type to facilitate communication between goroutines

	Declare:
		var <var_name> chan <data_type>
		e.g. var ch chan int

	Initialize:
		ch = make(chan int)

	Operations:
		using <- (channel operator)

		SEND operation
			ch <- 100

		RECEIVE operation
			result := <-ch
*/

func main() {
	// wg := &sync.WaitGroup{}
	// wg.Add(1)

	ch := make(chan int)

	// go add(100, 200, wg, ch)
	go add(100, 200, ch)

	result := <-ch
	// wg.Wait()
	fmt.Println("Result: ", result)
}

// func add(x, y int, wg *sync.WaitGroup, ch chan int) {
func add(x, y int, ch chan int) {
	// defer wg.Done()
	result := x + y
	ch <- result
}
