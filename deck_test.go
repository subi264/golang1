package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 2000 {
		t.Errorf("expected deck length of 20 ,but got %v", len(d))

	}

}
