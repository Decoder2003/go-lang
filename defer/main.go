package main

import "fmt"

func main() {
	defer fmt.Println("This will execute 1")

	fmt.Println("This will in 2")

	fmt.Println("This will be at 3")
}
