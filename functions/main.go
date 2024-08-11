package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("Learning function")
    fmt.Printf("\n")

    var x, y int
    fmt.Print("Enter the first num for addition: ")
    fmt.Scanln(&x)

    fmt.Print("Enter the second number for addition: ")
    fmt.Scanln(&y)

    added := add(x, y)
    fmt.Println("After adding we get:", added)
    fmt.Printf("\n")

    var a, b float64
    fmt.Print("Enter the first num for subtraction: ")
    fmt.Scanln(&a)

    fmt.Print("Enter the second number for subtraction: ")
    fmt.Scanln(&b)

    subtracted := sub(a, b)
    fmt.Println("After subtracting we get:", subtracted)
    fmt.Printf("\n")

    var p, q int
    fmt.Print("Enter the first num for multiplication: ")
    fmt.Scanln(&p)

    fmt.Print("Enter the second number for multiplication: ")
    fmt.Scanln(&q)

    multiplied := mul(p, q)
    fmt.Println("After multiplying we get:", multiplied)
}

func add(v1 int, v2 int) int {
    return v1 + v2
}

func sub(v1 float64, v2 float64) float64 {
    return math.Abs(v1 - v2)
}

func mul(v1 int, v2 int) int {
    return v1 * v2
}
