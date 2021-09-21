package main

// go env -w GO111MODULE=auto

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected 16 cards but got %v", len(d))
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")
	d := newDeck()
	d.saveToFile("_deckTesting")
	nd := newDeckFromFile("_deckTesting")
	if len(nd) != 16 {
		t.Errorf("16 cards not loaded %v", len(nd))
	}
	os.Remove("_deckTesting")
}
