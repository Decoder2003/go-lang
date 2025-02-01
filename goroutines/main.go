package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, wait for ME!")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Ended..")
}

func sayHi() {
	fmt.Println("Hi Dhruv!")
}

func main() {

	go sayHello()
	sayHi()

	// waiting for func sayHello to execute
	time.Sleep(2000 * time.Millisecond)
}
