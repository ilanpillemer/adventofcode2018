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
	//initSleigh("ABCDEF")
	ticker := make(chan struct{})

	//	wg.Add(10000)
	//go worker("santa", ticker)

	// closure begins here

	type cval struct {
		step string
		cost int
	}
	cmap := make(map[string]cval)
	santa := make(chan struct{})
	santago := func(<-chan struct{}) {
		//wait for time to tick
		fmt.Println("santa is waiting for time")
		<-ticker
		fmt.Println("time ticked for Santa")
		//check closure
		state := cmap["santa"]
		// if closure has something that can be finished this time slice
		if state.cost == 1 {
			//complete this step
			fmt.Println("santa is completing step", state.step)
			do("santa", state.step)
			//update closure
			cmap["santa"] = cval{}
		}

		// if closure has something that be worked on but is not finishable
		if state.cost > 1 {
			fmt.Println("santa is working on step", state.step)
			cmap["santa"] = cval{state.step, state.cost - 1}
			fmt.Printf("closure map now looks like this %v\n", cmap)
		}

		// if was idle previously and ready to take on work, check if there is anything available
		if state.cost == 0 {
			// if there is something available
			fmt.Printf("is there is anything available for santa? %+v\n", available)
			if available.size() != 0 {
				r, cost := available.pop()
				fmt.Printf("Yes, %s is available with cost %d.\n", string(r), cost)
				cmap["santa"] = cval{string(r), cost}
				fmt.Printf("closure map now looks like this %v\n", cmap)
			}

			//repeat same checks as above
			state := cmap["santa"]
			if state.cost == 1 {
				//complete this step
				fmt.Println("santa is completing step", state.step)
				do("santa", state.step)
				//update closure
				cmap["santa"] = cval{}
			}
			if state.cost > 1 {
				fmt.Println("santa works at", state.step)
				cmap["santa"] = cval{state.step, state.cost - 1}
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
		fmt.Println("moment is over for santa")
		santa <- struct{}{}

	}

	go func() {
		for i := 0; i < 600; i++ {
			// make sure available is updated for next tick
			available.update()
			go santago(santa)
			fmt.Println("time is about to tick")
			ticker <- struct{}{}

			fmt.Println("time ticked globally")
			fmt.Println("time is waiting for santa")
			<-santa // wait for santa
			fmt.Println("santa announced that he had a moment")
			if len(available.GetTodo()) == 0 {
			fmt.Println("assembled with steps",order)
				assembled <- struct{}{}
			}
			//fmt.Printf("%#v\n", available)
			//available.update()
			//fmt.Println(available)
		}
	}()
	fmt.Println("nodes", nodes)
	fmt.Println("C", nodes["C"])
	//assemble(nodes["C"])
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
	tickerC := make(chan struct{})
	tickerO := make(chan struct{})

	//	go worker("cesar", tickerC)
	//	go worker("ozzy", tickerO)
	go func() {
		for i := 0; i < 60; i++ {
			//fmt.Println("****** Second ****** ", i-31)
			// they must wait for each other here
			// any newly available tasks should only be released next tick
			tickerC <- struct{}{}
			tickerO <- struct{}{}
			available.update()

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
