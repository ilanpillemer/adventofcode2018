package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var rules = make(map[string]string)
var gens = flag.Int("gen", 0, "num gens")

func main() {
	flag.Parse()
	fmt.Println("Lets do some gardening!!")
	in := bufio.NewScanner(os.Stdin)
	// line 1
	in.Scan()
	curr := in.Text()[15:] //load initial state
	in.Scan()              // skip a line

	for in.Scan() != false { //load rules
		rule := strings.Fields(in.Text())
		rules[rule[0]] = rule[2]
		fmt.Printf("Loaded Rule [%s] => [%s] \n", rule[0], rules[rule[0]])
	}

	fmt.Println(curr)
	prevsum := 0

	if *gens <= 101 {
		for i := 0; i < *gens; i++ {
			curr = spread(curr)
			fmt.Printf("gen %d -> %d (diff is %d)\n", i+1, sum(curr), sum(curr)-prevsum)
			prevsum = sum(curr)
		}
		fmt.Println("sum", sum(curr))
	}

	// gen 100 is 4184 and then 38 gets added every generation
	// if sum is greater than 100 then sum = 4184 + (n-100 * 38)
	if *gens > 100 {
		fmt.Println("using precalc info.. sum", 4184+(*gens-100)*38)
	}
}

func sum(curr string) int {
	sum := 0
	for i, s := range curr {
		if s == '#' {
			sum += (i - center)
		}
	}
	return sum
}

var center = 0

func spread(curr string) string {
	next := ""
	curr = "..." + curr + "..."
	center++
	for i := 2; i < len(curr)-2; i++ {
		if _, ok := rules[curr[i-2:i+3]]; !ok {
			next = next + "."
		} else {
			next = next + rules[curr[i-2:i+3]]
		}
	}
	return strings.TrimSuffix(next, "...")
}
