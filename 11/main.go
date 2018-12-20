package main

import (
	"fmt"
)

type grid struct {
	sn    int
	cells [300][300]int
	sat   [300][300]int
}

func (g *grid) sumAreaTable() {
	for x, c := range g.cells {
		for y, _ := range c {
			if x == 0 && y == 0 {
				g.sat[x][y] = g.cells[0][0]
				continue
			}
			if x == 0 {
				g.sat[x][y] = g.liney(0, y)
				continue
			}
			if y == 0 {
				g.sat[x][y] = g.linex(x, 0)
				continue
			}
			g.sat[x][y] = g.sat[x-1][y-1] + g.linex(x-1, y) + g.liney(x, y-1) + g.cells[x][y]
		}
	}
}

func (g *grid) linex(x, y int) int {
	sum := 0
	for i := 0; i < x+1; i++ {
		sum += g.cells[i][y]
	}
	return sum
}

func (g *grid) liney(x, y int) int {
	sum := 0
	for i := 0; i < y+1; i++ {
		sum += g.cells[x][i]
	}
	return sum
}

func (g *grid) Power(x, y int) int {
	return g.cells[x-1][y-1]
}

func (g *grid) PowerSquare(x, y, size int) int {

	x = x - 1 - 1
	y = y - 1 - 1

	whole := g.sat[x+size][y+size]
	//fmt.Println(x, y+size)
	left := g.sat[x][y+size]
	right := g.sat[x+size][y]
	lr := g.sat[x][y]
	return whole + lr - left - right
}

func (g *grid) MaxPower(size int) (int, int, int) {

	max := 0
	i := -1
	j := -1
	for x, c := range g.cells {
		if x+1-1-1 < 0 {
			continue
		}
		if x+1-1-1+size > 299 {
			continue
		}
		for y, _ := range c {
			if y+1-1-1 < 0 {
				continue
			}
			if y+1-1-1+size > 299 {
				continue
			}
			//fmt.Println("powersquare", x+1, y+1, size)
			pwr := g.PowerSquare(x+1, y+1, size)
			if pwr > max {
				max = pwr
				i = x + 1
				j = y + 1
			}
		}
	}

	return max, i, j
}

func (g *grid) MaxPowerAllSize() (int, int, int, int) {
	max := 0
	i := -1
	j := -1
	size := -1
	for k := 0; k < 300; k++ {
		//fmt.Println(k)
		for x, c := range g.cells {
			if x+1-1-1 < 0 {
				continue
			}
			if x+1-1-1+k > 299 {
				continue
			}
			for y, _ := range c {
				if y+1-1-1 < 0 {
					continue
				}
				if y+1-1-1+k > 299 {
					continue
				}
				pwr := g.PowerSquare(x+1, y+1, k)
				if pwr > max {
					max = pwr
					i = x + 1
					j = y + 1
					size = k
				}
			}
		}
	}
	return max, i, j, size
}

func NewGrid(sn int) *grid {
	g := grid{
		sn: sn,
	}

	var cells [300][300]int
	for x, c := range cells {
		for y, _ := range c {
			cells[x][y] = g.calc(x+1, y+1)
		}
	}
	g.cells = cells
	g.sumAreaTable()
	return &g
}

func (g *grid) calc(x, y int) int {

	rackid := x + 10
	begin := rackid * y
	increased := begin + g.sn
	multiplied := increased * rackid
	hundreded := hundreds(multiplied)
	minused := hundreded - 5
	return minused
}

func main() {
	_, x, y := NewGrid(8444).MaxPower(3)
	fmt.Printf("Max Power for Grid is :%d,%d\n", x, y)
	size := -1
	_, x, y, size = NewGrid(8444).MaxPowerAllSize()
	fmt.Printf("Max All Possible Powers for Grid is :%d,%d,%d\n", x, y, size)
}

func hundreds(n int) int {
	return (n / 100) % 10
}

//Find the fuel cell's rack ID, which is its X coordinate plus 10.
//Begin with a power level of the rack ID times the Y coordinate.
//Increase the power level by the value of the grid serial number (your puzzle input).
//Set the power level to itself multiplied by the rack ID.
//Keep only the hundreds digit of the power level (so 12345 becomes 3; numbers with no hundreds digit become 0).
//Subtract 5 from the power level.
