package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type device struct {
	twos   int
	threes int
}

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

func main() {
	r := bufio.NewReader(os.Stdin)
	d := &device{}
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