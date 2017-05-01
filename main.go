package main

import (
	"fmt"
)

func main() {
	type person struct {
		name string
		age  int
	}

	type student struct {
		person
		speciality string
	}

	mark := student{person{"Mark", 18}, "Marketolog"}
	fmt.Printf("%s\n", mark.speciality)
}
