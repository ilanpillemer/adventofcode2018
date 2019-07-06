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
	p    pos
}

func (u unit) String() string {
	switch u.race {
	case elf:
		return fmt.Sprintf("elf%v", u.p)
	case goblin:
		return fmt.Sprintf("goblin%v", u.p)
	}
	return "?"
}

func display() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
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
	targets := make([]unit, 0)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch u.race {
			case elf:
				if u, ok := goblins[pos{x, y}]; ok {
					targets = append(targets, u)
				}
			case goblin:
				if u, ok := elves[pos{x, y}]; ok {
					targets = append(targets, u)
				}
			}
		}
	}
	return targets
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
	scan(bufio.NewScanner(os.Stdin))
	display()
	fmt.Println("Fight!!!!")
	fmt.Println("Initiative Determined...")
	initiative := initiatives()
	fmt.Println(initiative)
	// start round
	for _, u := range initiative {
		target := u.targets()
		if len(target) < 1 {
			fmt.Println("Game Over")
			if u.race == elf {
				fmt.Println("Elves Win")
			} else {
				fmt.Println("Goblins Win")
			}
			os.Exit(0)
		}
	}
	fmt.Println("Combat Over!!!!")
}

func initiatives() []unit {
	initiative := make([]unit, 0)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if u, ok := elves[pos{x, y}]; ok {
				initiative = append(initiative, u)
				fmt.Println(x, y)
			}
			if u, ok := goblins[pos{x, y}]; ok {
				initiative = append(initiative, u)
				fmt.Println(x, y)
			}
		}
	}
	return initiative
}

func determine(targets []unit) []unit {
	return nil
}

func scan(in *bufio.Scanner) {
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
				elves[pos{i, height}] = unit{elf, pos{i, height}}
			case 'G':
				goblins[pos{i, height}] = unit{goblin, pos{i, height}}
			}
		}
		height++
	}
	width++
}
