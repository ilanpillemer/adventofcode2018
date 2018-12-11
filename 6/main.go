package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

type grid struct {
	locs                            []pos          // possible teleport destinations
	projected                       map[string]int // the grid populated by its nearest locations
	top, bottom, left, right, width int
}

//init sets the values that determine the
//frame of the grid in the infinite void.
//it also leaves the positions sorted.
func (g *grid) init() {
	sort.Slice(g.locs, func(i, j int) bool {
		if g.locs[i].x < g.locs[j].x {
			return true
		}
		if g.locs[i].x > g.locs[j].x {
			return false
		}
		if g.locs[i].y < g.locs[j].y {
			return true
		}
		return false
	})
	g.top = g.locs[0].y
	g.left = g.locs[0].x
	g.right = g.locs[len(g.locs)-1].x
	g.bottom = g.locs[len(g.locs)-1].y
	g.width = g.right - g.left + 1
	g.projected = make(map[string]int)

}

//project assumes that init as been called.
//it works out to which location each position in the grid is closest
func (g *grid) project() {
	for y := 0; y <= g.bottom+1000; y++ {
		for x := 0; x <= g.right; x++ {
			key := fmt.Sprintf("%d,%d", x, y)
			c, _ := closest(pos{x, y}, g.locs)
			g.projected[key] = c
			//log.Printf("%d,%d --> %d", x, y, c)
		}
	}
	//log.Println(g.projected)
}

//prints out the nice little projected grid
func (g *grid) print() {
	fmt.Println()
	for i := 0; i <= g.bottom+500; i++ {
		for j := 0; j <= g.right; j++ {
			z := g.getOwner(j, i)
			if z == -1 {
				fmt.Print(".")
			} else {
				fmt.Printf("%c", (z%64)+65)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *grid) getLargestNonInfiniteArea() {
	counts := make(map[string]int)

	//count the areas
	for _, v := range g.projected {
		key := fmt.Sprintf("%c", (v%64)+65)
		counts[key] = counts[key] + 1
	}

	//remove infinites
	fmt.Println(counts)
	for i := 0; i <= g.bottom+500; i++ {
		for j := 0; j <= g.right; j++ {
			if i == 0 || j == 0 || i == g.bottom+500 || j == g.right {
				z := g.getOwner(j, i)
				key := fmt.Sprintf("%c", (z%64)+65)
				delete(counts, key)
			}
		}

	}
	fmt.Println(counts)
	max := 0
	key := ""
	for k, v := range counts {
		if v > max {
			key = k
			max = v
		}
	}
	fmt.Println("max", counts[key])

}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (g *grid) getOwner(x int, y int) int {
	key := fmt.Sprintf("%d,%d", x, y)
	return g.projected[key]
}

func manh(a pos, b pos) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

//returns index and pos of location that is closest, or -1 if more than one is closest
func closest(a pos, locs []pos) (int, pos) {
	min := 1<<63 - 1
	index := -1
	p := pos{}

	counts := make(map[int]int)
	for i, l := range locs {
		mh := manh(a, l)
		counts[mh] = counts[mh] + 1
		if mh < min {
			min = mh
			index = i
			p = l
		}
	}
	if counts[min] > 1 {
		return -1, p
	}

	return index, p
}

func main() {
	r := bufio.NewReader(os.Stdin)
	locations := make([]pos, 0)
	for {
		line, err := r.ReadString('\n')

		if strings.TrimSpace(line) == "" && err == io.EOF {
			g := &grid{locs: locations}
			g.init()
			g.project()
			//g.print()
			g.getLargestNonInfiniteArea()
			log.Println("Bye..")
			os.Exit(0)
		}
		fields := strings.Split(line, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(fields[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(fields[1]))
		locations = append(locations, pos{x, y})

	}
}