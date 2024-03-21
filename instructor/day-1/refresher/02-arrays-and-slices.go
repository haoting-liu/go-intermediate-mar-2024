package main

import "fmt"

func main() {
	nosArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println("before array: ", nosArray)
	multiplyBy2(nosArray)
	fmt.Println("after array: ", nosArray)

	nosSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("before slice: ", nosSlice)
	multiplyBy2Slice(nosSlice)
	fmt.Println("after slice: ", nosSlice)

	fmt.Println("before slice for-range: ", nosSlice)
	multiplyBy2SliceForEach(nosSlice)
	fmt.Println("after slice for-range: ", nosSlice)

}

func multiplyBy2(s [5]int) {
	for i := 0; i < len(s); i++ {
		s[i] = s[i] * 2
	}
}

func multiplyBy2Slice(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] = s[i] * 2
	}
}

func multiplyBy2SliceForEach(s []int) {
	for i, _ := range s {
		s[i] = s[i] * 2
	}
}
