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
			//fmt.Println(i, i-center)
			sum += (i - center)
		}
	}
	return sum
}

var center = 0

func spread(curr string) string {
	// initialise
	next := ""
	// iterate from beginning index 2 (position 3) to index len-3 (position 3 from the end) to  to end of string
	// chops off two on left each time
	curr = "..." + curr + "..."
	center += 1
	for i := 2; i < len(curr)-2; i++ {
		//fmt.Printf("%s\n", curr[i-2:i+3])
		if _, ok := rules[curr[i-2:i+3]]; !ok {
			next = next + "."
			//	fmt.Println("added .")
		} else {
			next = next + rules[curr[i-2:i+3]]
			//	fmt.Println("added", rules[curr[i-2:i+3]])
		}
		//next = next + curr[i-2:i+3]
	}

	//fmt.Println()
	//fmt.Println(next)

	// apply rules until one rules applies
	// append a string to next if rule applies
	// if no rule applies throw an error
	// if strings.HasPrefix(next, "...") {
	// 	next = strings.TrimSuffix(next, "...")
	// 	center -= 3
	// }
	//	next = strings.TrimSuffix(next, "...")
	return strings.TrimSuffix(next, "...")
	//	return strings.TrimSuffix((strings.TrimPrefix(next, "...")), "...")
}
