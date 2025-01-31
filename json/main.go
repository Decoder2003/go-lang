package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	IsAdult bool
}

func main() {
	fmt.Println("Using json package for the frst time")
	P := Person{Name: "john", Age: 34, IsAdult: true}
	Data, _ := json.Marshal(P)

	fmt.Println(string(Data))
	data := json.Unmarshal(Data)
	fmt.Println(string(Data))

	fmt.Println(P)
}
