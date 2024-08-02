package main

import "fmt"

func main() {
	fmt.Println("Learning Pointers")
	fmt.Print("\n")
	myNumber := 7
	var ptr = &myNumber
	fmt.Println("Address of pointer is :",ptr)
	fmt.Println("Value of pointer is :",*ptr)
	fmt.Print("\n")

	fmt.Println("Applying mathematics to the pointer")
	fmt.Print("\n")
	*ptr = *ptr * 3
	fmt.Println("Value of pointer after multiplication :",myNumber)
	*ptr = *ptr +2
	fmt.Println("Value of pointer after addition :",myNumber)
	*ptr = *ptr -20
	fmt.Println("Value of pointer after subtraction :",myNumber)
}