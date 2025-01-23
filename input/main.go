package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Heyy dhruv!")

	fmt.Print("Enter the 1st number : ")
	var no1 int
	fmt.Scan(&no1)

	fmt.Print("\nEnter the 2nd number : ")
	var no2 int
	fmt.Scan(&no2)

	fmt.Printf("\nThe addition of %d and %d is %d", no1, no2, add(no1, no2))
}
