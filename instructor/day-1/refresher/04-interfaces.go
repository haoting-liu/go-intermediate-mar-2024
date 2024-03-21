package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Name   string
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("%s with radius: %.2f", c.Name, c.Radius)
}

type Rectangle struct /* implements AreaFinder */ {
	Name   string
	Length float32
	Bredth float32
}

func (r Rectangle) Area() float32 {
	return (r.Bredth * r.Length)
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Bredth + r.Length)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("%s with length: %.2f, bredth: %.2f", r.Name, r.Length, r.Bredth)
}

type AreaFinder interface {
	Area() float32
}

type PerimeterFinder interface {
	Perimeter() float32
}

func main() {
	c := Circle{"small cicle", 12}
	// fmt.Println("Circle Area: ", c.Area())

	r := Rectangle{"large rec", 10, 12}
	// fmt.Println("Rectangle Area: ", r.Area())

	// PrintArea(c)
	// PrintArea(r)

	// fmt.Println(c)
	// fmt.Println(r)

	// PrintPerimeter(c)
	// PrintPerimeter(r)

	PrintShape(c)
	PrintShape(r)
}

// interface composition
type ShapeFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintShape(x ShapeFinder) {
	// PrintArea(x) -> should implement AreaFinder
	// PrintPerimeter(x) -> should implement PerimeterFinder
	PrintArea(x)
	PrintPerimeter(x)
}

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Permeter: ", x.Perimeter())
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area: ", x.Area())
}
