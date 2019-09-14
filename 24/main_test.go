package main

import (
	"sort"
	"testing"
)

func TestPower(t *testing.T) {
	g := group{}
	g.units = make([]unit, 18)
	for i := 0; i < 18; i++ {
		g.damage = 8
		g.units[i] = unit{}
	}

	want := 144
	got := g.power()
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}

func TestSort(t *testing.T) {
	a := army{}
	g0 := group{}
	g0.units = make([]unit, 18)
	for i := 0; i < 18; i++ {
		g0.damage = 1
		g0.units[i] = unit{}
	}

	g1 := group{}
	g1.units = make([]unit, 1)
	for i := 0; i < 1; i++ {
		g1.damage = 8
		g1.units[i] = unit{}
	}

	g2 := group{}
	g2.units = make([]unit, 180)
	for i := 0; i < 180; i++ {
		g2.damage = 8
		g2.units[i] = unit{}
	}

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
