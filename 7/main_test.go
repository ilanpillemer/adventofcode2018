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
	process(all)
	root.init()

	root.display(0)

	root.order()
	output := fmt.Sprint(ordered)
	if output != "CABDFE" {
		t.Errorf("want [%s] got [%s]\n", "CABDFE", output)
	}
}

func TestProcess2(t *testing.T) {

	all := []string{
		"Step A must be finished before step B can begin.",
		"Step A must be finished before step D can begin.",
		"Step B must be finished before step E can begin.",
		"Step D must be finished before step E can begin.",
		"Step F must be finished before step E can begin.",
		"Step C must be finished before step A can begin.",
		"Step C must be finished before step F can begin.",
	}
		process(all)
	root.init()

	root.display(0)

	root.order()
	output := fmt.Sprint(ordered)
	if output != "CABDFE" {
		t.Errorf("want [%s] got [%s]\n", "CABDFE", output)
	}
}

