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
	//fmt.Print(pos, "->")
	pos.p.x = pos.p.x + pos.v.x
	pos.p.y = pos.p.y + pos.v.y
	//fmt.Println(pos)
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

func (s *sky) draw() {
	for _, star := range s.stars {

		if star.p.x > -500 && star.p.x < 500 {
			goto cohesion
		}

		if star.p.y > -500 && star.p.y < 500 {
			goto cohesion
		}

		fmt.Println("not near cohesion")
		return
	}
cohesion:
	fmt.Println("cohesion area")

	img := image.NewRGBA(image.Rect(-500, -500, 500, 500))
	for _, star := range s.stars {
		//fmt.Println(star.p.x)
		img.Set(star.p.x, star.p.y, color.White)
		fmt.Println(star.p.x, star.p.y)
	}

	f, err := os.Create(fmt.Sprintf("stars%d.png", s.ticks))
	defer f.Close()

	if err != nil {
		panic(err.Error())
	}

	png.Encode(f, img)

	//img.Close()
	//f.Close()

	//	err = gif.EncodeAll(f, &gif.GIF{
	//		Image: images,
	//		Delay: delays,
	//	})

	//	if err != nil {
	//		panic(err.Error())
	//	}

}

func (s *sky) init(ps []position) {
	s.stars = ps
	//s.sortStars()
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
	for i := 0; i < 50000; i++ {
		s.draw()
		s.tick()
	}
}

func mustAtoi(a string) (i int) {
	i, err := strconv.Atoi(strings.TrimSpace(a))
	if err != nil {
		panic(err.Error())
	}
	return
}