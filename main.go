package main

import "fmt"

func main() {
	// card := newCard()
	cards := []string{newCard(), newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six") // returns new slice

	for i, card := range cards {
		fmt.Println(i, card)
	}

	for _, card := range cards {
		fmt.Println(card)
	}
}

func newCard() string {
	return "Five of diamons"
}
