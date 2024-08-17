package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://www.linkedin.com:3000/dhruv-sharma-331379154?GO=Developer"

func main() {
	fmt.Println("Learning Handling URL in GO")
	fmt.Println("The testing url is: ",myUrl)
	result, _ := url.Parse(myUrl)
	fmt.Println("Scheme is: ",result.Scheme)
	fmt.Println("Host is: ",result.Host)
	fmt.Println("Path is: ",result.Path)
	fmt.Println("Port is: ",result.Port())
	fmt.Println("RawQuery is: ",result.RawQuery)
	fmt.Printf("\n")
	fmt.Printf("Making URL from components: \n")
	parts :=&url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tut.css",
		RawPath: "user=dhruv",
	}
	anotherURL :=parts.String()
	fmt.Println("The url is: ",anotherURL)

}