package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected 16 but got %v", len(d))
	}

	if d[0] != "Ace of Heart" {
		t.Errorf("Expected Ace of Heart but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Club" {
		t.Errorf("Expected Four of Club but got %v", d[len(d)-1])
	}
}
