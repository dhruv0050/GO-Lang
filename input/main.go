package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	println("This is our rating page")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give us a rating : ")

	input,_ := reader.ReadString('\n')
	fmt.Println("Thanks for a",input)
}