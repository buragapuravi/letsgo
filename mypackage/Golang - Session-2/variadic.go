package main

import "fmt"

func main() {
	maxcpu := maxcpu(22, 45, 33, 29, 56, 68)
	fmt.Println("Max CPU :", maxcpu)
}

func maxcpu(cpu ...int) int {
	high := cpu[0]
	for _, i := range cpu {
		if i > high {
			high = i
		}
	}
	return high
}
