package main

import "fmt"

func main() {

	messages := make(chan string)
	signal := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received messags", msg)
	default:
		fmt.Println("Message not received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("received messags", msg)
	default:
		fmt.Println("Message not received")
	}

	select {
	case msg := <-messages:
		fmt.Println("received messags", msg)
	case sig := <-signal:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("Message not received")
	}
}
