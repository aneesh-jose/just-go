package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	d := make(chan int)
	go routine1(c, d)
	go routine2(d)
	fmt.Println(<-c)
}

func routine1(x chan int, y chan int) {
	x <- <-y
}
func routine2(x chan int) {
	x <- 1
	time.Sleep(2 * time.Second)
}
