package main

import (
	"fmt"
	"testing"
)

var c1 = &claim{
	id:     "1",
	left:   1,
	top:    3,
	width:  4,
	height: 4,
}

var c2 = &claim{
	id:     "2",
	left:   3,
	top:    1,
	width:  4,
	height: 4,
}

var c3 = &claim{
	id:     "3",
	left:   5,
	top:    5,
	width:  2,
	height: 2,
}

func TestHello(t *testing.T) {
	// no-op
}

func TestClaim(t *testing.T) {
	c := NewClaim("#1 @ 1,3: 4x4")

	if c.id != "#1" {
		t.Errorf("want %s, got %s", "#1", c.id)
	}

	if c.left != 1 {
		t.Errorf("want %d, got %d", 1, c.left)
	}

	if c.top != 3 {
		t.Errorf("want %d, got %d", 3, c.top)
	}

	if c.width != 4 {
		t.Errorf("want %d, got %d", 4, c.width)
	}

	if c.height != 4 {
		t.Errorf("want %d, got %d", 4, c.height)
	}
}

func TestOverlapCells(t *testing.T) {
	cells := overlapCells(c1, c2)
	if fmt.Sprint(cells) != "[{3 3} {3 4} {4 3} {4 4}]" {
		t.Errorf("want %s, got %s", "[{3 3} {3 4} {4 3} {4 4}]", fmt.Sprint(cells))
	}

}

func TestOverlapSize(t *testing.T) {
	if overlapSize(c1, c2) != 4 {
		t.Errorf("want %d, got %d", 4, overlapSize(c1, c2))
	}

	if overlapSize(c2, c1) != 4 {
		t.Errorf("want %d, got %d", 4, overlapSize(c2, c1))
	}
}

func TestOverlap(t *testing.T) {

	if !overlap(c1, c2) {
		t.Error("c1 and c2 should overlap")
	}

	if !overlap(c2, c1) {
		t.Error("c2 and c1 should overlap")
	}

	if overlap(c2, c3) {
		t.Errorf("c2 [%v] and c3[%v] should not overlap", c2, c3)
	}

	if overlap(c3, c2) {
		t.Errorf("c3 [%v] and c2[%v] should not overlap", c3, c2)
	}

	if overlap(c1, c3) {
		t.Errorf("c2 [%v] and c3[%v] should not overlap", c2, c3)
	}

	if overlap(c3, c1) {
		t.Errorf("c3 [%v] and c2[%v] should not overlap", c3, c2)
	}

}

//id   l t  w h
//#1 @ 1,3: 4x4
//#2 @ 3,1: 4x4
//#3 @ 5,5: 2x2
