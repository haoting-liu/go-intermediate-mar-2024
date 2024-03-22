package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
HTOP: runs as htop

- Press F2 for SETUP
- In the screen -> Main -> Active Columns -> Available Columns
- Choose NLWP (Press ENTER)
- Press q to comeout of the setup mode
- F4 for filter and add the binary name

*/

// Its a semaphone based counter
func main() {
	count := flag.Int("count", 0, "Number of go routines to spin up")
	flag.Parse()
	fmt.Printf("%d number of goroutings staring.. Press ENTER\n", *count)
	fmt.Scanln()

	wg := &sync.WaitGroup{}
	for i := 0; i < *count; i++ {
		wg.Add(1)
		go f1(i, wg)
	}

	// waiting here till the counter becomes zero
	wg.Wait()
	fmt.Println("All goroutines finished. Press ENTER to complete")
	fmt.Scanln()
}

func f1(id int, wg *sync.WaitGroup) {
	// decrement the counter
	defer wg.Done()
	fmt.Printf("fn[%d] invoked\n", id)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}
