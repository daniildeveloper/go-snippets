package main

import (
	"fmt"
)

func main() {
	x := 3
	y := 4
	a, b := sumAndProduct(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, a)
	fmt.Printf("%d * %d = %d\n", x, y, b)
}

func sumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}
