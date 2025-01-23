package main

import "fmt"

func main() {
	fmt.Println("First time declaring an array!")
	arr := [...]int{1, 2, 3, 4, 5} // implicit type variable declaration
	fmt.Println("Orginal array :", arr)
}
