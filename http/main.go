package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Requesting server...")

	payload, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error feteching API!")
	}
	defer payload.Body.Close()

	data, _ := ioutil.ReadAll(payload.Body)
	fmt.Println(string(data))
}
