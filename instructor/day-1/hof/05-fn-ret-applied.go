package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// logAdd := getLogOperation(add)
	// logAdd(100, 200)

	// logSub := getLogOperation(sub)
	// logSub(100, 200)

	// profileAdd := getProfileOperation(add)
	// profileAdd(100, 200)
	// profileSub := getProfileOperation(sub)
	// profileSub(100, 200)

	// -profile
	// 	- logging
	// 		- add

	logAdd := getLogOperation(add)
	profileLogAdd := getProfileOperation(logAdd)
	profileLogAdd(100, 200)

	logSub := getLogOperation(sub)
	profileLogSub := getProfileOperation(logSub)
	profileLogSub(100, 200)
}

func getProfileOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		operation(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elasped: ", elapsed)
	}
}

func getLogOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("operation started")
		operation(x, y)
		log.Println("operation completed")
	}
}

func add(x, y int) {
	fmt.Println("Add result:", x+y)
}

func sub(x, y int) {
	fmt.Println("Sub result:", x-y)
}
