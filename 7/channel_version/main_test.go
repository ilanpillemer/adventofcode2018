package main

import (
	"fmt"
	"testing"
)

func TestProcessOneWorker(t *testing.T) {
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
	root := addRoot(starts(nodes))
	initSleigh("_ABCDEF")

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
			}
		}
	}()
	assemble(root)

	tests := []struct {
		index string
		want  string
	}{
		{"A", "{id:A outs:[{from:A to:B} {from:A to:D}] ins:[{from:C to:A}]}"},
		{"B", "{id:B outs:[{from:B to:E}] ins:[{from:A to:B}]}"},
		{"C", "{id:C outs:[{from:C to:A} {from:C to:F}] ins:[{from:_ to:C}]}"},
		{"D", "{id:D outs:[{from:D to:E}] ins:[{from:A to:D}]}"},
		{"E", "{id:E outs:[] ins:[{from:B to:E} {from:D to:E} {from:F to:E}]}"},
		{"F", "{id:F outs:[{from:F to:E}] ins:[{from:C to:F}]}"},
	}

	for _, test := range tests {

		if got := fmt.Sprintf("%+v", nodes[test.index]); got != test.want {
			t.Errorf("want [%s], got [%s]", test.want, got)
		}
	}

	if order != "_CABDFE" {
		t.Errorf("want %s got %s", "CABDFE", order)
	}
}

func TestProcessTwoWorkers(t *testing.T) {
	fmt.Println(rune(0))
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
	root := addRoot(starts(nodes))
	initSleigh("_ABCDEF")
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
			}
		}
	}()
	assemble(root)

	tests := []struct {
		index string
		want  string
	}{
		{"A", "{id:A outs:[{from:A to:B} {from:A to:D}] ins:[{from:C to:A}]}"},
		{"B", "{id:B outs:[{from:B to:E}] ins:[{from:A to:B}]}"},
		{"C", "{id:C outs:[{from:C to:A} {from:C to:F}] ins:[{from:_ to:C}]}"},
		{"D", "{id:D outs:[{from:D to:E}] ins:[{from:A to:D}]}"},
		{"E", "{id:E outs:[] ins:[{from:B to:E} {from:D to:E} {from:F to:E}]}"},
		{"F", "{id:F outs:[{from:F to:E}] ins:[{from:C to:F}]}"},
	}

	for _, test := range tests {

		if got := fmt.Sprintf("%+v", nodes[test.index]); got != test.want {
			t.Errorf("want [%s], got [%s]", test.want, got)
		}
	}
	//_CABFDE CABFDE
	if order != "_CABFDE" {
		t.Errorf("want %s got %s", "CABFDE", order)
	}
}
