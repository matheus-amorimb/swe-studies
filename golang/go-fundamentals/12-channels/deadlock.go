package main

import "fmt"

func deadlock() {
	ch := make(chan int)
	ch <- 1
	v := <-ch
	fmt.Println(v)
}
