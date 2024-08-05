package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Learning Slices")
	fmt.Printf("\n")
	var fruitList = []string{"Apple", "MAngo", "Banana","Peach"}
	fmt.Println("Right Now the list contains :", fruitList)
	fmt.Printf("\n")
	fmt.Printf("Adding more fruits to the list\n")
	fruitList = append(fruitList, "Banana", "Watermelon")
	fmt.Println("After appending the list contains :", fruitList)
	fmt.Printf("\n")
	fmt.Printf("Fetching only fruits from index 2 to 5:\n")
	fmt.Println(fruitList[2:6])
	fmt.Printf("\n")
	var nums = []int{1,2,100,348,113,223,454,3,5}
	fmt.Println("The number list is :",nums)
	fmt.Printf("\n")
	fmt.Println("Sorting this nums list in increasing order:")
	sort.Ints(nums)
	fmt.Println("THe sorted nums list is:",nums)
}