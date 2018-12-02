package main

import (
	"bufio"
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

func nice(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return i
}