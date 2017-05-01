package main

import (
	"fmt"
)

func main() {
	var arr [10]int
	arr[0] = 143

	var arr2 = [3]int{1, 3, 5}

	fmt.Printf("first is %d\n", arr[0])
	fmt.Printf("second elem of second array is %d \n", arr2[1])
}
