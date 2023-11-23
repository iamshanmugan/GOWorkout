package main

import "fmt"

type bot interface {
	getGreeting() string
}

type engbot struct{}
type spbot struct{}

func main() {
	eb := engbot{}
	sp := spbot{}
	fmt.Println("Interface sample")

	printGreeting(eb)
	printGreeting(sp)
}

func printGreeting(e bot) {
	fmt.Println(e.getGreeting())
}

func (engbot) getGreeting() string {
	return "Hi There"
}
func (spbot) getGreeting() string {
	return "Holla!!"
}
