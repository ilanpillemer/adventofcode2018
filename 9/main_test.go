package main

import "testing"

func TestHighScores(t *testing.T) {
	tests := []struct {
		players    int
		lastMarble int
		want       int
	}{
		{9, 25, 32},
		{10, 1618, 8317},
		{17, 1104, 2764},
		{21, 6111, 54718},
		{30, 5807, 37305},
	}

	for _, test := range tests {
		got := play(test.players, test.lastMarble)
		if test.want != got {
			t.Errorf("want %d got %d", test.want, got)
		}
	}
}
