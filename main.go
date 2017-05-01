package main

import (
	"fmt"
)

func main() {
	var arr [10]int
	arr[0] = 143

	var arr2 = [3]int{1, 3, 5}
	var slice []int
	slice = arr2[:2]
	fmt.Printf("%d\n", len(slice)) //1
}
