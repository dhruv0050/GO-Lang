package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Learning to make GET requests")
	PerformGetRequest()
}

func errHandling(err error) {
	if err != nil {
		panic(err)
	}
}

func PerformGetRequest() {
	const myurl = "http://localhost:3000"
	response, err := http.Get(myurl)
	errHandling(err)
	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	// fmt.Println("Content Length: ", response.ContentLength)

	// content, err := io.ReadAll(response.Body)
	// errHandling(err)
	// fmt.Println(string(content))

	var responseString strings.Builder
	content, err := io.ReadAll(response.Body)
	errHandling(err)
	byteCount,_:= responseString.Write(content)
	fmt.Println("ByteCount is ",byteCount)
	fmt.Println(responseString.String())
}
