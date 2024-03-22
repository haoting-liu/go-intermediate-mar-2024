package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go getNos(ch)     //1. scheduled
	fmt.Println(<-ch) // 3. blocking
	fmt.Println(<-ch) // 5. blocking
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func getNos(ch chan int) {
	time.Sleep(500 * time.Millisecond) // 2. blocking
	ch <- 10                           //4. blocking
	time.Sleep(500 * time.Millisecond)
	ch <- 20
	time.Sleep(500 * time.Millisecond)
	ch <- 30
	time.Sleep(500 * time.Millisecond)
	ch <- 40
	time.Sleep(500 * time.Millisecond)
	ch <- 50
}
