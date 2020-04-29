package main

import (
	"fmt"
	"time"
)

func from(from string) {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d %s \n", i, from)
	}

	// the job of this is to delay enough so that all the printf for the gorountine can show up
	time.Sleep(time.Second)
}

func main() {
	go from("goroutine")
	from("shit")
}
