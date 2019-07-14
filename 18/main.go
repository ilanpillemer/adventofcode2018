package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type pos struct{ x, y int }

var height int
var width int
var gens = flag.Int("gens", 10, "generations")
var state map[pos]rune
var cache = make(map[int]int)

func main() {
	flag.Parse()
	state = map[pos]rune{}
	in := bufio.NewScanner(os.Stdin)
	y := 0
	for in.Scan() {
		line := in.Text()
		for x, c := range line {
			state[pos{x, y}] = c
		}
		if width < len(line) {
			width = len(line)
		}
		y++
	}
	height = y
	// loaded initial state
	fmt.Println()
	fmt.Println("mutating")

	//TODO work out how to do this without manual inspection
	// cycle is 28
	// cycle start at 580
	// 580 should be 165376
	// 609 should be 163494
	// 741 should be 186686
	skip := *gens

	if *gens > 579 {
		jump := (*gens - 580) % 28
		fmt.Println("can skip in part 2")
		skip = 580 + jump
	}
	fmt.Println("reduced skip to ", skip)

	for i := 0; i < skip; i++ {
		state = mutate(state)
	}

	fmt.Println()
	display()
	fmt.Println(score(state))
}

var counter = 0

func cycle(state map[pos]rune) (map[pos]rune, int, bool) {
	before := score(state)
	state = mutate(state)
	if _, ok := cache[before]; ok {
		return state, counter, true
	}
	cache[before] = counter
	counter++
	return state, 0, false
}

func score(state map[pos]rune) int {
	ly := 0
	tree := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if state[pos{x, y}] == '#' {
				ly++
			}
			if state[pos{x, y}] == '|' {
				tree++
			}
		}
	}
	return tree * ly
}

func display() {
	ly := 0
	tree := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Printf("%c", state[pos{x, y}])
			if state[pos{x, y}] == '#' {
				ly++
			}
			if state[pos{x, y}] == '|' {
				tree++
			}
		}
		fmt.Println()
	}
}

func mutate(current map[pos]rune) map[pos]rune {

	//check cache
	var adjacent = func(e pos) [8]pos {
		var adj [8]pos
		adj[0] = pos{e.x, e.y - 1}
		adj[1] = pos{e.x, e.y + 1}
		adj[2] = pos{e.x + 1, e.y}
		adj[3] = pos{e.x - 1, e.y}
		adj[4] = pos{e.x - 1, e.y - 1}
		adj[5] = pos{e.x - 1, e.y + 1}
		adj[6] = pos{e.x + 1, e.y - 1}
		adj[7] = pos{e.x + 1, e.y + 1}
		return adj
	}

	// An open acre will become filled with trees if three or more
	// adjacent acres contained trees. Otherwise, nothing happens.

	var change1 = func(in pos) (rune, bool) {
		if current[in] != '.' {
			return current[in], false
		}
		adj := adjacent(in)
		tree := 0
		for _, v := range adj {
			if current[v] == '|' {
				tree++
			}
			if tree >= 3 {
				return '|', true
			}
		}
		return current[in], true
	}

	// An acre filled with trees will become a lumberyard if three or more
	// adjacent acres were lumberyards. Otherwise, nothing happens.

	var change2 = func(in pos) (rune, bool) {
		if current[in] != '|' {
			return current[in], false
		}
		adj := adjacent(in)
		lumberyards := 0
		for _, v := range adj {
			if current[v] == '#' {
				lumberyards++
			}
			if lumberyards >= 3 {
				return '#', true
			}
		}
		return current[in], true
	}

	// An acre containing a lumberyard will remain a lumberyard if it was
	// adjacent to at least one other lumberyard and at least one acre
	// containing trees. Otherwise, it becomes open.

	var change3 = func(in pos) (rune, bool) {
		if current[in] != '#' {
			return current[in], false
		}
		adj := adjacent(in)
		lumberyards := 0
		trees := 0
		for _, v := range adj {
			if current[v] == '|' {
				trees++
			}
			if current[v] == '#' {
				lumberyards++
			}

			if lumberyards >= 1 && trees >= 1 {
				return '#', true
			}
		}
		return '.', true
	}

	state := make(map[pos]rune)
	for k := range current {
		var ok bool
		if state[k], ok = change1(k); ok {
			continue
		}
		if state[k], ok = change2(k); ok {
			continue
		}
		if state[k], ok = change3(k); ok {
			continue
		}
	}
	//	cache[score(current)] = score(state)
	return state
}
