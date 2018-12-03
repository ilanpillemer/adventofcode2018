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

type freq struct {
	value int
}

func (f *freq) apply(input int) {
	f.value = f.value + input
}

var calibrating = flag.Bool("c", false, "calibrating mode")

func main() {
	f := freq{0}
	r := bufio.NewReader(os.Stdin)

	flag.Parse()
	if *calibrating {
		fmt.Println("calibrating")
		a := turnIntoArray(r)
		seen := make(map[int]int)
		seen[0] = seen[0] + 1
		i := 0
		for {
			count := seen[f.value]
			if count == 2 {
				fmt.Println(f.value)
				os.Exit(0)
			}
			f.apply(a[i])
			seen[f.value] = seen[f.value] + 1
			i = (i + 1) % len(a)
		}
	}

	// not calibrating, so part 1
	for {
		input, err := r.ReadString('\n')
		input = strings.TrimSpace(input)
		if strings.TrimSpace(input) == "" && err == io.EOF {
			fmt.Println(f.value)
			os.Exit(0)
		}
		f.apply(nice(input))
	}

}

func turnIntoArray(r *bufio.Reader) []int {
	a := make([]int, 0)
	for {
		input, err := r.ReadString('\n')
		input = strings.TrimSpace(input)
		if strings.TrimSpace(input) == "" && err == io.EOF {
			return a
		}
		if strings.TrimSpace(input) != "" {
			a = append(a, nice(input))
		}
	}
}

func nice(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	if i == 0 {
		fmt.Println("wtf")
	}
	return i
}