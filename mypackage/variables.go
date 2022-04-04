package main

import (
	"fmt"
)

func main() {
	name := "This is me, Ravi!"
	fmt.Println("From main(): ", name)
	letsupdate(&name)
	fmt.Println("After update ..from main(): ", name)
}

func letsupdate(myname *string) {
	fmt.Println("From letsupdate(): ", *myname)
	*myname = "Updated to Ravikumar Buragapu"
	fmt.Println("Updated Name : ", myname)
	fmt.Println("Updated Name : ", *myname)
}
