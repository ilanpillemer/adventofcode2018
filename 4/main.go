package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

// a little db

type record struct {
	minutes [60]int // could be a bit or a boolean for better space
}

const layout = "[2006-01-02 15:04]"

type raw struct {
	date string // month-day eg 11-01
	rest string // rest of scrawl on wall
}

func NewRaw(input string) raw {
	//[1518-11-01 00:00] Guard #10 begins shift"
	r := raw{}
	r.date = input[0:18]
	r.rest = strings.TrimSpace(input[19:])
	return r
}

type byDate []raw

func (d byDate) Len() int      { return len(d) }
func (d byDate) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
func (d byDate) Less(i, j int) bool {
	left, _ := time.Parse(layout, d[i].date)
	right, _ := time.Parse(layout, d[j].date)
	return right.After(left)
}

type guard struct {
	id            string
	data          map[string]record
	sleepyMinutes map[int]int
	totalMinutes  int
	asleep        bool
	fellAsleep    int
}

func (g *guard) update(input raw) {
	date, _ := time.Parse(layout, input.date)
	key := fmt.Sprintf("%d-%d", date.Month(), date.Day())
	record := g.data[key]
	minute := date.Minute()
	fallingAsleep := strings.Contains(input.rest, "falls asleep")
	wakingUp := strings.Contains(input.rest, "wakes up")
	if fallingAsleep {
		g.asleep = true
		g.fellAsleep = minute
	}

	if wakingUp {
		g.asleep = false
		for i := g.fellAsleep; i < date.Minute(); i++ {
			if g.sleepyMinutes == nil {
				g.sleepyMinutes = map[int]int{}
			}
			g.sleepyMinutes[i] = g.sleepyMinutes[i] + 1
			record.minutes[i] = record.minutes[i] + 1
			g.totalMinutes = g.totalMinutes + 1
		}
		g.fellAsleep = minute
	}

}

func main() {
	scrawls := make([]raw, 0)
	guards := make(map[string]guard)

	r := bufio.NewReader(os.Stdin)
	for {
		input, err := r.ReadString('\n')
		if input == "" && err == io.EOF {
			sort.Sort(byDate(scrawls))

			currentGuardId := "#UNKNOWN"
			for _, scrawl := range scrawls {
				//fmt.Println(scrawl)
				if strings.Contains(scrawl.rest, "#") {
					currentGuardId = strings.Fields(scrawl.rest)[1]
					continue
				}
				g := guards[currentGuardId]
				g.update(scrawl)
				guards[currentGuardId] = g
			}

			for k, v := range guards {
				fmt.Printf("guard[%s], totalAsleep[%d]\n", k, v.totalMinutes)
			}
			os.Exit(0)
		}
		scrawls = append(scrawls, NewRaw(input))
	}

}