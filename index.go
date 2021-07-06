package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	argsWithProg := os.Args[1:4]
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	go speed(c1, argsWithProg[0])
	go speed(c2, argsWithProg[1])
	go speed(c3, argsWithProg[1])

	select {
	case s1 := <-c1:
		fmt.Println(s1)
	case s2 := <-c2:
		fmt.Println(s2)
	case s3 := <-c3:
		fmt.Println(s3)
	}
}

func speed(ch chan string, page string) {
	http.Get(page)
	ch <- page
}
