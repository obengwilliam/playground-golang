package main

import "fmt"

func main() {
	buffered := make(chan string, 3)
	buffered <- "1"
	buffered <- "2"

	// basically it does not block when sending messages
	fmt.Println("third buffer coming in \n ")
	buffered <- "3"

	// it seems close can go anywhere
	// close channel
	close(buffered)

	// app panices when u create a channel that takes more than it required
	//buffered <- "4"

	fmt.Println(<-buffered, "1...")
	fmt.Println(<-buffered, "2...")
	fmt.Println(<-buffered, "3...")
}
