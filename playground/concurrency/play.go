package main

import (
	"fmt"
)

var message chan string = make(chan string, 3)

func main() {
	message <- "shit"
	message <- "good"
	message <- "good"

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}

// var done chan bool = make(chan bool)

// func printMessages(source string) {

// 	for i := 0; i < 9; i++ {
// 		fmt.Println("MEssage", i, "from", source)
// 	}

// 	if source == "goroutine" {
// 		done <- true
// 	}
// }

// func main() {

// 	go printMessages("goroutine")
// 	<-done
// 	printMessages("slow")
// }

// func getMessage(message chan string) {
// 	message <- "Good boy"
// }

// func main() {
// 	message := make(chan string)

// 	go getMessage(message)
// 	fmt.Println("Sample message", <-message)
// 	close(message)
// }
