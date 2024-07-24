package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	println("This is our rating page")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give us a rating : ")

	input,_ := reader.ReadString('\n')
	fmt.Println("Thanks for a",input)

	num , err := strconv.ParseFloat(strings.TrimSpace(input),64)
	if err!= nil{
		fmt.Println((err))
	} else{
		fmt.Println("Adding 1 to you rating" , num+1)
	}
}