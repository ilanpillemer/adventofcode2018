package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type device struct {
	twos   int
	threes int
	boxes  []string
}

var isSearching = flag.Bool("f", false, "turn on fabric search mode")

func (d *device) scan(input string) {
	counts := make(map[rune]int)
	for _, c := range input {
		counts[c] = counts[c] + 1
	}

	hasTwo, hasThree := false, false
	for _, v := range counts {
		if hasTwo && hasThree {
			break
		}
		switch v {
		case 2:
			{
				hasTwo = true
			}
		case 3:
			{
				hasThree = true
			}
		}
	}
	if hasTwo {
		d.twos = d.twos + 1
	}
	if hasThree {
		d.threes = d.threes + 1
	}
}

func (d *device) checksum() int {
	return d.twos * d.threes
}

// part 2 device functionality

func (d *device) check(input string) (string, bool) {
	for _, fabric := range d.boxes {
		if matches(input, fabric) {
			return fabric, true
		}
	}
	d.boxes = append(d.boxes, input)
	return "", false
}

func matches(i1 string, i2 string) bool {
	for _, s := range i1 {
		i2 = strings.Replace(i2, string(s), "", 1)
	}
	if len(i2) <= 1 {
		return true
	}
	return false
}

func main() {
	r := bufio.NewReader(os.Stdin)
	d := &device{}

	flag.Parse()
	if *isSearching {
		fmt.Println("searching mode is ON")
		for {
			input, err := r.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "" && err == io.EOF {
				fmt.Println("GoodBye...")
				os.Exit(0)
			}
			result, ok := d.check(input)
			if ok {
				fmt.Printf("Two Boxes are [%s] and [%s]\n", input, result)
				os.Exit(0)
			}
		}
	}

	// default checksum mode (part 1)

	for {
		input, err := r.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" && err == io.EOF {
			fmt.Printf("checksum: [%d]\n", d.checksum())
			os.Exit(0)
		}
		d.scan(input)
	}
}