package main

import (
	"fmt"
	"time"
)

func main() {
	go f1()
	f2()

	// time.Sleep(1 * time.Second)
	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 invoked")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
