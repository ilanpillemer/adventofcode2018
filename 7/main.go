package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	nodes := make(map[string]node)
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

func main() {
	r := bufio.NewReader(os.Stdin)
	all := make([]string, 0)
	for {
		line, err := r.ReadString('\n')
		if strings.TrimSpace(line) == "" && err == io.EOF {
			edges := process(all)
			nodes := createNodes(edges)
				fmt.Println("ends", ends(nodes))
	fmt.Println("starts", starts(nodes))
			fmt.Println(edges)
		
			os.Exit(0)
		}
		all = append(all, strings.TrimSpace(line))
	}
}
