package main

import (
	"fmt"
	"time"
)

// consumer
func main() {
	ch := getNos()

	// go func() {
	// 	ch <- 1000
	// }()

	for data := range ch {
		fmt.Println(data)
	}
}

// producer
func getNos() <-chan int {
	ch := make(chan int)

	go func() {
		count := 5
		for i := 1; i <= count; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i * 10
		}
		close(ch) // I don't have any more data
	}()

	return ch
}
