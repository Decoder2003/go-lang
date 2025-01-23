package main

import "fmt"

func menu() {
	fmt.Println("-------Menu------")
	fmt.Println("1. add")
	fmt.Println("2. sub")
	fmt.Println("3. multiply")
	fmt.Println("4. division")
	fmt.Println("5. end")
}

func main() {

	fmt.Println("\nThis is my calculator")

	
	var number int
	fmt.Print("Enter the number : ")
	fmt.Scan(&number)
	fmt.Println(number)

	menu()
}
