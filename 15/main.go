package main

import (
	"bufio"
	"fmt"
	"math"
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
	hp   int
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

func (u *unit) up() pos {
	return pos{x: u.p.x, y: u.p.y - 1}
}

func (u *unit) down() pos {
	return pos{x: u.p.x, y: u.p.y + 1}
}

func (u *unit) left() pos {
	return pos{x: u.p.x - 1, y: u.p.y}
}

func (u *unit) right() pos {
	return pos{x: u.p.x + 1, y: u.p.y}
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

func (u *unit) attackable(target []unit) (*unit, bool) {

	all := make([]unit, 0)
	minhp := 201
	for _, v := range target {
		switch {
		case v.p == u.up():
			all = append(all, v)
			if v.hp < minhp {
				minhp = v.hp
			}
		case v.p == u.down():
			all = append(all, v)
			if v.hp < minhp {
				minhp = v.hp
			}
		case v.p == u.left():
			all = append(all, v)
			if v.hp < minhp {
				minhp = v.hp
			}
		case v.p == u.right():
			all = append(all, v)
			if v.hp < minhp {
				minhp = v.hp
			}
		}
	}

	//no opposition in range
	if len(all) == 0 {
		return nil, false
	}

	//only one opposition in range
	if len(all) == 1 {
		return &all[0], true
	}

	//filter to weakest opponents
	weakest := make([]unit, 0, len(all))
	for _, v := range all {
		if v.hp == minhp {
			weakest = append(weakest, v)
		}
	}
	if len(weakest) == 1 {
		return &weakest[0], true
	}

	//filter by reading order if more than one weakest
	var closest unit
	minx := width
	miny := height
	for _, v := range weakest {
		if v.p.y < miny {
			closest = v
			miny = v.p.y
		}
		if v.p.y == miny {
			if v.p.x < minx {
				closest = v
				minx = v.p.x
			}
		}
	}

	return &closest, true
}

func (u *unit) canMoveTo(target []unit) (*unit, bool) {
	all := make(map[pos]unit, 0)
	reachable := make(map[pos]int, 0)
	closest := make(map[pos]int, 0)
	dist := math.MaxInt64
	//exist positions that can be moved to
	for _, v := range target {
		if _, ok := caverns[v.up()]; ok {
			all[v.up()] = v
		}
		if _, ok := caverns[v.down()]; ok {
			all[v.down()] = v
		}
		if _, ok := caverns[v.left()]; ok {
			all[v.left()] = v
		}
		if _, ok := caverns[v.right()]; ok {
			all[v.right()] = v
		}
	}

	//filter to those positions that are reachable
	for k := range all {
		if moves, ok := distance(u.p, k); ok {
			if moves < dist {
				dist = moves
			}
			reachable[k] = dist
		}
	}
	if len(reachable) == 0 {
		return nil, false
	}

	//filter to closest of the reachable positions
	for k, v := range reachable {
		if v == dist {
			closest[k] = v
		}
	}

	if len(closest) == 1 {
		return nil, true //TODO: fill in value
	}

	return nil, false
}

func distance(pos, pos) (int, bool) {

	return -1, false
}

func (u *unit) attack(unit) {
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

		//Battle Over?
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

		if opp, ok := u.attackable(target); ok {
			fmt.Printf("%v attacks %v\n", u, opp)
			u.attack(*opp)
			continue
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
				elves[pos{i, height}] = unit{elf, pos{i, height}, 200}
			case 'G':
				goblins[pos{i, height}] = unit{goblin, pos{i, height}, 200}
			}
		}
		height++
	}
	width++
}
