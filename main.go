package main

import (
	"fmt"
)

func main() {
	r1 := Rect{88, 67}
	fmt.Printf("Area of r1 is %f m2\n", area(r1))
	fmt.Printf("Perimetr is %f m\n", r1.perim())
}

/**
* Recatngle
 */
type Rect struct {
	width, height float64
}

func area(r Rect) float64 {
	return r.width * r.height
}

func (r Rect) perim() float64 {
	return r.width*2 + r.height*2
}
