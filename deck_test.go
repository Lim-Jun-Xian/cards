package main

import "testing"

// Test the function newDeck().
func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Test whether function consist of 16 cards
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// Test whether first card is Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	// Test whether last card is Four of Clubs
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}
