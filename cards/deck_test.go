package main

import (
	"os"
	"testing"
)

// t -> test handler of type testing.T
func TestNewDeck(t *testing.T) {
	d := newDeck()

	// if we have less cards than the 56 needed
	if len(d) != 56 {
		// created a formatted string -> must inject value into string
		t.Errorf("Expected deck length of 56 but got: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got: %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got: %v", d[len(d)-1])
	}
}

//
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// remove old file in case there was a previous failure
	os.Remove("_decktesting")

	// create a new deck and save
	deck := newDeck()
	deck.saveToFile("_decktesting")

	// load prev deck
	loadedDeck := newDeckFromFile("_decktesting")

	// test length of loaded deck = average deck length
	if len(loadedDeck) != 56 {
		t.Errorf("Expected 56 cards in deck, but got %v", len(loadedDeck))
	}

	// remove old deck
	os.Remove("_decktesting")
}
