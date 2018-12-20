package main

import (
	"fmt"
	"testing"
)

func TestGrid8(t *testing.T) {

	tests := []struct {
		g    *grid
		x    int
		y    int
		want int
	}{
		{NewGrid(8), 3, 5, 4},
		{NewGrid(57), 122, 79, -5},
		{NewGrid(39), 217, 196, 0},
		{NewGrid(71), 101, 153, 4},
	}

	for _, test := range tests {
		got := test.g.Power(test.x, test.y)
		if got != test.want {
			t.Errorf("want %d got %d", test.want, got)
		}
	}
}

func TestPowerSquare(t *testing.T) {
	tests := []struct {
		g    *grid
		x    int
		y    int
		size int
		want int
	}{
		{NewGrid(18), 33, 45, 3, 29},
		{NewGrid(42), 21, 61, 3, 30},
		{NewGrid(18), 90, 269, 16, 113},
		{NewGrid(42), 232, 251, 12, 119},
	}

	for _, test := range tests {
		got := test.g.PowerSquare(test.x, test.y, test.size)
		if got != test.want {
			t.Errorf("want %d got %d", test.want, got)
		}
	}
}

func TestMaxPower(t *testing.T) {
	tests := []struct {
		g     *grid
		wantx int
		wanty int
	}{
		{NewGrid(18), 33, 45},
		{NewGrid(42), 21, 61},
	}

	for _, test := range tests {
		_, gotx, goty := test.g.MaxPower(3)
		if gotx != test.wantx || goty != test.wanty {
			t.Errorf("want %d,%d got %d,%d", test.wantx, test.wanty, gotx, goty)
		}
	}
}

func TestMaxPowerAllSizes(t *testing.T) {
	tests := []struct {
		g        *grid
		wantx    int
		wanty    int
		wantsize int
	}{
		{NewGrid(18), 90, 269, 16},
		{NewGrid(42), 232, 251, 12},
	}

	for _, test := range tests {
		_, gotx, goty, gotsize := test.g.MaxPowerAllSize()
		if gotx != test.wantx || goty != test.wanty || gotsize != test.wantsize {
			t.Errorf("want %d,%d,%d got %d,%d,%d", test.wantx, test.wanty, test.wantsize, gotx, goty, gotsize)
		}
	}
}

func TestStuff(t *testing.T) {
	got := hundreds(12345)
	want := 3
	if got != want {
		t.Errorf("want %d got %d", want, got)
	}

	g := NewGrid(18)
	fmt.Printf("%d\t%d\t%d\t%d \n", g.cells[0][0], g.cells[0][1], g.cells[0][2], g.cells[0][3])
	fmt.Printf("%d\t%d\t%d\t%d \n", g.cells[1][0], g.cells[1][1], g.cells[1][2], g.cells[1][3])
	fmt.Printf("%d\t%d\t%d\t%d \n", g.cells[2][0], g.cells[2][1], g.cells[2][2], g.cells[2][3])
	fmt.Println()
	fmt.Printf("%d\t%d\t%d\t%d \n", g.sat[0][0], g.sat[0][1], g.sat[0][2], g.sat[0][3])
	fmt.Printf("%d\t%d\t%d\t%d \n", g.sat[1][0], g.sat[1][1], g.sat[1][2], g.sat[1][3])
	fmt.Printf("%d\t%d\t%d\t%d \n", g.sat[2][0], g.sat[2][1], g.sat[2][2], g.sat[2][3])
	fmt.Println()
	fmt.Printf("%d\t%d\t%d \n", g.cells[32][44], g.cells[32][45], g.cells[32][46])
	fmt.Printf("%d\t%d\t%d \n", g.cells[33][44], g.cells[33][45], g.cells[33][46])
	fmt.Printf("%d\t%d\t%d \n", g.cells[34][44], g.cells[34][45], g.cells[34][46])
	fmt.Println()
	fmt.Printf("%d\t%d\t%d \n", g.sat[32][44], g.sat[32][45], g.sat[32][46])
	fmt.Printf("%d\t%d\t%d \n", g.sat[33][44], g.sat[33][45], g.sat[33][46])
	fmt.Printf("%d\t%d\t%d \n", g.sat[34][44], g.sat[34][45], g.sat[34][46])

}

//Fuel cell at  122,79, grid serial number 57: power level -5.
//Fuel cell at 217,196, grid serial number 39: power level  0.
//Fuel cell at 101,153, grid serial number 71: power level  4.
