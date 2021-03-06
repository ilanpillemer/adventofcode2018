package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
)

type unitType int
type pos struct {
	x, y int
}

var elfpower = flag.Int("elfpower", 3, "elfpower")

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
		return fmt.Sprintf("elf%v [%d]", u.p, u.hp)
	case goblin:
		return fmt.Sprintf("goblin%v [%d]", u.p, u.hp)
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

func (p *pos) up() pos {
	return pos{x: p.x, y: p.y - 1}
}

func (p *pos) down() pos {
	return pos{x: p.x, y: p.y + 1}
}

func (p *pos) left() pos {
	return pos{x: p.x - 1, y: p.y}
}

func (p *pos) right() pos {
	return pos{x: p.x + 1, y: p.y}
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

	for y := 0; y < height; y++ {
		for x := 0; x < height; x++ {
			if e, ok := elves[pos{x, y}]; ok {
				fmt.Printf("E %d %d,%d\n", e.hp, y, x)
			}
			if g, ok := goblins[pos{x, y}]; ok {
				fmt.Printf("G %d %d,%d\n", g.hp, y, x)
			}
		}
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

func (u *unit) attack(victim unit) {
	switch u.race {
	case goblin:
		victim.hp = victim.hp - 3
	case elf:
		victim.hp = victim.hp - *elfpower
	}
	if victim.hp <= 0 {
		switch victim.race {
		case elf:
			delete(elves, victim.p)
			caverns[victim.p] = true
		case goblin:
			delete(goblins, victim.p)
			caverns[victim.p] = true
		}
		//fmt.Printf("Arrggghhhh.. %v is dead\n", victim)
		//display()
		return
	}

	switch victim.race {
	case elf:
		elves[victim.p] = victim
	case goblin:
		goblins[victim.p] = victim
	}

}

func distance(dest pos, src pos) (int, bool) {
	//fmt.Println("distance", dest, src)
	//breadth first search
	queue := make([]pos, 0)
	seen := make(map[pos]int)
	queue = append(queue, src)
	seen[src] = -1 // actual positions for src and dest not counted in example in AoC

	for len(queue) != 0 {
		// pop
		//	fmt.Println(queue, seen)
		q := queue[0]
		queue = queue[1:]

		if q == dest {
			//	fmt.Println("found!!", seen[q])
			return seen[q], true
		}

		// push
		if _, ok := caverns[q.up()]; ok || dest == q.up() {
			if _, ok := seen[q.up()]; !ok {
				seen[q.up()] = seen[q] + 1
				queue = append(queue, q.up())
			}
		}

		if _, ok := caverns[q.down()]; ok || dest == q.down() {
			if _, ok := seen[q.down()]; !ok {
				seen[q.down()] = seen[q] + 1
				queue = append(queue, q.down())
			}
		}

		if _, ok := caverns[q.left()]; ok || dest == q.left() {
			if _, ok := seen[q.left()]; !ok {
				seen[q.left()] = seen[q] + 1
				queue = append(queue, q.left())
			}
		}

		if _, ok := caverns[q.right()]; ok || dest == q.right() {
			if _, ok := seen[q.right()]; !ok {
				seen[q.right()] = seen[q] + 1
				queue = append(queue, q.right())
			}
		}
	}
	return -1, false
}

type node struct {
	n        pos
	children []node
}

func path(dest pos, src pos) (map[pos]node, bool) {
	//fmt.Println("distance", dest, src)
	//breadth first search
	queue := make([]pos, 0)
	seen := make(map[pos]int)
	queue = append(queue, src)
	seen[src] = -1 // actual positions for src and dest not counted in example in AoC
	tree := make(map[pos]node)
	tree[src] = node{src, []node{}}

	for len(queue) != 0 {
		// pop
		//	fmt.Println(queue, seen)
		q := queue[0]
		queue = queue[1:]
		if q == dest {
			if (src == pos{20, 4} && dest == pos{19, 8}) {
				fmt.Println("bfs tree:", tree)
			}

			//	fmt.Println("found!!", seen[q])
			return tree, true
		}
		tree[q] = node{q, []node{}}

		// push in priority reading order
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				switch {
				case q.up() == pos{x, y}:
					if _, ok := caverns[q.up()]; ok || dest == q.up() {

						if _, ok := seen[q.up()]; !ok || dest == q.up() {
							seen[q.up()] = seen[q] + 1
							t := tree[q]
							t.children = append(t.children, node{q.up(), []node{}})
							tree[q] = t
							queue = append(queue, q.up())
						}
					}
				case q.down() == pos{x, y}:
					if _, ok := caverns[q.down()]; ok || dest == q.down() {
						if _, ok := seen[q.down()]; !ok || dest == q.down() {
							seen[q.down()] = seen[q] + 1
							t := tree[q]
							t.children = append(t.children, node{q.down(), []node{}})
							tree[q] = t

							queue = append(queue, q.down())
						}
					}

				case q.left() == pos{x, y}:
					if _, ok := caverns[q.left()]; ok || dest == q.left() {
						if _, ok := seen[q.left()]; !ok || dest == q.left() {
							seen[q.left()] = seen[q] + 1
							t := tree[q]
							t.children = append(t.children, node{q.left(), []node{}})
							tree[q] = t

							queue = append(queue, q.left())
						}
					}

				case q.right() == pos{x, y}:
					if _, ok := caverns[q.right()]; ok || dest == q.right() {
						if _, ok := seen[q.right()]; !ok || dest == q.right() {
							seen[q.right()] = seen[q] + 1
							t := tree[q]
							t.children = append(t.children, node{q.right(), []node{}})
							tree[q] = t

							queue = append(queue, q.right())
						}
					}

				}
			}
		}

	}
	return tree, false
}

func inrange(target []unit) ([]pos, bool) {
	in := make([]pos, 0, len(target))
	for _, u := range target {
		if _, ok := caverns[u.p.up()]; ok {
			in = append(in, u.p.up())
		}
		if _, ok := caverns[u.p.down()]; ok {
			in = append(in, u.p.down())
		}
		if _, ok := caverns[u.p.left()]; ok {
			in = append(in, u.p.left())
		}
		if _, ok := caverns[u.p.right()]; ok {
			in = append(in, u.p.right())
		}
	}

	if len(in) == 0 {
		return in, false
	}
	return in, true
}

func nearest(u unit, target []unit) (pos, bool) {
	nearest := make(map[pos]int, 0)
	mindist := math.MaxInt64
	if in, ok := inrange(target); ok {
		//fmt.Printf("[%v] inrange: %v\n", u, in)
		for _, i := range in {
			if dist, ok := distance(i, u.p); ok {
				//fmt.Println(i, "<->", u.p, dist)
				nearest[i] = dist
				if dist < mindist {
					mindist = dist
				}
			}
		}
	}

	for k, v := range nearest {
		if v != mindist {
			delete(nearest, k)
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if _, ok := nearest[pos{x, y}]; ok {
				return pos{x, y}, true
			}
		}
	}

	return pos{}, false
}

func main() {
	flag.Parse()
	fmt.Println("Scan....")
	scan(bufio.NewScanner(os.Stdin))
	fmt.Println("Fight!!!!")

	totelves := len(elves)
	// initiative := initiatives()

	count := 0
	// start round
	display()
	for {
		count++
		//fmt.Println("Initiative Determined...")
		invalid := map[pos]bool{} // indicates that a unit in this position in the iniative is no longer valid
		initiative := initiatives()
		//fmt.Println("initiative:", initiative)

		//fmt.Println("Round", count)
		for _, u := range initiative {
			//Battle Over?
			// there has been an attack.. check you still exist!
			//BUG still here.. can steal an extra turn
			//I no longer exist if there is a cavern here
			//but what if someone now is here?

			if _, ok := caverns[u.p]; ok {
				//I no longer exist
				continue
			}

			if invalid[u.p] {
				//this turn is not mine
				fmt.Println("bug fixed..")
				continue
			}

			target := u.targets()
			if len(target) < 1 {
				fmt.Println("Game Over")
				display()
				if u.race == elf {
					fmt.Println("Elves Win")
				} else {
					fmt.Println("Goblins Win")
				}
				fmt.Println("Rounds fully Completed", count-1)
				sum := 0
				for _, v := range elves {
					sum += v.hp
					fmt.Println(v)
				}
				for _, v := range goblins {
					sum += v.hp
					fmt.Println(v)
				}
				fmt.Println("Total HP remaining", sum)
				fmt.Println("Outcome ", sum*(count-1))
				fmt.Printf("Elves alive %d out of %d\n", len(elves), totelves)
				if totelves == len(elves) {
					fmt.Println("Timeline restored.. all elves alive.")
				}
				os.Exit(0)
			}

			if opp, ok := u.attackable(target); ok {
				//fmt.Printf("%v attacks %v\n", u, opp)
				//check opp still exists
				if _, ok := caverns[opp.p]; !ok {
					//fmt.Println("potential bug if this unit has already actually attacked")
					//fmt.Printf("%v attacks %v\n", u, opp)
					u.attack(*opp)
					continue
				}

			}

			if near, ok := nearest(u, target); ok {
				//fmt.Printf("%v: nearest -> %v \n", u, near)
				tree, _ := path(near, u.p)
				//fmt.Printf("tree: %v\n", tree)
				move := readingOrderNextMove(near, u.p, tree)
				invalid[move] = true
				//fmt.Printf("%v moves to %v\n", u, move)
				switch u.race {
				case elf:
					e := elves[u.p]
					e.p = move
					elves[move] = e
					delete(elves, u.p)
					caverns[u.p] = true
					delete(caverns, move)
					//after moving.. it can still attack if possible
					//target := e.targets()
					if opp, ok := e.attackable(target); ok {
						//	fmt.Printf("%v attacks after a move %v\n", e, opp)
						e.attack(*opp)
						continue
					}

				case goblin:
					g := goblins[u.p]
					g.p = move
					goblins[move] = g
					delete(goblins, u.p)
					caverns[u.p] = true
					delete(caverns, move)
					//after moving.. it can still attack if possible
					//target := g.targets()
					if opp, ok := g.attackable(target); ok {
						//	fmt.Printf("%v attacks after a move %v\n", g, opp)
						g.attack(*opp)
						continue
					}

				}
				//display()
			}

		}
		//display()
		//	fmt.Println(count)

	}
	fmt.Println("Combat Over!!!!")
}

func readingOrderNextMove(dest pos, src pos, tree map[pos]node) pos {

	//remove all paths that dont lead to destination
	changes := true
	for changes {
		changes = false
		for k, v := range tree {
			if len(v.children) == 0 {
				delete(tree, k)
				changes = true
				continue
			}
			children := v.children
			for i := len(v.children) - 1; i > -1; i-- {
				if _, ok := tree[children[i].n]; !ok && children[i].n != dest {
					children = append(children[0:i], children[i+1:]...)
					changes = true
				}
			}
			t := tree[k]
			t.children = children
			tree[k] = t
		}
	}

	//return the reading order element of the first of possible shortest paths (if more than one)
	//fmt.Printf("cleaned:: %v\n", tree)
	path := tree[src]
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			for _, v := range path.children {
				if (pos{x, y} == v.n) {
					return v.n
				}
			}
		}
	}
	panic("unreachable")

}

func initiatives() []unit {
	initiative := make([]unit, 0)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if u, ok := elves[pos{x, y}]; ok {
				initiative = append(initiative, u)
				//				fmt.Println(x, y)
			}
			if u, ok := goblins[pos{x, y}]; ok {
				initiative = append(initiative, u)
				//				fmt.Println(x, y)
			}
		}
	}
	//	fmt.Println("inside initiative", initiative)
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
