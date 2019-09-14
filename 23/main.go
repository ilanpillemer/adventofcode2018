package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
)

type pos struct{ x, y, z, r int }

func (p pos) up() pos      { return pos{p.x, p.y - 1, p.z, p.r} }
func (p pos) down() pos    { return pos{p.x, p.y + 1, p.z, p.r} }
func (p pos) left() pos    { return pos{p.x - 1, p.y, p.z, p.r} }
func (p pos) right() pos   { return pos{p.x + 1, p.y, p.z, p.r} }
func (p pos) forward() pos { return pos{p.x, p.y, p.z + 1, p.r} }
func (p pos) back() pos    { return pos{p.x, p.y, p.z - 1, p.r} }

var nano = map[pos]bool{}
var maxr = 0
var maxnano pos

var maxx, maxy, maxz, minx, miny, minz int
var l sync.Mutex
var l2 sync.Mutex

type Item struct {
	distance int
	count    int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		//fmt.Println(line)
		p0 := strings.Split(line, ",")
		xs := strings.TrimPrefix(p0[0], "pos=<")
		ys := p0[1]
		zs := strings.TrimSuffix(p0[2], ">")
		rs := strings.TrimPrefix(p0[3], " r=")
		x, _ := strconv.Atoi(xs)
		y, _ := strconv.Atoi(ys)
		z, _ := strconv.Atoi(zs)
		r, _ := strconv.Atoi(rs)

		nano[pos{x, y, z, r}] = true
		if r > maxr {
			maxr = r
			maxnano = pos{x, y, z, r}
		}

		if x > maxx {
			maxx = x
		}
		if y > maxy {
			maxy = y
		}
		if z > maxz {
			maxz = z
		}

		if x < minx {
			minx = x
		}
		if y < miny {
			miny = y
		}
		if z < minz {
			minz = z
		}
	}
	count := 0
	for k := range nano {
		if maxnano.inrange(k) {
			count++
		}
	}

	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()
	fmt.Println("Number in range of max nano is", count)
	fmt.Println("max", maxx, maxy, maxz)
	fmt.Println("min", minx, miny, minz)

	// used the idea from the subreddit https://www.reddit.com/r/adventofcode/comments/a8s17l/2018_day_23_solutions/
	// also looked at https://github.com/BarDweller/aoc2018/blob/master/day23/part2.go
	queue := PriorityQueue{}
	heap.Init(&queue)

	for k := range nano {
		distance := distFromOrigin(k)
		maxd := distance + k.r
		mind := distance - k.r
		heap.Push(&queue, &Item{distance: mind, count: 1})
		heap.Push(&queue, &Item{distance: maxd, count: -1})
	}

	best := 0
	enc := 0
	max := 0
	for len(queue) > 0 {
		item := (heap.Pop(&queue)).(*Item)
		enc += item.count
		if enc > max {
			max = enc
			best = item.distance
		}
	}

	fmt.Println("most encircled ", best)

}

var cache = map[pos]int{}
var seen = map[pos]bool{}
var cand = map[pos]int{}

var maxjumps = 10000000
var jumps = 0

func (n pos) findMaximum() {
	//fmt.Println("trying ", n)

	current := n.totN()
	up := n.up().totN()
	down := n.down().totN()
	left := n.left().totN()
	right := n.right().totN()
	forward := n.forward().totN()
	back := n.back().totN()

	switch {
	case current < up:
		n.up().findMaximum()
	case current < down:
		n.down().findMaximum()
	case current < left:
		n.left().findMaximum()
	case current < right:
		n.right().findMaximum()
	case current < forward:
		n.forward().findMaximum()
	case current < back:
		n.back().findMaximum()
	default:
		addToCandidates(n, n.totN())
	}

	//after 100 jumps found the following candidates
	//	fmt.Println("ok.. after jump,", maxjumps)

	//	os.Exit(0)

}

func getRandomPos() pos {
	x := rand.Intn(maxx+abs(minx)) - abs(minx)
	y := rand.Intn(maxy+abs(miny)) - abs(miny)
	z := rand.Intn(maxz+abs(minz)) - abs(minz)
	//fmt.Println("randomly jumping to:", x, y, z)
	return pos{x, y, z, 0}
}

func addToCandidates(n pos, c int) {
	newCandList := make(map[pos]int)
	biggest := true
	for k, v := range cand {
		if v > c {
			return // existing candidates are better
		}
		if v == c {
			fmt.Printf("found %v -> %d : %d\n", n, c, distFromOrigin(n))
			biggest = false
			newCandList[k] = v
		}
	}
	if biggest {
		fmt.Printf("\nnew max %v -> %d : %d\n", n, c, distFromOrigin(n))
	}

	newCandList[n] = c
	//	l.Lock()
	cand = newCandList
	//	l.Unlock()
}

func (n pos) totN() int {
	l2.Lock()
	defer l2.Unlock()
	if i, ok := cache[n]; ok {
		return i
	}
	count := 0
	for k := range nano {
		if n.withinNanoBotRange(k) {
			//fmt.Printf("%v is in range of %v\n", n, k)
			count++
		} else {
			//fmt.Printf("%v is in not in range of %v\n", n, k)
		}
	}

	cache[n] = count
	return count
}

func (n pos) withinNanoBotRange(t pos) bool {
	if (abs(t.x-n.x) + abs(t.y-n.y) + abs(t.z-n.z)) <= t.r {
		return true
	}
	return false
}

func (n pos) inrange(t pos) bool {
	if (abs(t.x-n.x) + abs(t.y-n.y) + abs(t.z-n.z)) <= n.r {
		return true
	}
	return false
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func distFromOrigin(t pos) int {
	return abs(t.x) + abs(t.y) + abs(t.z)
}
