package main

import (
	"fmt"
)

func main() {

	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(b)

	//logically creates an array and returns a slice out of it
	c := []int{6, 7, 8} //creates and array and returns a slice reference
	fmt.Println(c)
	// how system creates a slice
	i := make([]int, 5, 5)
	fmt.Println(i)
	// Scalable buffered  array slice
	slicevar := make([]int, 1, 4)
	fmt.Println(slicevar)
}
