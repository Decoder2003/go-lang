package main

import "fmt"

func main() {
	// table of 5
	for i := 1; i <= 10; i++ {
		fmt.Printf("5 x %d = %d\n", i, i*5)
	}
}
