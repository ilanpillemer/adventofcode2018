package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type pos struct{ x, y int }

func (p pos) down() pos {
	return pos{p.x, p.y + 1}
}

func (p pos) up() pos {
	return pos{p.x, p.y - 1}
}

func (p pos) left() pos {
	return pos{p.x - 1, p.y}
}

func (p pos) right() pos {
	return pos{p.x + 1, p.y}
}

var minx = math.MaxInt64
var maxx = -1
var miny = math.MaxInt64
var maxy = -1
var spring = pos{500, 0}
var clay = map[pos]bool{}
var water = map[pos]bool{}
var settled = map[pos]bool{}

func main() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		parts := strings.Split(line, ",")
		left, right := parts[0], parts[1]
		var xpart, ypart []string
		if strings.Index(left, "x") != -1 {
			xpart = strings.Split(left, "=")
			ypart = strings.Split(right, "=")
			x, _ := strconv.Atoi(xpart[1])
			yrange := strings.Split(ypart[1], "..")
			y1, _ := strconv.Atoi(yrange[0])
			y2, _ := strconv.Atoi(yrange[1])
			vein(x, x, y1, y2)

		} else {
			xpart = strings.Split(right, "=")
			ypart = strings.Split(left, "=")
			y, _ := strconv.Atoi(ypart[1])
			xrange := strings.Split(xpart[1], "..")
			x1, _ := strconv.Atoi(xrange[0])
			x2, _ := strconv.Atoi(xrange[1])
			vein(x1, x2, y, y)
		}
	}
	// begins
	water[pos{500, 1}] = true
	ticks := 0
	for tick() {
		ticks++
		if ticks%10000 == 0 {
			display()
		}
		//fmt.Println("ticks", ticks)
	}

	display()
	fmt.Println("watery", count()) //mustnt include above min y so -2
	fmt.Println("settled", len(settled))

}

func count() int {
	total := len(water)
	less := 0
	for k := range water {
		if k.y < miny {
			less++
		}
	}
	return total - less
}

func tick() bool {
	before := len(water) + len(settled)
	tock()
	after := len(water) + len(settled)
	if before == after {
		return false
	}
	return true
}

func tock() {
	//water flows down
	for k := range water {
		if !blocked(k.down()) {
			if k.down().y > maxy {
				continue
			}
			water[k.down()] = true
			continue
		}
		//water flows left
		if !blocked(k.left()) && blocked(k.left().down()) {
			water[k.left()] = true
		}
		//water flows right
		if !blocked(k.right()) && blocked(k.right().down()) {
			water[k.right()] = true
		}

		if !blocked(k.right()) && blocked(k.down()) {
			water[k.right()] = true
		}

		if !blocked(k.left()) && blocked(k.down()) {
			water[k.left()] = true
		}

	}
	settle()
	//does any water settle?
	//if it has clay on both sides and is blocked below it can settle

}

func blocked(p pos) bool {
	if _, ok := clay[p]; ok {
		return true
	}

	if _, ok := settled[p]; ok {
		return true
	}

	return false
}

func watery(p pos) bool {
	_, ok := water[p]
	return ok
}

func clayey(p pos) bool {
	_, ok := clay[p]
	return ok
}

func shouldSettle(p pos) bool {
	return blocked(p.down())
}

func display() {
	for y := miny; y < maxy+1; y++ {
		for x := minx - 1; x < maxx+1; x++ {
			if _, ok := settled[pos{x, y}]; ok {
				fmt.Print("~")
				continue
			}
			if _, ok := clay[pos{x, y}]; ok {
				fmt.Print("#")
				continue
			}
			if _, ok := water[pos{x, y}]; ok {
				fmt.Print("|")
				continue
			}

			fmt.Print(".")
		}
		fmt.Println()
	}
}

func settle() {
	for y := miny; y < maxy+1; y++ {
		//	inclay := false
		enclosed := false
		start, end := 0, 0
		waterlevel := false
		for x := minx - 1; x < maxx+1; x++ {
			if water[pos{x, y}] {
				waterlevel = true
			}

			if clayey(pos{x, y}) {
				enclosed = true
				//	fmt.Println("enclosed", pos{x, y})
				start, end = x, x
				continue
			}
			// if y == 6 && (x > 495 || x < 501) {
			// 	fmt.Println("water?", pos{x, y}, watery(pos{x, y}))
			// 	fmt.Println("enclosed?", pos{x, y}, enclosed)
			// }
			if watery(pos{x, y}) && enclosed && blocked(pos{x, y}.down()) {
				//	fmt.Println("could be...", start, end)
				end = x
			} else {
				enclosed = false
			}
			if clayey(pos{x, y}.right()) && enclosed {
				block(y, start, end)
				enclosed = false
				continue
			}
		}
		if !waterlevel {
			return
		}
	}
}

func block(y int, start int, end int) {
	for x := start + 1; x < end+1; x++ {
		settled[pos{x, y}] = true
	}
}

func vein(x1, x2, y1, y2 int) {
	if y1 < miny {
		miny = y1
	}

	if y2 > maxy {
		maxy = y2
	}

	if x1 < minx {
		minx = x1
	}

	if x2 > maxx {
		maxx = x2
	}
	for x := x1; x < x2+1; x++ {
		for y := y1; y < y2+1; y++ {
			clay[pos{x, y}] = true
		}
	}
}

// 36789 is wrong?
