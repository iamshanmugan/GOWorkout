package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuite := deck{"Spade", "Heart", "Diamond", "Clubs"}
	cardValue := deck{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuite {
		for _, val := range cardValue {
			cards = append(cards, val+"  of  "+suit)
		}

	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")

}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 666)
}

func readFromFile(fileName string) deck {

	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}
