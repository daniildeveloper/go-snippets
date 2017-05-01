package main

import (
	"fmt"
	"runtime"
)

func main() {
	go say("world")
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Printf("%s \n", s)
	}
}
