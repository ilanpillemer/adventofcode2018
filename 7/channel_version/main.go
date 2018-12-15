package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"sync"
)

type edge struct {
	from string
	to   string
}

var nodes map[string]node

type node struct {
	id   string
	outs []edge
	ins  []edge
}

// closure stuff

var UniversalTickerTime = make(chan struct{})

type availableQ struct {
	sync.Mutex
	steps string
	todo  string
	ready string
}

//var todo = "_ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var order = ""

var available = &availableQ{}

func (a *availableQ) pop() (rune, int) {
	a.Lock()
	defer a.Unlock()

	if len(a.steps) == 0 {
		return 0, 0
	}

	a.prioritise()
	fmt.Println("prioritised", a.steps)
	p := a.steps[0]
	a.steps = strings.Replace(a.steps, string(p), "", 1)
	//fmt.Println(string(p), p-64)
	if string(p) == "_" {
		return rune(p), 1
	}
	return rune(p), int(p - 64)
}

func (a *availableQ) push(id string) {
	a.Lock()
	defer a.Unlock()
	if strings.Contains(a.todo, id) && !strings.Contains(a.steps, id) && complete(id) {
		a.ready = a.ready + id
	}
	a.prioritise()
}

func (a *availableQ) size() int {
	return len(a.steps)
}

func (a *availableQ) clear() {
	a.Lock()
	defer a.Unlock()
	a.steps = ""
}

func (a *availableQ) prioritise() {
	//pun intended
	//	a.Lock()
	//	defer a.Unlock()

	runed := []rune(a.steps)
	sort.Slice(runed, func(i, j int) bool {
		return runed[i] < runed[j]
	})
	a.steps = string(runed)
}

func (a *availableQ) update() {
	a.Lock()
	defer a.Unlock()
	a.steps = a.steps + a.ready
	a.ready = ""
	a.prioritise()

}

func (a *availableQ) setTodo(todo string) {
	a.todo = todo
}

func (a *availableQ) GetTodo() string {
	return a.todo
}

func (a *availableQ) LeftTodo() int {
	return len(a.todo)
}

var assembled = make(chan struct{})

func initSleigh(todo string) {
	//todo = "_ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	order = ""
	available.clear()
	available.setTodo(todo)
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

func do(id string, step string) bool {

	if strings.Contains(available.GetTodo(), step) {
		fmt.Println(id, "doing", step, "STEPS done", order, available.GetTodo(), "available:", available.steps, available.size())
		order = order + step
		available.Lock()
		available.setTodo(strings.Replace(available.GetTodo(), step, "", 1))
		available.Unlock()
		//fmt.Println(id, " says he has done STEP", step)
		o := nodes[step]
		for _, out := range o.outs {
			fmt.Printf("adding %s to available.\n", out.to)
			available.push(out.to)
		}
		return true
	}
	fmt.Printf("maybe this is the bug? worker %s step %s\n", id, step)
	return false
}

func assemble(n node) {
	available.push(n.id)
	available.update()
	<-assembled
}

func main() {

	//go worker("rudolph")

	r := bufio.NewReader(os.Stdin)
	all := make([]string, 0)
	for {
		line, err := r.ReadString('\n')
		if strings.TrimSpace(line) == "" && err == io.EOF {
			edges := process(all)
			nodes := createNodes(edges)
			root := addRoot(starts(nodes))
			assemble(root)
		}
		all = append(all, strings.TrimSpace(line))
	}

}

func complete(id string) bool {
	ins := nodes[id]
	for _, in := range ins.ins {
		if strings.Contains(available.GetTodo(), in.from) {
			return false
		}

	}
	return true
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

func starts(nodes map[string]node) []node {
	s := make([]node, 0)
	for k, v := range nodes {
		if len(v.ins) == 0 {
			s = append(s, nodes[k])
		}
	}
	return s
}
