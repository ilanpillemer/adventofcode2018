package main

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

//[1518-11-01 00:00]
func TestTime(t *testing.T) {
	const layout = "[2006-01-02 15:04]"
	ti, _ := time.Parse(layout, "[1518-11-01 23:58]")
	fmt.Println(ti)
}

func TestNewRaw(t *testing.T) {
	r := NewRaw("[1518-11-01 00:00] Guard #10 begins shift")
	if r.date != "[1518-11-01 00:00]" {
		t.Errorf("want %s got %s", "[1518-11-01 00:00]", r.date)
	}
	if r.rest != "Guard #10 begins shift" {
		t.Errorf("want %s got %s", "Guard #10 begins shift", r.rest)
	}
}

func TestSort(t *testing.T) {
	r1 := raw{"[1518-11-01 00:30]", "falls asleep"}
	r2 := raw{"[1518-11-01 00:55]", "wakes up"}
	r3 := raw{"[1518-11-01 00:30]", "eats chocalate"}

	input := []raw{r1, r2, r3}
	sort.Sort(byDate(input))
	if fmt.Sprintf("%v", input) != "[{[1518-11-01 00:30] falls asleep} {[1518-11-01 00:30] eats chocalate} {[1518-11-01 00:55] wakes up}]" {
		t.Errorf("time not sorting...")
	}

}