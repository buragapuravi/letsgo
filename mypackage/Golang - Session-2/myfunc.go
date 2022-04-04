package main

import (
	"fmt"
	"strings"
)

func main() {
	fname := "Ravi"
	lname := "Buragapu"
	location := "san jose"
	jod := "05 / 24 / 2021"
	empno := 007
	fmt.Println(demofunc(fname, lname, location, jod, empno))
}

func demofunc(fname, lname, location, jod string, empno int) (a, b, c, d string, e int) {
	fname = strings.ToUpper(fname)
	lname = strings.ToUpper(lname)
	location = strings.Title(location)
	return fname, lname, location, jod, empno
}
