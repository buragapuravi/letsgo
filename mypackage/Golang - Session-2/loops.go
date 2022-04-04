package main

import (
	"fmt"
)

func main() {
	// Regular For loop
	count := 0
	for i := 1; i < 5; i++ {
		count += i
	}
	fmt.Println("For Loop : ", count)

	//While loop
	j := 1
	for j < 5 {
		j *= 2
	}
	fmt.Println("While Loop: ", j)

	// For Each loop
	people := []string{"person1", "person2", "person3"}
	fmt.Println("For Each Loop ")
	for i, s := range people {
		fmt.Println(i, s)
	}

}
