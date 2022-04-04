package main

import (  
    "fmt"
)


func main() {  
    var a [3]int //int array with length 3
    fmt.Println(a)
}

func main() {  
    var a [3]int //int array with length 3
    a[0] = 12 // array index starts at 0
    a[1] = 78
    a[2] = 50
    fmt.Println(a)
}

func main() {  
    a := [3]int{12, 78, 50} // short hand declaration to create array
    fmt.Println(a)
}

//It is not necessary that all elements in an array have to be assigned a value 
//during short hand declaration.

func main() {  
    a := [3]int{12} 
    fmt.Println(a)
}

// Dynamic array
func main() {  
    a := [...]int{12, 78, 50} // ... makes the compiler determine the length
    fmt.Println(a)
}

// Multi dimensional arrays
var b [3][2]string
b[0][0] = "apple"
b[0][1] = "samsung"
b[1][0] = "microsoft"
b[1][1] = "google"
b[2][0] = "AT&T"
b[2][1] = "T-Mobile"
fmt.Printf("\n")
