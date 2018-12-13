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
	id    string
	edges []edge
}

var nodes = make(map[string]*node)
var added = make(map[string]bool)

func (n *node) init(s []string) {

	n.id = "ROOT"
	n.edges = make([]edge, 0)
	for _, id := range s {
		edge := edge{
			from: "ROOT",
			to:   id,
		}
		n.edges = append(n.edges, edge)
	}
	sort.Slice(n.edges, func(i, j int) bool {
		return n.edges[i].to < n.edges[j].to
	})
	nodes[n.id] = n
}

func (n *node) construct(edges []edge, depth int) {
	//	fmt.Print(".")
	//	fmt.Println(edges)

	sort.Slice(n.edges, func(i, j int) bool {
		return n.edges[i].to < n.edges[j].to
	})
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].to < edges[j].to
	})

	// all the edges in the current node
	for _, ne := range n.edges {
		// all possible edges
		for i, me := range edges {
			if ne.to == me.from {
				//				found = true
				next, ok := nodes[ne.to]
				if !ok {
					next = &node{id: ne.to}
				}
				lessedges := make([]edge, len(edges))
				copy(lessedges, edges)

				lessedges = append(lessedges[:i], lessedges[i+1:]...)
				fmt.Println(lessedges)
				_, ok = added[next.id+me.String()]
				if !ok {
					next.edges = append(next.edges, me)
					added[next.id+me.String()] = true
					//					fmt.Println(len(edges))
					//					fmt.Println(i)
					//					fmt.Println(edges)

				}

				nodes[ne.to] = next
				next.construct(lessedges, depth+1)
			}
		}
		next, ok := nodes[ne.to]
		if !ok {
			next = &node{id: ne.to}
			nodes[ne.to] = next
			next.construct(edges, depth+1)
		}
	}
}

func (e edge) String() string {
	return e.from + "->" + e.to
}

func starts(edges []edge) (starts []string) {
	froms := make(map[string]bool)
	tos := make(map[string]bool)
	for _, e := range edges {
		froms[e.from] = true
		tos[e.to] = true
	}

	for k, _ := range tos {
		delete(froms, k)
	}

	for k, _ := range froms {
		starts = append(starts, k)
	}
	return
}

//var root = &node{}

func process(input []string) []edge {
	edges := make([]edge, 0)
	for _, line := range input {
		fields := strings.Fields(line)
		e := edge{fields[1], fields[7]}
		edges = append(edges, e)
	}
	//fmt.Printf("edges : %v\n", edges)
	return edges
}

func main() {
	r := bufio.NewReader(os.Stdin)
	all := make([]string, 0)
	for {
		line, err := r.ReadString('\n')
		if strings.TrimSpace(line) == "" && err == io.EOF {
			edges := process(all)
			s := starts(edges)
			fmt.Println(edges)
			fmt.Println("starts", s)
			root := &node{}
			root.init(s)
			root.construct(edges, 0)
			os.Exit(0)
		}
		all = append(all, strings.TrimSpace(line))
	}
}
