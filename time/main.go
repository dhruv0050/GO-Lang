package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println("Present Time is :")
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	//Creating time manually
	fmt.Println("Created Time is: ")
	createdTime := time.Date(1983 , time.June , 25 , 22 , 0 , 0 , 0, time.UTC)
	fmt.Println(createdTime)
	fmt.Println(createdTime.Format("01-02-2006 15:04:05 Monday"))
}