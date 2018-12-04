package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type claim struct {
	id     string
	left   int
	top    int
	width  int
	height int
}

func NewClaim(input string) claim {
	//#1 @ 1,3: 4x4
	fields := strings.Fields(input)
	c := claim{}
	c.id = fields[0]

	pos := strings.TrimSuffix(fields[2], ":")
	coord := strings.Split(pos, ",")
	c.left, _ = strconv.Atoi(coord[0])
	c.top, _ = strconv.Atoi(coord[1])
	quants := strings.Split(fields[3], "x")
	c.width, _ = strconv.Atoi(quants[0])
	c.height, _ = strconv.Atoi(quants[1])
	return c
}

type fabricCell struct {
	x int
	y int
}

func (f *fabricCell) String() string {
	return fmt.Sprintf("%d,%d", f.x, f.y)
}

// overlap is defined as not too left, too right, too low, too high
func overlap(c1 *claim, c2 *claim) bool {
	// too left
	if c1.left+c1.width <= c2.left {
		return false
	}

	// too right
	if c1.left >= c2.left+c2.width {
		return false
	}

	// too high
	if c1.top+c1.height <= c2.top {
		return false
	}

	// too low
	if c1.top >= c2.top+c2.height {
		return false
	}

	// too true
	return true
}

func overlapSize(c1 *claim, c2 *claim) int {
	if !overlap(c1, c2) {
		return 0
	}
	return h(c1, c2) * w(c1, c2)
}

func overlapCells(c1 *claim, c2 *claim) []fabricCell {
	cells := make([]fabricCell, 0)
	if c1.top < c2.top {
		c1, c2 = c2, c1
	}
	top := c1.top
	bottom := c1.top + h(c1, c2)

	if c1.left < c2.left {
		c1, c2 = c2, c1
	}
	left := c1.left
	right := c1.left + w(c1, c2)

	for i := top; i < bottom; i++ {
		for j := left; j < right; j++ {
			cells = append(cells, fabricCell{i, j})
		}
	}
	return cells
}

var perfect = flag.Bool("f", false, "find that perfect little claim")

func main() {
	flag.Parse()
	r := bufio.NewReader(os.Stdin)
	claims := make([]claim, 0)
	for {
		input, err := r.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" && err == io.EOF {
			break
		}
		claims = append(claims, NewClaim(input))
	}

	crosshatch := make(map[string]int)
	//cross join land of the friend of friends
	if *perfect {
		FAIL := false
		for _, c1 := range claims {
			for _, c2 := range claims {
				if c1.id != c2.id {
					if overlap(&c1, &c2) {
						FAIL = true
						break
					}
				}
			}
			if !FAIL {
				fmt.Printf("Perfect Claim is .... [%s]\n", c1.id)
			}
			FAIL = false
		}
		os.Exit(0)
	}
	// default behaviour of finding wasted inches
	for _, c1 := range claims {
		for _, c2 := range claims {
			if c1.id != c2.id {
				cells := overlapCells(&c1, &c2)
				// and another inner loop ... this is getting icky..
				// though this is an inc(h)y (w)inchy loop
				for _, cell := range cells {
					crosshatch[cell.String()] = crosshatch[cell.String()] + 1
				}
			}
		}
	}
	// count those little inchy squares
	sum := 0
	for _, v := range crosshatch {
		if v >= 2 {
			sum = sum + 1
		}
	}
	fmt.Printf("wasted square inches from claims: [%d]\n", sum)
}

func w(c1 *claim, c2 *claim) int {
	if c1.left < c2.left {
		c1, c2 = c2, c1
	}
	return min(c2.width-abs(c1.left-c2.left), c1.width)
}

func h(c1 *claim, c2 *claim) int {
	if c1.top < c2.top {
		c1, c2 = c2, c1
	}
	return min(c2.height-abs(c1.top-c2.top), c1.height)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}