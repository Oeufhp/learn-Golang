package main

import "fmt"

type deck []string

func newDeck() deck {
	cardSuits := []string{"Spades", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Jack", "Queen", "King"}

	cards := deck{}

	for _, val := range cardValues {
		for _, suit := range cardSuits {
			cards = append(cards, val+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
