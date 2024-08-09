package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learning if-else")
	var age int

	fmt.Print("Enter your current age: ")
	_, err := fmt.Scan(&age)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if age>=18 {
		fmt.Println("Eligible to vote")
	}else{
		fmt.Println("Not eligible to vote")
	}
}