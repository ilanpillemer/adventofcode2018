package main

import "testing"

func TestHighScores(t *testing.T) {
	tests := []struct {
		players    int
		lastMarble int
		want       int
	}{
		{10, 1618, 8371},
	}

	for _, test := range tests {
		got := play(test.players, test.lastMarble)
		if test.want != got {
			t.Errorf("want %d got %d", test.want, got)
		}
	}
}