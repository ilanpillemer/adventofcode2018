package main

import (
	"bufio"
	"io"
	"regexp"
	"strings"
	"testing"
)

func TestDraw(t *testing.T) {
	s := sky{}

	testdata := `position=< 9,  1> velocity=< 0,  2>
position=< 7,  0> velocity=<-1,  0>
position=< 3, -2> velocity=<-1,  1>
position=< 6, 10> velocity=<-2, -1>
position=< 2, -4> velocity=< 2,  2>
position=<-6, 10> velocity=< 2, -2>
position=< 1,  8> velocity=< 1, -1>
position=< 1,  7> velocity=< 1,  0>
position=<-3, 11> velocity=< 1, -2>
position=< 7,  6> velocity=<-1, -1>
position=<-2,  3> velocity=< 1,  0>
position=<-4,  3> velocity=< 2,  0>
position=<10, -3> velocity=<-1,  1>
position=< 5, 11> velocity=< 1, -2>
position=< 4,  7> velocity=< 0, -1>
position=< 8, -2> velocity=< 0,  1>
position=<15,  0> velocity=<-2,  0>
position=< 1,  6> velocity=< 1,  0>
position=< 8,  9> velocity=< 0, -1>
position=< 3,  3> velocity=<-1,  1>
position=< 0,  5> velocity=< 0, -1>
position=<-2,  2> velocity=< 2,  0>
position=< 5, -2> velocity=< 1,  2>
position=< 1,  4> velocity=< 2,  1>
position=<-2,  7> velocity=< 2, -2>
position=< 3,  6> velocity=<-1, -1>
position=< 5,  0> velocity=< 1,  0>
position=<-6,  0> velocity=< 2,  0>
position=< 5,  9> velocity=< 1, -2>
position=<14,  7> velocity=<-2,  0>
position=<-3,  6> velocity=< 2, -1>
`

	r := bufio.NewReader(strings.NewReader(testdata))
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

	s.init(ps)
	s.draw()

}
