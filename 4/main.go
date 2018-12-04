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
	date    time.Time // month-day eg 11-01
	id      string    // eg #10
	minutes [60]int   // could be a bit or a boolean for better space
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

func main() {
	scrawls := make([]raw, 0)

	r := bufio.NewReader(os.Stdin)
	for {
		input, err := r.ReadString('\n')
		if input == "" && err == io.EOF {
			sort.Sort(byDate(scrawls))

			for _, scrawl := range scrawls {
				fmt.Println(scrawl)
			}
			os.Exit(0)
		}
		scrawls = append(scrawls, NewRaw(input))
	}

}