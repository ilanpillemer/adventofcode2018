package main

import (
	"testing"
)

func TestMe(t *testing.T) {

	tests := []struct {
		lic  string
		want int
	}{
		{"0 3 10 11 12", 10 + 11 + 12},
		{"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2", 138},
		{"0 1 99", 99},
		{"1 1 0 1 99 2", 101},
	}

	for _, test := range tests {
		got := sumMetadata(test.lic)
		if test.want != got {
			t.Errorf("want %d got %d\n", test.want, got)
		}
	}

}