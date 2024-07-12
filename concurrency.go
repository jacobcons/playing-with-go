package main

import (
	"fmt"
	"time"
)

func attack(target string, channel chan bool) {
	fmt.Println(<-channel)
	channel <- true
}

func main() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()

	channel := make(chan string, 2)
	channel <- "First message"
	channel <- "second"
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
