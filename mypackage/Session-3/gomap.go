package main

import "fmt"

func main() {
	emp := map[int]string{
		123: "Ravikumar Buragapu",
		345: "Dhoni MS",
		567: "Kohli Virat",
	}
	fmt.Println(emp[123])
	//Add a key
	emp[789] = "Ben Stokes"
	fmt.Println(emp)
	//Delete a key
	delete(emp, 345)
	fmt.Println(emp)
	//update a new key
	emp[789] = "Shakib AI Hasan"
	fmt.Println(emp)

	//Check for a key existed or not
	checkkey, mybool := emp[1234]
	fmt.Println(checkkey, mybool)
	fmt.Println(emp)
}
