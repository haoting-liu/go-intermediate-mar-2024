package main

import (
	"fmt"
	"log"
)

func main() {
	// add(100, 200)
	// sub(100, 200)

	// log.Println("operation started")
	// add(100, 200)
	// log.Println("operation completed")

	// log.Println("operation started")
	// sub(100, 200)
	// log.Println("operation completed")

	// logAdd(100, 200)
	// logSub(100, 200)
	// DRY: Do not repeat yourself. remove duplicacy

	// logOperation(add, 100, 200)
	// logOperation(sub, 100, 200)

	logAdd := getLogOperation(add)
	logAdd(100, 200)

	logSub := getLogOperation(sub)
	logSub(100, 200)

}

func getLogOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("operation started")
		operation(x, y)
		log.Println("operation completed")
	}
}

func logOperation(operation func(int, int), x, y int) {
	log.Println("operation started")
	operation(x, y)
	log.Println("operation completed")
}

// func logAdd(x, y int) {
// 	log.Println("operation started")
// 	add(x, y)
// 	log.Println("operation completed")
// }

// func logSub(x, y int) {
// 	log.Println("operation started")
// 	sub(x, y)
// 	log.Println("operation completed")
// }

func add(x, y int) {
	fmt.Println("Add result:", x+y)
}

func sub(x, y int) {
	fmt.Println("Sub result:", x-y)
}
