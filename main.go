package main

import (
	"fmt"
)

func main() {
	// var numbers map[string]int

	numbers := make(map[string]int)
	numbers["one"] = 1
	fmt.Printf("First number is %d\n", numbers["one"])
}
