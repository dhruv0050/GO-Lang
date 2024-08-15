package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://go.dev"
func main() {
	fmt.Println("Handling web request")
	res , err :=http.Get(url)
	if err!=nil{
		panic(err)
	}
	data , err := io.ReadAll(res.Body)
	if err!=nil{
		panic(err)
	}
	content :=string(data)
	fmt.Println(content)
	defer res.Body.Close()
	fmt.Println("Data Fetched Successfully")
}