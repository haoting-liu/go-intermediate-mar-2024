package main

import "fmt"

// Higher order functions: functions as values
// functions as first citizens
func main() {

	// functions are also type safe in Go
	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	var addProcessor func(int, int)
	addProcessor = func(x, y int) {
		result := x + y
		fmt.Println("Result: ", result)
	}
	addProcessor(100, 200)

	mathOperation := func(x, y int) int {
		return x + y
	}
	fmt.Println(mathOperation(100, 200))

	mathOperation = func(x, y int) int {
		return x - y
	}
	fmt.Println(mathOperation(100, 200))

	// passing functions as arguments
	// returning functions as return values
}

// func fn() {
// 	fmt.Println("fn invoked")
// }
