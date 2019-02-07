package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expecting deck length of 16, but got, %d", len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	d := newDeck()
	d.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expecting deck length of 16, but got, %d", len(d))
	}

	os.Remove("_deckTesting")
}