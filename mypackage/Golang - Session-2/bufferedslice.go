package main

import (
	"fmt"
)

func main() {
	slicevar := make([]int, 1, 4)
	fmt.Println("Length of the slice: ", len(slicevar), " Capacity of slice : ", cap(slicevar))

	for i := 1; i < 20; i++ {
		slicevar = append(slicevar, i)
		fmt.Println("Length of the slice: ", len(slicevar), " Capacity of slice : ", cap(slicevar))
	}

}
