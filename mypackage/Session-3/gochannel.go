package main

import (
	"fmt"
	"time"
)

func main() {
	mychannel := make(chan int)

	fmt.Println("Sending value to channel>>>")
	go send(mychannel)

	fmt.Println("<<<Receiving from channel")
	go receive(mychannel)

	time.Sleep(time.Second * 1)
}

func send(mychannel chan int) {
	mychannel <- 1
}

func receive(mychannel chan int) {
	val := <-mychannel
	fmt.Printf("Channel value received = %d in receive function\n", val)
}
