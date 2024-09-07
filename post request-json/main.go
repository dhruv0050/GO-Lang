package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Learning the concept of Post request with JSON data")
	PerformPostJsonRequest()
}

func errHandling(err error) {
	if err != nil {
		panic(err)
	}
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:3000/post"
	requestBody := strings.NewReader(`
	{
		"cname" : "GOLANG",
		"Code editor":"VS CODE"
	}`)
	
	response, err := http.Post(myurl, "application/json", requestBody)
	errHandling(err)
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	errHandling(err)

	fmt.Println(string(content))
}
