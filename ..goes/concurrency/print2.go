package main

import (
	"fmt"
)

var done = make(chan bool)

func printNames(source string) {

	for i := 0; i < 9; i++ {
		fmt.Println(source, "index", i)
	}
	if source == "goroutine" {
		done <- true
	}
}

func main() {

	go printNames("goroutine")
	printNames("main")
	<-done
}
