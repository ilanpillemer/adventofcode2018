package main

import "testing"

func TestPower(t *testing.T) {
	group := make([]unit, 18)
	for i := 0; i < 18; i++ {
		group[i] = unit{damage: 8}
	}

	want := 144
	got := power(group)
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}
