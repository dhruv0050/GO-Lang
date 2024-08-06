package main

import (
	"fmt"

)

func main() {
	fmt.Println("Learning Maps in Go")
	fmt.Printf("\n")
	fmt.Println("Maping programing languages with their extensions:")
	fmt.Printf("\n")
	languages := make(map[string]string)

	languages["Ruby"] = ".rb"
	languages["Go"] = ".go"
	languages["Python"] = ".py"
	languages["JavaScript"] = ".js"
	languages["Java"] = ".java"
	fmt.Println(languages)
	fmt.Printf("\n")

	fmt.Println("The extension for Ruby is :",languages["Ruby"])
	delete(languages,"Ruby")
	fmt.Println("Deleted Ruby")
	fmt.Printf("\n")
	for key,value := range languages{
		fmt.Printf("For %v, extension is %v\n",key,value)
	}



}