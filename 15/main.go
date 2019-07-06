package main

import (
	"bufio"
	"fmt"
	"os"
)

type unitType int
type pos struct {
	x, y int
}

const (
	elf unitType = iota
	goblin
)

var (
	roundCount = 0
	walls      = make(map[pos]bool)
	caverns    = make(map[pos]bool)
	elves      = make(map[pos]unit)
	goblins    = make(map[pos]unit)
	width      = 0
	height     = 0
)

type unit struct {
	race unitType
}

func display() {
	for y := 0; y < height; y++ {
		for x := 0; x < width+1; x++ {
			if _, ok := walls[pos{x, y}]; ok {
				fmt.Print("#")
				continue
			}
			if _, ok := caverns[pos{x, y}]; ok {
				fmt.Print(".")
			}
			if _, ok := elves[pos{x, y}]; ok {
				fmt.Print("E")
			}
			if _, ok := goblins[pos{x, y}]; ok {
				fmt.Print("G")
			}
		}
		fmt.Println()
	}
}

func (u *unit) targets() []unit {
	return nil
}

func (u *unit) canAttack([]unit) bool {
	return false
}

func (u *unit) canMoveTo([]unit) bool {
	return false
}

func (u *unit) attack([]unit) {
}

func (u *unit) move([]unit) {
}

func main() {
	fmt.Println("Scan....")
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		for i, c := range line {
			if i > width {
				width = i
			}
			switch c {
			case '#':
				walls[pos{i, height}] = true
			case '.':
				caverns[pos{i, height}] = true
			case 'E':
				elves[pos{i, height}] = unit{elf}
			case 'G':
				goblins[pos{i, height}] = unit{goblin}
			}
		}
		height++
	}
	display()
	fmt.Println("Fight!!!!")
	os.Exit(1)
	// rounds
rounds:
	for {
		for _, u := range startPositions() {
			targets := u.targets()
			if len(targets) == 0 {
				break rounds
			}

			inRange := determine(targets)
			if u.canAttack(inRange) {
				u.attack(inRange)
				continue
			}
			if u.canMoveTo(inRange) {
				u.move(inRange)
				continue
			}
		}
	}
	fmt.Println("Combat Over!!!!")
}

func startPositions() []unit {
	return nil
}

func determine(targets []unit) []unit {
	return nil
}
