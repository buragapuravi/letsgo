package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var myRoutinesWaitGrp sync.WaitGroup
	myRoutinesWaitGrp.Add(2)

	go func() {
		defer myRoutinesWaitGrp.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("First Routine in Execution")
	}()

	go func() {
		defer myRoutinesWaitGrp.Done()
		fmt.Println("Second Routine in Execution")
	}()
	myRoutinesWaitGrp.Wait()
}
