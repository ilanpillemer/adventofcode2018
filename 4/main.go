package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
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

func (g *guard) sleepiestMinute() (int, int) {
	max := 0
	var key int
	for k, v := range g.sleepyMinutes {
		if v > max {
			max = v
			key = k
		}
	}
	return key, max
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
				g.id = currentGuardId
				g.update(scrawl)
				guards[currentGuardId] = g
			}

			// part 1
			max := 0
			var lazy guard

			for k, v := range guards {
				if v.totalMinutes >= max {
					max = v.totalMinutes
					lazy = guards[k]
				}
			}
			fmt.Printf("lazy guard[%v]\n", lazy)

			max = 0
			bingo := 0
			for k, v := range lazy.sleepyMinutes {
				if v > max {
					max = v
					bingo = k
				}
			}
			fmt.Printf("sleepy minute[%d]\n", bingo)

			// part 2
			max = 0
			var lethargic guard

			for k, v := range guards {
				_, amount := v.sleepiestMinute()
				if amount >= max {
					max = amount
					lethargic = guards[k]
				}
			}
			minute, amount := lethargic.sleepiestMinute()
			fmt.Printf("lethargic guard[%v] slept at minute [%d], [%d] times\n", lethargic, minute, amount)

			// part 1 output
			x, _ := strconv.Atoi(strings.TrimPrefix(lazy.id, "#"))
			fmt.Printf("%d X %d = %d\n", x, bingo, x*bingo)

			// part 2 output
			x2, _ := strconv.Atoi(strings.TrimPrefix(lethargic.id, "#"))
			fmt.Printf("%d X %d = %d\n", x2, minute, x2*minute)
			os.Exit(0)
		}
		scrawls = append(scrawls, NewRaw(input))
	}

}