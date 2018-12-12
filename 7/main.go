package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type node struct {
	id   string
	next []*node
}

var root = &node{}

func (n *node) lookup(id string) (*node, bool) {
	if n.id == "" {
		n.id = id
		return n, true
	}

	if n.id == id {
		return n, true
	}

	for _, next := range n.next {
		found, ok := next.lookup(id)
		if ok {
			return found, true
		}
	}
	return nil, false
}

func (n *node) display(i int) {
	fmt.Print(strings.Repeat(".", i))
	fmt.Println(n.id)
	for _, next := range n.next {
		next.display(i + 4)
	}
}

var ordered string

func (n *node) init() {
	root = &node{}
	ordered = ""
}

func (n *node) order() {
	ordered = strings.Replace(ordered, n.id, "", -1)
	ordered += n.id
	sort.Slice(n.next, func(i, j int) bool {
		return n.next[i].id < n.next[j].id
	})
	for _, next := range n.next {
		next.order()
	}
}




// https://en.wikipedia.org/wiki/Topological_sorting
//L ← Empty list that will contain the sorted elements
//S ← Set of all nodes with no incoming edge
//while S is non-empty do
//    remove a node n from S
//    add n to tail of L
//    for each node m with an edge e from n to m do
//        remove edge e from the graph
//        if m has no other incoming edges then
//            insert m into S
//if graph has edges then
//    return error   (graph has at least one cycle)
//else
//    return L   (a topologically sorted order)
func process(all []string) {
	nodes := make(map[string][]string)
	for _, line := range all {
		fields := strings.Fields(line)
		id, nextid := fields[1], fields[7]
		deps := nodes[id]
		deps = append(deps, nextid)
		nodes[id] = deps
	}

	var start = ""

	for k, _ := range nodes {
		found := true
		for _, v := range nodes {
			for _, dep := range v {
				if k == dep {
					found = false
				}
			}
		}
		if found {
			start = k
			break
		}
	}
	fmt.Println("start", start)
	L := make([]map[string][]string, 0)
	item := make(map[string][]string)
	item[start] = nodes[start]
	L = append(L, item)
	delete(nodes, start)
	fmt.Println("L", L)
	fmt.Println("ALL", nodes)

	for _, v := range item {
		for _,next := range v {
			item := make(map[string][]string)
			item[next] = nodes[next]
			L = append(L, item)
			delete(nodes, next)
			fmt.Println("L", L)
			fmt.Println("ALL", nodes)
		}
	}
}


func (n *node) makeGraph(start string, edges map[string][]string) {


}

func main() {
	input := make([]string, 0)
	r := bufio.NewReader(os.Stdin)
	for {
		line, err := r.ReadString('\n')
		if strings.TrimSpace(line) == "" && err == io.EOF {
			root.init()
			root.order()
			fmt.Println(ordered)
			os.Exit(0)
		}
		input = append(input, (strings.TrimSpace(line)))
	}

}