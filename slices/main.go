package main

import "fmt"

func main() {
	arr := [6]int{1, 2, 3, 4, 5} // array

	// this is my first code to declare a slice
	slice := []int{1, 2, 3, 4, 5}
	// arr = append(arr, 6)
	slice = append(slice, 6)

	fmt.Println(arr)
	fmt.Println(slice)
}
