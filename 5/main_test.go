package main

import "testing"

func TestReact(t *testing.T) {
	tests := []struct {
		input polymer
		want  polymer
	}{
		{polymer("Aa"), polymer("")},
		{polymer("abBA"), polymer("")},
		{polymer("abAB"), polymer("abAB")},
		{polymer("aabAAB"), polymer("aabAAB")},
		{polymer("dabAcCaCBAcCcaDA"), polymer("dabCBAcaDA")},
	}
	for _, test := range tests {
		if test.input.react() != test.want {
			t.Errorf("want %s got %s", test.want, test.input.react())
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		input string
		want  polymer
	}{
		{"a", polymer("dbcCCBcCcD")},
		{"B", polymer("daAcCaCAcCcaDA")},
		{"c", polymer("dabAaBAaDA")},
		{"D", polymer("abAcCaCBAcCcaA")},
	}

	p := polymer("dabAcCaCBAcCcaDA")

	for _, test := range tests {
		if p.remove(test.input) != test.want {
			t.Errorf("want %s got %s", test.want, p.remove(test.input))
		}
	}
}

func TestUnits(t *testing.T) {
	p := polymer("dabAcCaCBAcCcaDA").react()
	if p.units() != 10 {
		t.Errorf("want 10 got %d", p.units())
	}
}