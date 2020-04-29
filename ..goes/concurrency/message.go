package main

import "fmt"

func getMessage(message chan string) {
	message <- "channel body"
}
func main() {
	message := make(chan string)
	go getMessage(message)
	fmt.Println("Hello message:", <-message)
	close(message)
}
