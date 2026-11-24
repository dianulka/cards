package main

func main() {
	cards := deck{newCard(), newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six") // returns new slice

	cards.print()
}

func newCard() string {
	return "Five of diamons"
}
