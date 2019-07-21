package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pos struct{ x, y, z, r int }

var nano = map[pos]bool{}
var maxr = 0
var maxnano pos

func main() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		fmt.Println(line)
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
	}
	count := 0
	for k := range nano {
		if maxnano.inrange(k) {
			count++
		}
	}

	fmt.Println("Number in range of max nano is", count)

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
