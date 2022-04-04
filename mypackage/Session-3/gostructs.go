package main

import "fmt"

type employee struct {
	name  string
	dept  string
	empno int
}

var (
	emp1 = employee{"Ravikumar Buragapu", "ACT", 123}
	emp2 = employee{name: "David Lee"}
	emp3 = employee{}
	emp  = &employee{"Tom Hanks", "Hollywood", 345}
)

func main() {
	fmt.Println(emp1, emp2, emp3, emp)
}
