package main

import "fmt"


const Float_Number float64 = 300.5004       //Public declaration withe a constant value 
const Name string = "Dhruv"
const Number  = 1983

func main() {
	fmt.Println("Variables")
	//String
	var username string = "Dhruv"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	//Boolean
	var isthere bool = true
	fmt.Println(isthere)
	fmt.Printf("Variable is of type : %T \n", isthere)

	var num int = 400
	fmt.Println(num)
	fmt.Printf("Variable is of type : %T \n", num)

	var float float64 = 07.000567898
	fmt.Println(float)
	fmt.Printf("Variable is of type : %T \n", float)

	var float2 float32 = 07.000567898
	fmt.Println(float2)
	fmt.Printf("Variable is of type : %T \n", float)

	fmt.Printf("The global Float Number is %f\n", Float_Number)
	
	fmt.Printf("The global Name is %s\n", Name)

	fmt.Printf("The global Number is %d\n", Number)
	
}