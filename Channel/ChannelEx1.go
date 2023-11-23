package main

import "fmt"

func main() {

	message := make(chan string)

	go func() { message <- "ping channel message" }()
	msg := <-message
	fmt.Println(msg)

}
