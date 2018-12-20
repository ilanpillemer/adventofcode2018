package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type velocity struct {
	x int
	y int
}

type position struct {
	p point
	v velocity
}

func (pos position) tick() position {
	pos.p.x = pos.p.x + pos.v.x
	pos.p.y = pos.p.y + pos.v.y
	return pos
}

func (pos position) translate(x, y int) position {
	pos.p.x = pos.p.x + x
	pos.p.y = pos.p.y + y
	return pos
}

type sky struct {
	stars []position
	ticks int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (s *sky) withinCohesionRange(limit int) bool {

	s.sortStars()
	if abs(s.stars[0].p.y-s.stars[len(s.stars)-1].p.y) < limit {
		s.sortStars2()
		if abs(s.stars[0].p.x-s.stars[len(s.stars)-1].p.x) < limit {
			return true
		}
	}
	return false
}

var maxx = 0
var maxy = 0
var minx = 100000
var miny = 100000

func (s *sky) draw() {
	if !s.withinCohesionRange(500) {
		return
	}
	// for 100 cohesion
	//maxx 270
	//maxy 138
	//minx 170
	//miny 89

	//for 500 cohesion
	//maxx 470
	//maxy 338
	//minx -30
	//miny -111

	img := image.NewRGBA(image.Rect(-30, -111, 470, 338))
	for _, star := range s.stars {
		img.Set(star.p.x, star.p.y, color.White)
		if star.p.x > maxx {
			maxx = star.p.x
		}
		if star.p.y > maxy {
			maxy = star.p.y
		}
		if star.p.x < minx {
			minx = star.p.x
		}
		if star.p.y < miny {
			miny = star.p.y
		}
	}

	f, err := os.Create(fmt.Sprintf("stars%d.png", s.ticks))
	defer f.Close()

	if err != nil {
		panic(err.Error())
	}

	png.Encode(f, img)
}

func (s *sky) init(ps []position) {
	s.stars = ps
}

func (s *sky) tick() {
	for i, star := range s.stars {
		s.stars[i] = star.tick()
	}
	s.ticks++
}

func (s *sky) translate(x, y int) {
	for i, star := range s.stars {
		s.stars[i] = star.translate(x, y)
	}
	s.ticks++
}

func (s *sky) sortStars() {
	sort.Slice(s.stars, func(i, j int) bool {
		if s.stars[i].p.y == s.stars[j].p.y {
			return s.stars[i].p.x < s.stars[j].p.x
		}
		return s.stars[i].p.y < s.stars[j].p.y
	})
}

func (s *sky) sortStars2() {
	sort.Slice(s.stars, func(i, j int) bool {
		if s.stars[i].p.x == s.stars[j].p.x {
			return s.stars[i].p.y < s.stars[j].p.y
		}
		return s.stars[i].p.x < s.stars[j].p.x
	})
}

func main() {
	s := sky{}

	r := bufio.NewReader(os.Stdin)
	ps := make([]position, 0)
	for {
		line, err := r.ReadString('\n')
		if line == "" && err == io.EOF {
			break
		}
		re := regexp.MustCompile("^position=<(.+)> velocity=<(.+)>")
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			coords := strings.Split(matches[1], ",")
			vector := strings.Split(matches[2], ",")
			pos := position{
				p: point{mustAtoi(coords[0]), mustAtoi(coords[1])},
				v: velocity{mustAtoi(vector[0]), mustAtoi(vector[1])},
			}
			ps = append(ps, pos)
		}
	}
	fmt.Println("starting to draw")
	s.init(ps)
	//s.translate(50000, 50000)
	//g := &GIF{}

	for i := 0; i < 50000; i++ {
		s.draw()
		s.tick()
	}
	fmt.Println("maxx", maxx)
	fmt.Println("maxy", maxy)
	fmt.Println("minx", minx)
	fmt.Println("miny", miny)
	fmt.Println("finished draw")
}

func mustAtoi(a string) (i int) {
	i, err := strconv.Atoi(strings.TrimSpace(a))
	if err != nil {
		panic(err.Error())
	}
	return
}