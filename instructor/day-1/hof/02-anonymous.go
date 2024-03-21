package main

import "fmt"

func main() {

	// anonymous functions are for immediate consumption
	func() {
		fmt.Println("anonymous fn invoked")
	}()

	// pass values to anonymous functions
	func(x, y int) {
		result := x + y
		fmt.Println("Result: ", result)
	}(100, 200)

	fmt.Println(func(x, y int) int {
		return x + y
	}(100, 200))
}
