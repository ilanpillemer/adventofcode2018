package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y, z, t int
}

var all = []point{}
var e = map[int][]int{} // edges

func main() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		points := strings.Split(line, ",")
		p := point{
			Atop(points[0]),
			Atop(points[1]),
			Atop(points[2]),
			Atop(points[3]),
		}
		all = append(all, p)
	}

	for i, p1 := range all {
		for j, p2 := range all {
			if connected(p1, p2) {
				edges := e[i]
				if edges == nil {
					edges = []int{}
				}
				edges = append(edges, j)
				e[i] = edges
			}
		}
	}

	queue := []int{}
	seen := map[int]bool{}
	count := 0
	for i := range all {
		if _, ok := seen[i]; ok {
			continue
		}
		count++
		queue = append(queue, i)
		for len(queue) > 0 {
			top := queue[0]
			queue = queue[1:]
			if _, ok := seen[top]; ok {
				continue
			}
			seen[top] = true
			edges := e[top]
			for _, j := range edges {
				queue = append(queue, j)
			}

		}

	}
	fmt.Println(count)
}

func connected(p1 point, p2 point) bool {
	dist := abs(p1.x-p2.x) + abs(p1.y-p2.y) + abs(p1.z-p2.z) + abs(p1.t-p2.t)
	if dist <= 3 {
		return true
	}
	return false
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Atop(str string) int {
	str = strings.TrimSpace(str)
	p, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return p
}
