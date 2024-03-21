package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func (emp Employee) Format() string {
	return fmt.Sprintf("Id: %d, Name: %s, Salary: %.2f", emp.Id, emp.Name, emp.Salary)
}

func (emp *Employee) AwardBonus(bonus float32) {
	emp.Salary += bonus
}

func main() {

	// pointers in Go
	// pass by reference and pass by value
	// Pass by value

	// Pointers on Arrays
	nos1 := [5]int{1, 2, 4, 3, 5}
	nos2 := [5]int{1, 2, 4, 3, 5}

	fmt.Println(nos1 == nos2)

	// Pointers on Structs
	e1 := Employee{100, "employee-1", 5000}
	e2 := Employee{100, "employee-1", 5000}
	fmt.Println(e1 == e2)

	emp := Employee{100, "employee-1", 5000}
	fmt.Println("Before increase: ", emp)
	IncreaseSalary(&emp, 10)
	fmt.Println("After increase: ", emp)

	// pointers on primitives
	var no int
	no = 100

	noPtr := &no
	fmt.Println(no, noPtr, *noPtr)

	var somePtr *int // what should be the type here?
	// type safe pointers
	somePtr = &no

	// Dereferencing
	fmt.Println(*somePtr) // dereferencing = getting the value from the memory location (pointer)

	fmt.Println(no == *(&no))

	fmt.Println("Before incrementing: ", no)
	increment(&no)
	fmt.Println("After incrementing: ", no)

	// pointer on structs
	newEmployee := &Employee{300, "emp-3", 100}
	fmt.Println(newEmployee)
	fmt.Printf("%p\n", newEmployee)
	fmt.Println(reflect.TypeOf(newEmployee))
	fmt.Println(reflect.TypeOf(emp))

	fmt.Println("Employee ID: ", (*newEmployee).Id)
	fmt.Println("Employee ID: ", newEmployee.Id)

	// pointers on methods = function with receiver
	fmt.Println(emp.Format())
	fmt.Println("Before awarding bonus: ", emp)
	emp.AwardBonus(2000)
	fmt.Println("After awarding bonus: ", emp)
}

func increment(val *int) {
	*val++
}

func IncreaseSalary(emp *Employee, percentage float32) {
	emp.Salary = ((100 + percentage) / 100) * emp.Salary
}
