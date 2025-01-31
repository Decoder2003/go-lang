package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Testing my url using package")
	Myurl := "https://jsonplaceholder.typicode.com/comments?postId=9"
	payload, _ := http.Get(Myurl)

	data, _ := ioutil.ReadAll(payload.Body)
	fmt.Println("Response:", string(data))

	// Parsing the url
	ParseUrl, _ := url.Parse(Myurl)
	fmt.Println("Scheme:", ParseUrl.Scheme)
	fmt.Println("Path:", ParseUrl.Path)
	fmt.Println("Host:", ParseUrl.Host)
	fmt.Println("Query:", ParseUrl.RawQuery)

	// creating new url by changing Path & RawQuery
	ParseUrl.Path = "/posts/1"
	ParseUrl.RawQuery = ""
	fmt.Println("New url:", ParseUrl)

	payload, _ = http.Get(ParseUrl.String())

	data, _ = ioutil.ReadAll(payload.Body)
	fmt.Println("New Response:", string(data))
}
