package main

import (
	"fmt"
	"time"
)

func concurrency(message string) {
	//This anonymous function is executed concurrently
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}
