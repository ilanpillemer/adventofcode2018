package main

import (
	"fmt"
	"testing"
)

func TestProcess(t *testing.T) {
	all := []string{
		"Step C must be finished before step A can begin.",
		"Step C must be finished before step F can begin.",
		"Step A must be finished before step B can begin.",
		"Step A must be finished before step D can begin.",
		"Step B must be finished before step E can begin.",
		"Step D must be finished before step E can begin.",
		"Step F must be finished before step E can begin.",
	}
	edges := process(all)
	s := starts(edges)
	fmt.Println(edges)
	fmt.Println("starts", s)
	root := &node{}
	root.init(s)
	root.construct(edges, 0)

	tests := []struct {
		index string
		want  string
	}{
		{"ROOT", "&{ROOT [{ROOT C}]}"},
		{"A", "&{A [{A B} {A D}]}"},
		{"B", "&{B [{B E}]}"},
		{"C", "&{C [{C A} {C F}]}"},
		{"D", "&{D [{D E}]}"},
		{"E", "&{E []}"},
		{"F", "&{F [{F E}]}"},
	}

	for _, test := range tests {

		if got := fmt.Sprint(nodes[test.index]); got != test.want {
			t.Errorf("want [%s], got [%s]", test.want, got)
		}
	}
}
