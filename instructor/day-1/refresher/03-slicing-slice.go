package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []int{1, 2, 3, 4, 5}

	fmt.Println(s, len(s), cap(s))

	p := s[2:4]
	p[1] = 99
	fmt.Println(p, len(p), cap(p))

	s = append(s[:3], s[4:]...)

	// s = append(s, 98, 99, 100)
	fmt.Println(s)

	// deleting an element from slice - performant
	s[2] = s[len(s)-1]
	s = s[:len(s)-1]
	fmt.Println(s)

	// using delete function
	s = slices.Delete(s, 1, 2)
	fmt.Println(s)
}
