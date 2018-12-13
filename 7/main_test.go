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
	nodes := createNodes(edges)

	fmt.Println("ends", ends(nodes))
	fmt.Println("starts", starts(nodes))
	fmt.Println(edges)
	fmt.Println(nodes)


	tests := []struct {
		index string
		want  string
	}{
		{"A", "{id:A outs:[{from:A to:B} {from:A to:D}] ins:[{from:C to:A}]}"},
		{"B", "{id:B outs:[{from:B to:E}] ins:[{from:A to:B}]}"},
		{"C", "{id:C outs:[{from:C to:A} {from:C to:F}] ins:[]}"},
		{"D", "{id:D outs:[{from:D to:E}] ins:[{from:A to:D}]}"},
		{"E", "{id:E outs:[] ins:[{from:B to:E} {from:D to:E} {from:F to:E}]}"},
		{"F", "{id:F outs:[{from:F to:E}] ins:[{from:C to:F}]}"},
	}

	for _, test := range tests {

		if got := fmt.Sprintf("%+v", nodes[test.index]); got != test.want {
			t.Errorf("want [%s], got [%s]", test.want, got)
		}
	}
}
