package main

import (
	"container/heap"
	"flag"
	"fmt"
	"os"
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
	fmt.Println("finding route")
	traverse()

}

type tool int

const (
	torch tool = iota
	gear
	neither
)

type pos struct{ x, y int }

func (p pos) up() pos    { return pos{p.x, p.y - 1} }
func (p pos) down() pos  { return pos{p.x, p.y + 1} }
func (p pos) left() pos  { return pos{p.x - 1, p.y} }
func (p pos) right() pos { return pos{p.x + 1, p.y} }

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

func (p pos) valid() bool {
	if p.x < 0 || p.y < 0 {
		return false
	}

	return true

}

func (p pos) allowed() []tool {
	switch p.rtype() {
	case rocky:
		return ([]tool{gear, torch})
	case wet:
		return ([]tool{gear, neither})
	case narrow:
		return ([]tool{torch, neither})
	default:
		panic("allowed: illegal tool")
	}
}

func (p pos) isAlllowed(t tool) bool {
	switch p.rtype() {
	case rocky:
		if t == gear || t == torch {
			return true
		}
	case wet:
		if t == gear || t == neither {
			return true
		}
	case narrow:
		if t == torch || t == neither {
			return true
		}
	}
	return false
}

type move struct {
	p pos
	e tool
}

func (m move) options(seen map[move]bool) []move {
	moves := []move{}
	directions := [](func() pos){
		m.p.up, m.p.down, m.p.left, m.p.right,
	}

	for _, f := range directions {
		if f().valid() {
			for _, t := range m.p.allowed() {
				if f().isAlllowed(t) {
					option := move{f(), t}
					if _, ok := seen[option]; !ok {
						moves = append(moves, option)
					}
				}
			}
		}
	}
	return moves
}

func (m move) distance() int {
	//assumes x and y is always postive
	if m.p.x < 0 || m.p.y < 0 {
		panic("distance: negative")
	}
	return m.p.x + m.p.y
}

func traverse() {
	frontier := make(pqueue, 0)
	seen := map[move]bool{}
	first := move{p: pos{0, 0}, e: torch}
	i := item{value: first}
	frontier = append(frontier, &i)

	for len(frontier) > 0 {
		i := heap.Pop(&frontier).(*item)
		cand := i.value
		if _, ok := seen[cand]; ok { //strictly longer route
			continue
		}
		seen[cand] = true

		// are we there?
		if (cand.p == pos{*tx, *ty}) {
			if cand.e != torch {
				//perhaps but if you add 7 suddenly now... may be wrong
				tot := i.time + 7
				fmt.Println("perhaps: ", tot)
				os.Exit(0)
			} else {
				fmt.Println("time: ", i.time)
				os.Exit(0)
			}
		}
		// where can we go now? at what time cost?

		for _, m := range cand.options(seen) {
			newtime := i.time + 1
			if m.e != cand.e {
				newtime += 7
			}
			newpriority := newtime + m.distance()
			o := item{
				value:    m,
				priority: newpriority,
				time:     newtime,
			}
			heap.Push(&frontier, &o)
		}
	}

}

type pqueue []*item
type item struct {
	value    move
	priority int
	index    int
	time     int //not part of uniqeness of a move
}

func (pq pqueue) Len() int           { return len(pq) }
func (pq pqueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq pqueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *pqueue) Push(x interface{}) {
	n := len(*pq)
	i := x.(*item)
	i.index = n
	*pq = append(*pq, i)
}

func (pq *pqueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
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
