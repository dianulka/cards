package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Expected Spades of Ace, but got %v", d[0])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_deckTesting"
	os.Remove(filename)

	deck := newDeck()

	deck.saveToFile(filename)
	newDeck := newDeckFromFile(filename)

	if len(newDeck) != len(deck) {
		t.Errorf("Expected the same length")
	}

	if deck[0] != newDeck[0] {
		t.Errorf("Expected the same card")
	}

	os.Remove(filename)

}
