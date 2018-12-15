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
	
	
	ticker = make(chan struct{})
cmap = make(map[string]cval)

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

var ticker = make(chan struct{})
var cmap = make(map[string]cval)

type cval struct {
	step string
	cost int
}

func worker(id string, done chan<- struct{}) {
	//wait for time to tick
	fmt.Printf("%s is waiting for time\n", id)
	<-ticker
	fmt.Printf("time ticked for %s\n", id)
	//check closure
	state := cmap[id]
	// if closure has something that can be finished this time slice
	if state.cost == 1 {
		//complete this step
		fmt.Printf("%s is completing stepp %s\n", id, state.step)
		do(id, state.step)
		//update closure
		cmap[id] = cval{}
	}

	// if closure has something that be worked on but is not finishable
	if state.cost > 1 {
		fmt.Printf("%s is working on step %s\n", id, state.step)
		cmap[id] = cval{state.step, state.cost - 1}
		fmt.Printf("closure map now looks like this %v\n", cmap)
	}

	// if was idle previously and ready to take on work, check if there is anything available
	if state.cost == 0 {
		// if there is something available
		fmt.Printf("is there is anything available for %s? %+v\n", id, available)
		if available.size() != 0 {
			r, cost := available.pop()
			fmt.Printf("Yes, %s is available with cost %d.\n", string(r), cost)
			cmap[id] = cval{string(r), cost}
			fmt.Printf("closure map now looks like this %v\n", cmap)
		}

		//repeat same checks as above
		state := cmap[id]
		if state.cost == 1 {
			//complete this step
			fmt.Printf("%s is completing step %s\n", id, state.step)
			do("id", state.step)
			//update closure
			cmap[id] = cval{}
		}
		if state.cost > 1 {
			fmt.Printf("%s works at %s\n", id, state.step)
			cmap[id] = cval{state.step, state.cost - 1}
			fmt.Printf("closure map now looks like this %v\n", cmap)
		}
	}

	// if was idle previously and there is nothing available
	if state.cost == 0 {
		if available.size() == 0 {
			// stay idle
		}
	}
	//indicate that santa is done
	fmt.Printf("moment is over for %s\n", id)
	done <- struct{}{}
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
