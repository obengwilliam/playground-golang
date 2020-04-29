package main

import "fmt"

func main() {
	message := make(chan string)

	message <- "hello"
	// message <- "shello"
	// message <- "sello"

	close(message)

	// 	for value := range message {
	// 		fmt.Println(value)
	// 	}

	fmt.Println(<-message)
	// fmt.Println(<-message)
	// fmt.Println(<-message)
}
