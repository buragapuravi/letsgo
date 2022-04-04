package main

import "fmt"

func main() {
	cpu := 85

	switch {
	case cpu < 30:
		fmt.Println("CPU is less than 30%")
	case cpu < 50:
		fmt.Println("CPU is less than 50%")
		fallthrough
	case cpu < 80:
		fmt.Println("CPU is less than 80%")
	default:
		fmt.Println("CRITICAL : CPU is over 80% ")
	}

}
