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
	initSleigh("_ABCDEF", 0)

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
			fmt.Printf("time is waiting for %s\n", worker1)
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
	initSleigh("_ABCDEF", 0)
	count := 0
	go func() {
		worker1 := "ozzy"
		worker2 := "cesar"
		for {
			count++
			// make sure available is updated for next tick
			available.update()
			worker1C := make(chan struct{})
			worker2C := make(chan struct{})
			go worker(worker1, worker1C)
			go worker(worker2, worker2C)
			fmt.Println("time is about to tick")
			ticker <- struct{}{}
			ticker <- struct{}{}

			fmt.Println("time ticked ")
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
	fmt.Println("ticks", count-1)

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

	if (count - 1) != 15 {
		t.Errorf("wanted 15 got %d", count-1)
	}
}
