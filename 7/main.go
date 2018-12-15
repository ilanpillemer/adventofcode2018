package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type edge struct {
	from string
	to   string
}

type node struct {
	id   string
	outs []edge
	ins  []edge
}

var nodes map[string]node

var todo = "_ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var times = makeTimes()
var available = "_"
var order = ""
var timing = 0

func makeTimes() map[rune]int {
	times := make(map[rune]int)
	for i, r := range todo {
		if i == 0 {
			continue
		}
		times[r] = 60 + i
	}
	return times
}

func add(id string) {
	if strings.Contains(todo, id) && !strings.Contains(available, id) && complete(id) {
		available += id
	}
}

func rejig() {
	//pun intended
	runed := []rune(available)
	sort.Slice(runed, func(i, j int) bool {
		return runed[i] < runed[j]
	})
	available = string(runed)
}

func do() bool {

	if len(todo) == 0 {
		fmt.Println("all done")
		return true
	}

	if len(available) == 0 {
		return true
	}
	rejig()
	next := string(available[0])
	if strings.Contains(todo, next) {
		order = order + next
		todo = strings.Replace(todo, next, "", 1)
		available = strings.Replace(available, next, "", 1)
		o := nodes[next]
		for _, out := range o.outs {
			add(out.to)
		}
		do()
	}
	return false
}

func addRoot(starts []node) node {
	root := node{id: "_"}

	for _, n := range starts {
		e := edge{"_", n.id}
		root.outs = append(root.outs, e)
		s := nodes[n.id]
		s.ins = append(s.ins, e)
		nodes[n.id] = s
	}
	nodes["_"] = root
	return root
}

func complete(id string) bool {
	ins := nodes[id]
	for _, in := range ins.ins {
		if strings.Contains(todo, in.from) {
			return false
		}

	}
	return true
}

func (e edge) String() string {
	return e.from + "->" + e.to
}

func starts(nodes map[string]node) []node {
	s := make([]node, 0)
	for k, v := range nodes {
		if len(v.ins) == 0 {
			s = append(s, nodes[k])
		}
	}
	return s
}

func ends(nodes map[string]node) []node {
	e := make([]node, 0)
	for k, v := range nodes {
		if len(v.outs) == 0 {
			e = append(e, nodes[k])
		}
	}
	return e
}

func process(input []string) []edge {
	edges := make([]edge, 0)
	for _, line := range input {
		fields := strings.Fields(line)
		e := edge{fields[1], fields[7]}
		edges = append(edges, e)
	}
	return edges
}

func createNodes(edges []edge) map[string]node {
	nodes = make(map[string]node)
	for _, e := range edges {
		in, ok := nodes[e.from]
		if !ok {
			in = node{id: e.from}
		}
		in.outs = append(in.outs, e)
		nodes[e.from] = in

		out, ok := nodes[e.to]
		if !ok {
			out = node{id: e.to}
		}
		out.ins = append(out.ins, e)
		nodes[e.to] = out
	}

	return nodes
}

func walk(n node) {
	add(n.id)
	do()
	fmt.Println()
}

func walkConcurrent(n node) {
	add(n.id)
	do()
	fmt.Println()
}

func main() {
	r := bufio.NewReader(os.Stdin)
	all := make([]string, 0)
	for {
		line, err := r.ReadString('\n')
		if strings.TrimSpace(line) == "" && err == io.EOF {
			edges := process(all)
			nodes := createNodes(edges)
			root := addRoot(starts(nodes))
			walk(root)
			fmt.Println(order)
			goto part2
		}
		all = append(all, strings.TrimSpace(line))
	}
part2:
fmt.Println("welcome to part2")
}
