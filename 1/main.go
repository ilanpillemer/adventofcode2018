package main

import (
	"bufio"
	"flag"
	"fmt"
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

func main() {
	f := freq{0}
	r := bufio.NewReader(os.Stdin)

	flag.Parse()
	switch flag.Arg(0) {

// this solves part 1
	default:
		for {
			input, _ := r.ReadString('\n')
			input = strings.TrimSpace(input)
			if strings.TrimSpace(input) == "" {
				fmt.Println(f.value)
				os.Exit(0)
			}
			f.apply(nice(input))
		}
	}
}

func nice(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return i
}