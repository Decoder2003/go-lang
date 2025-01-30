package main

import "fmt"

func mul(a, b int) int {
	return a * b
}

func main() {
	fmt.Println("Code in running.")

	a := 10
	b := 20
	fmt.Println(mul(a, b))

	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println(mul(a, b))
}
