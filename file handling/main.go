package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Learning File handling in GO")
	content := "Lorem ipsum dolor, sit amet consectetur adipisicing elit. Odio eos iure fugiat deserunt quasi, amet accusamus rerum modi ab. Consequatur enim eligendi veritatis ducimus omnis praesentium aliquid itaque tempore reprehenderit?"

	file, err := os.Create("./random_file.txt")
	if err != nil {
		panic(err)
	}
	_, err = io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	readFile("./random_file.txt")

	fmt.Println("File writing completed successfully!")
}

func readFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("The data of the file is\n", string(data))
}
