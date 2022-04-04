package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//Create file
	file, err := os.Create("dummyfile2.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file created")
	//raise error
	myfile, err := os.Open("filename.foo")
	if err != nil {
		log.Fatal(err)
		defer myfile.Close()
	}
	defer myfile.Close()

}
