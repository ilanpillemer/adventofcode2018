package main

import (
	"testing"
)

var locations = []pos{
	pos{1, 1},
	pos{1, 6},
	pos{8, 3},
	pos{3, 4},
	pos{5, 5},
	pos{8, 9},
}

var g = grid{locs: locations}

func TestFrame(t *testing.T) {
	g.init()
	if g.top != 1 {
		t.Errorf("top of grid... want %d, got %d", 1, g.top)
	}

	if g.bottom != 9 {
		t.Errorf("bottom of grid... want %d, got %d", 9, g.bottom)
	}

	if g.left != 1 {
		t.Errorf("left of grid... want %d, got %d", 1, g.left)
	}

	if g.right != 8 {
		t.Errorf("right of grid... want %d, got %d", 8, g.right)
	}
}

func TestManhattan(t *testing.T) {

	tests := []struct {
		a    pos
		b    pos
		want int
	}{
		{pos{0, 0}, pos{3, 4}, 7},
		{pos{5, 0}, pos{8, 3}, 6},
		{pos{5, 0}, pos{5, 5}, 5},
		{pos{5, 0}, pos{1, 1}, 5},
	}

	for _, test := range tests {
		if got := manh(test.a, test.b); got != test.want {
			t.Errorf("want:%d got:%d", test.want, got)
		}
	}

}

func TestClosest(t *testing.T) {
	locs := []pos{pos{x: 1, y: 1}, pos{x: 1, y: 6}, pos{x: 8, y: 3}, pos{x: 3, y: 4}, pos{x: 5, y: 5}, pos{x: 8, y: 9}}
	tests := []struct {
		locs []pos
		p    pos
		want int
	}{
		{locs, pos{0, 0}, 0},                       //A
		{locs, pos{5, 0}, -1},                      //.
		{locs, pos{5, 1}, -1},                      //.
		{locs, pos{5, 2}, 4},                       //E
		{locs, pos{5, 3}, 4},                       //E
		{locs, pos{5, 4}, 4}, {locs, pos{6, 5}, 4}, //E
		{locs, pos{4, 5}, 4}, {locs, pos{5, 5}, 4}, {locs, pos{6, 5}, 4}, {locs, pos{7, 5}, 4}, //E
		{locs, pos{5, 6}, 4}, //E
		{locs, pos{5, 7}, 4}, //E
		{locs, pos{5, 8}, 4}, //E
	}
	for _, test := range tests {
		got, _ := closest(test.p, test.locs)
		if got != test.want {
			t.Errorf("want :%d got :%d", test.want, got)
		}
	}
}

func TestArea(t *testing.T) {
	g.init()
	g.project()
	g.getLargestNonInfiniteArea()
}

func TestProjected(t *testing.T) {
	g.init()
	g.project()

	tests := []struct {
		x    int
		y    int
		want int
	}{
		{5, 1, -1},
		{5, 2, 4},
	}
	g.print()
	for _, test := range tests {
		got := g.getOwner(test.x, test.y)
		if got != test.want {
			//		t.Errorf("Want %d, got %d", test.want, got)

		}
	}
}

//1, 1 A 0
//1, 6 B 1
//8, 3 C 2
//3, 4 D 3
//5, 5 E 4
//8, 9 F 5
