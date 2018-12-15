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
	steps  string
	todo   string
	ready  string
	offset int
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
	p := a.steps[0]
	a.steps = strings.Replace(a.steps, string(p), "", 1)

	if string(p) == "_" {
		return rune(p), 1
	}
	return rune(p), int(p-64) + a.offset
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

func initSleigh(todo string, offset int) {
	//todo = "_ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	order = ""
	available.clear()
	available.setTodo(todo)
	available.offset = offset

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
	<-ticker
	//check closure
	state := cmap[id]
	// if closure has something that can be finished this time slice
	if state.cost == 1 {
		//complete this step
		fmt.Printf("%s, %s || ", id, state.step)
		do(id, state.step)
		//update closure
		cmap[id] = cval{}
	}

	// if closure has something that be worked on but is not finishable
	if state.cost > 1 {
		fmt.Printf("%s, %s || ", id, state.step)
		cmap[id] = cval{state.step, state.cost - 1}
	}

	// if was idle previously and ready to take on work, check if there is anything available
	if state.cost == 0 {
		// if there is something available
		if available.size() != 0 {
			r, cost := available.pop()
			cmap[id] = cval{string(r), cost}
		}

		//repeat same checks as above
		state := cmap[id]
		if state.cost == 1 {
			//complete this step
			fmt.Printf("%s, %s || ", id, state.step)
			do("id", state.step)
			//update closure
			cmap[id] = cval{}
		}
		if state.cost > 1 {
			fmt.Printf("%s, %s || ", id, state.step)
			cmap[id] = cval{state.step, state.cost - 1}
		}
	}

	// if was idle previously and there is nothing available
	if state.cost == 0 {
		if available.size() == 0 {
			fmt.Printf("%s, %s || ", id, state.step)
		}
	}
	//indicate that santa is done
	done <- struct{}{}
}

func do(id string, step string) bool {

	if strings.Contains(available.GetTodo(), step) {
		order = order + step
		available.Lock()
		available.setTodo(strings.Replace(available.GetTodo(), step, "", 1))
		available.Unlock()
		o := nodes[step]
		for _, out := range o.outs {
			available.push(out.to)
		}
		return true
	}
	return false
}

func assemble(n node) {
	available.push(n.id)
	available.update()
	<-assembled
}

func oneWorker(all []string, sleigh string) {
	edges := process(all)
	nodes := createNodes(edges)
	root := addRoot(starts(nodes))
	initSleigh(sleigh, 0)

	go func() {
		worker1 := "ozzy"
		for {
			// make sure available is updated for next tick
			available.update()
			worker1C := make(chan struct{})
			go worker(worker1, worker1C)
			fmt.Println("time is about to tick")
			ticker <- struct{}{}

			fmt.Println("time ticked globally")
			fmt.Println("time is waiting for santa")
			<-worker1C // wait for santa
			fmt.Printf("%s announced that he had a moment\n", worker1)
			if len(available.GetTodo()) == 0 {
				fmt.Println("assembled with steps", order)
				assembled <- struct{}{}
				break
			}
		}
	}()
	assemble(root)
}

func twoWorker(all []string, sleigh string) {
	edges := process(all)
	nodes := createNodes(edges)
	root := addRoot(starts(nodes))
	initSleigh(sleigh, 0)
	go func() {
		worker1 := "ozzy"
		worker2 := "cesar"
		for {
			// make sure available is updated for next tick
			available.update()
			worker1C := make(chan struct{})
			worker2C := make(chan struct{})
			go worker(worker1, worker1C)
			go worker(worker2, worker2C)
			fmt.Println("time is about to tick")
			ticker <- struct{}{}
			ticker <- struct{}{}

			fmt.Println("time is waiting for worker 1 and worker 2")
			<-worker1C // wait for worker1
			fmt.Printf("%s announced that he had a moment\n", worker1)
			<-worker2C // wait for worker2
			fmt.Printf("%s announced that he had a moment\n", worker2)
			if len(available.GetTodo()) == 0 {
				fmt.Println("assembled with steps", order)
				assembled <- struct{}{}
				break
			}
		}
	}()
	assemble(root)
}

func fiveWorker(all []string, sleigh string) {
	edges := process(all)
	nodes := createNodes(edges)
	root := addRoot(starts(nodes))
	initSleigh(sleigh, 60)
	count := 0
	go func() {
		worker1 := "ilan "
		worker2 := "cesar"
		worker3 := "ozzy "
		worker4 := "flo  "
		worker5 := "erin i"
		for {
			// make sure available is updated for next tick
			count++
			available.update()
			worker1C := make(chan struct{})
			worker2C := make(chan struct{})
			worker3C := make(chan struct{})
			worker4C := make(chan struct{})
			worker5C := make(chan struct{})
			go worker(worker1, worker1C)
			go worker(worker2, worker2C)
			go worker(worker3, worker3C)
			go worker(worker4, worker4C)
			go worker(worker5, worker5C)
			fmt.Println()
			ticker <- struct{}{}
			ticker <- struct{}{}
			ticker <- struct{}{}
			ticker <- struct{}{}
			ticker <- struct{}{}

			<-worker1C // wait for worker1
			<-worker2C // wait for worker2
			<-worker3C // wait for worker1
			<-worker4C // wait for worker2
			<-worker5C // wait for worker2

			if len(available.GetTodo()) == 0 {
				assembled <- struct{}{}
				break
			}
		}
	}()
	assemble(root)
	fmt.Println("ticks", count-1)
}

func main() {

	r := bufio.NewReader(os.Stdin)
	all := make([]string, 0)
	for {
		line, err := r.ReadString('\n')
		if strings.TrimSpace(line) == "" && err == io.EOF {
			fiveWorker(all, "_ABCDEFGHIJKLMNOPQRSTUVWXYZ")
			fmt.Println("ordered", order)
			os.Exit(0)
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
