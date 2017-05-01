package main

import (
	"errors"
	"fmt"
)

func main() {
	// var i inte
	err := errors.New("some error")

	if err != nil {
		fmt.Print(err)
	}
}