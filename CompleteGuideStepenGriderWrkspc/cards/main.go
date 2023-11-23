package main

import "fmt"

func main() {

	printState()

	//slices example:

	//cards := deck{"ace", "diamond", "heartin"}
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	cards := newDeck()
	cards.print()

	hand, remainingcards := deal(cards, 5)
	hand.print()
	remainingcards.print()

	greeting := "Hi There!"
	fmt.Println([]byte(greeting))

	fmt.Println(cards.toString())

	fmt.Println(cards.saveToFile("myCards"))

	cards = readFromFile("myCards1")
	cards.print()

}
