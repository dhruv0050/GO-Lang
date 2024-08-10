package main

import "fmt"

func main() {
    fmt.Println("Learning loop, break, continue, goto in Go")
    fmt.Printf("\n")
    
    days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
    
    for index, day := range days {
        fmt.Printf("Day %v is %v\n", index+1, day)
    }
	fmt.Printf("\n")
    
    value := 1
    for value < 15 {
        if value == 12 {
			fmt.Println("Value is 12 we need to break now")
            break
        }
        if value == 3 {
            value++
			fmt.Println("Skipped 3")
            continue
        }
        if value == 11 {
			fmt.Println("Processing goto stattement")
            goto random
        }
        fmt.Printf("Loop At %v\n", value)
        value++
    }
    
random:
    fmt.Println("Hello GO")
}
