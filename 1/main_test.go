package main

import "testing"

func TestHello(t *testing.T) {
	//no-op
}

func TestApplyFrequency(t *testing.T) {
	f := &freq{0}
	f.apply(+1)
	f.apply(-2)
	f.apply(+3)
	f.apply(+1)

	if f.value != 3 {
		t.Errorf("want:%d got %d", 3, f.value)
	}
}