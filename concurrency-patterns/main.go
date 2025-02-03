package main

// 1. go routines : fork-main model
// 2. channels

import (
	"fmt"
)

func sayHi(x int) {
	fmt.Println(x)
}

func main() {
	go sayHi(1)
	go sayHi(2)
	go sayHi(3)

	// time.Sleep(time.Second * 1)
	fmt.Println("Hello there!")

	myChannel := make(chan string)
	go func() {
		myChannel <- "data"
	}()

	fmt.Println(<-myChannel)
}
