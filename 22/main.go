package main

import (
	"flag"
	"fmt"
)

type region int

const (
	rocky region = iota
	wet
	narrow
)

var depth = flag.Int("depth", 510, "depth")
var ty = flag.Int("y", 10, "target y")
var tx = flag.Int("x", 10, "target x")
var display = flag.Bool("display", true, "display")

var cache = map[pos]region{}
var gicache = map[pos]int{}

func main() {
	flag.Parse()
	fmt.Println("target", *tx, *ty)
	fmt.Println("depth", *depth)
	if *display {
		render()
	}
	fmt.Println("risk level", risklevel())

}

type pos struct{ x, y int }

func (p pos) String() string {
	switch p.rtype() {
	case rocky:
		return "."
	case wet:
		return "="
	case narrow:
		return "|"
	default:
		return "?"
	}
}

func (p pos) gindex() int {
	if gi, ok := gicache[p]; ok {
		return gi
	}

	if p.x == 0 && p.y == 0 {
		return 0
	}

	if p.x == *tx && p.y == *ty {
		return 0
	}

	if p.y == 0 {
		return p.x * 16807
	}

	if p.x == 0 {
		return p.y * 48271
	}

	gicache[p] = pos{x: p.x - 1, y: p.y}.elevel() * pos{x: p.x, y: p.y - 1}.elevel()

	return gicache[p]

}

func (p pos) elevel() int {
	return (p.gindex() + *depth) % 20183
}

func (p pos) rtype() region {
	if rt, ok := cache[p]; ok {
		return rt
	}
	cache[p] = region(p.elevel() % 3)
	return cache[p]
}

func render() {
	for y := 0; y < *ty+1; y++ {
		for x := 0; x < *tx+1; x++ {
			fmt.Printf("%s", pos{x, y})
		}
		fmt.Println()
	}
}

func risklevel() int {
	rl := 0
	for y := 0; y < *ty+1; y++ {
		for x := 0; x < *tx+1; x++ {
			switch (pos{x, y}).rtype() {
			case rocky:
			case wet:
				rl++
			case narrow:
				rl += 2
			}
		}
	}
	return rl
}
