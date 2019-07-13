package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct{ x, y int }

var height int
var width int

var state map[pos]rune

func main() {
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
	display()
	fmt.Println()
	fmt.Println("mutating")

	for i := 0; i < 10; i++ {
		state = mutate(state)
	}
	fmt.Println()
	display()

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
	fmt.Println("score", tree*ly)
}

func mutate(current map[pos]rune) map[pos]rune {

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

	return state
}
