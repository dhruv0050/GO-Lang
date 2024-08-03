package main

import "fmt"

func main() {
	var myList1 [4]string
	myList1[0] = "GO"
	myList1[1] = "Python"
	myList1[2] = "C++"
	myList1[3] = "C"

	fmt.Println("The programming languages i know are: ", myList1)
	fmt.Println("The length of myList1 is: ", len(myList1))
	fmt.Printf("\n")

	var list2 = [7]string{"HTML5", "CSS3", "JAVASCRIPT", "MONGODB", "EXPRESS", "REACT", "NODEJS"}
	fmt.Println("The Web Development includes: ", list2)
	fmt.Println("The length of list2 is: ", len(list2))
	fmt.Printf("\n")

	var myList3 [4]string
	copy(myList3[:], myList1[:])
	fmt.Println("The myList3 is: ", myList3)
	fmt.Println("The length of myList3 is: ", len(myList3))

}
