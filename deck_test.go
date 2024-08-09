package main

import (
	"fmt"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 16 {
		//	fmt.Println("the test failed")
		t.Errorf("Expected deck length of 20 but got %v", len(cards))
		return
	}
	fmt.Println("The test passed")
}

func TestCards(t *testing.T) {
	cards := newDeck()

	if cards[0] != "Ace of Spades" {
		t.Errorf("expected Ace of spadess but got %v", cards[0])
	}
}
