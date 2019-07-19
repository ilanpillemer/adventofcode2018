package main

import (
	"bufio"
	"fmt"
	"os"
)

type state int

const (
	finished state = iota
	explored
)

type pos struct{ x, y int }

func (p pos) n() pos { return pos{p.x, p.y - 1} }
func (p pos) s() pos { return pos{p.x, p.y + 1} }
func (p pos) w() pos { return pos{p.x - 1, p.y} }
func (p pos) e() pos { return pos{p.x + 1, p.y} }

//graph of maze. the key is a node and it has a list of nodes as its edges
//these are stored bidirectional
var maze = map[pos][]pos{}

func initMaze() {
	maze = map[pos][]pos{}
	dists = make([]int, 0)
}

func walk(path string, s pos) (pos, string, state) {
	entry := s
	if !exists(s) {
		maze[s] = []pos{}
	}
	for {
		switch car(path) {
		case '^':
			walk(cdr(path), s)
			path = cdr(path)
		case '$':
			return s, path, finished
		case 'N':
			grow(s, s.n())
			grow(s.n(), s)
			s = s.n()
			path = cdr(path)
		case 'S':
			grow(s, s.s())
			grow(s.s(), s)
			path = cdr(path)
			s = s.s()
		case 'W':
			grow(s, s.w())
			grow(s.w(), s)
			path = cdr(path)
			s = s.w()
		case 'E':
			grow(s, s.e())
			grow(s.e(), s)
			path = cdr(path)
			s = s.e()
		case '(':
			next, remain, _ := walk(cdr(path), s)
			s = next
			path = remain
		case ')':
			return s, cdr(path), explored
		case '|':
			if car(cdr(path)) == ')' {
				//in all cases where there is a |)
				//every option brings the walker
				//back to where they started anyway
				//so even if I returned s instead of entry
				//this would work!
				//I checked this on the puzzle input by inspection as well
				return entry, cdr(cdr(path)), explored
			}
			s = entry
			path = cdr(path)
		}

	}
}

func car(i string) byte {
	return i[0]
}

func cdr(i string) string {
	return i[1:]
}

func exists(p pos) bool {
	_, ok := maze[p]
	return ok
}

//add bidirectional edges in the maze graph
func grow(s pos, d pos) {
	opts := maze[s]
	for _, o := range opts {
		if o == d {
			return
		}
	}
	opts = append(opts, d)
	maze[s] = opts
	return
}

var seen = map[pos]bool{}
var dists []int

//walk every path until you reach a dead or circle back onto yourself
func explore(s pos, seen map[pos]bool, dist int) {
	seen[s] = true
	directions := maze[s]
	for _, front := range directions {
		if _, ok := seen[front]; !ok { //make sure you have not been here
			seen[front] = true           //ok I can go somewhere
			explore(front, seen, dist+1) //lets go there noting it took longer
		}
	}
	//you cant get any further so note down how long it took
	dists = append(dists, dist)
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	input := in.Text()
	initMaze()
	walk(input, pos{})
	explore(pos{}, map[pos]bool{}, 0)
	fmt.Println("longest", longest())
	fmt.Println("many", count())
}

func longest() int {
	max := 0
	for _, i := range dists {
		if max < i {
			max = i
		}
	}
	return max
}

func count() int {
	num := 0
	for _, i := range dists {
		if i >= 1000 {
			num++
		}
	}
	return num
}
