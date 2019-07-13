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

var minx = math.MaxInt64
var maxx = -1
var miny = math.MaxInt64
var maxy = -1

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

	display()
}

var clay = map[pos]bool{}

func display() {
	for y := miny; y < maxy+1; y++ {
		for x := minx; x < maxx+1; x++ {
			if _, ok := clay[pos{x, y}]; ok {
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
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
