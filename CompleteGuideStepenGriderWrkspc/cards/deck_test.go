package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected length is 16 but we got %v", len(d))
	}
	if len(d) != 2000 {
		t.Errorf("Expected length is 16 but we got %v", len(d))
	}
}
