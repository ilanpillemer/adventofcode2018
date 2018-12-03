package main

import "testing"

func TestCheckSum(t *testing.T) {
	d := &device{3, 4}
	if d.checksum() != 12 {
		t.Errorf("want %d, got %d", 12, d)
	}
}

func TestScanner(t *testing.T) {
	tests := []struct {
		input  string
		twos   int
		threes int
	}{
		{"abcdef", 0, 0},
		{"bababc", 1, 1},
		{"abbcde", 1, 0},
		{"abcccd", 0, 1},
		{"aabcdd", 1, 0},
		{"abcdee", 1, 0},
		{"ababab", 0, 1},
	}

	for _, test := range tests {
		d := &device{0, 0}
		d.scan(test.input)
		if d.twos != test.twos || d.threes != test.threes {
			t.Errorf("want %d %d, got %d %d for %s", test.twos, test.threes, d.twos, d.threes, test.input)
		}
	}

}