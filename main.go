package main

import (
	"fmt"
)

func main() {
	type person struct {
		name string
		age  int
	}
	var P person
	P.name = "Daniil"
	P.age = 21
	fmt.Printf("My name is %s\n", P.name)
	Mark := person{age: 21, name: "Mark"}
	fmt.Printf("My name is %s\n", Mark.name)

}
