package main

import (
	"sort"
	"testing"
)

func TestPower(t *testing.T) {
	g := group{}
	g.units = 18
	g.damage = 8

	want := 144
	got := g.power()
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}

func TestSort(t *testing.T) {
	a := army{}
	g0 := group{}
	g0.units = 18
	g0.damage = 1

	g1 := group{}
	g1.units = 1
	g1.damage = 8

	g2 := group{}
	g2.units = 180

	g2.damage = 8

	a.groups = []group{g0, g1, g2}
	sort.Sort(a)
	if a.groups[0].power() != 1440 {
		t.Errorf("want 8 got %d", a.groups[0].power())
	}

	if a.groups[1].power() != 18 {
		t.Errorf("want 18 got %d", a.groups[1].power())
	}

	if a.groups[2].power() != 8 {
		t.Errorf("want 1440 got %d", a.groups[2].power())
	}

}
