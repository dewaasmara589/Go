package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, World")
}

func main() {
	go sayHello()

	// Goroutines
	time.Sleep(time.Second)

	// Channels
	ch := make(chan string)

	go func() {
		ch <- "Hello, Channel!"
	}()

	fmt.Println(<-ch)
}
