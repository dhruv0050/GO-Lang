package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Generating diceNumber of LUDO game using SWITCH-CASE: ")
	fmt.Printf("\n")
	diceNumber := rand.Intn(6) + 1
	fmt.Println("You got ",diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Move 1 spot ahead")
	case 2:
		fmt.Println("Move 2 spots ahead")
	case 3:
		fmt.Println("Move 3 spots ahead")
	case 4:
		fmt.Println("Move 4 spots ahead")
	case 5:
		fmt.Println("Move 5 spots ahead")
	case 6:
		fmt.Println("Move 6 spots ahead and roll dice again")
	default:
		fmt.Println("Soemthing went wrong!")
	}
}