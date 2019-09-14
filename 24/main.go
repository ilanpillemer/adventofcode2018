package main

import (
	"fmt"
	"sort"
)

type attackType int

const (
	radiation attackType = iota
	bludgeoning
	fire
	slashing
)

type unit struct {
	hp int
}

type group struct {
	units      []unit
	damage     int
	initiative int
	attack     attackType
	immune     []attackType
	weak       []attackType
}

func (a army) Len() int {
	return len(a.groups)
}

func (a army) Less(i, j int) bool {
	if a.groups[i].power() == a.groups[j].power() {
		return a.groups[i].initiative > a.groups[j].initiative
	}

	return a.groups[i].power() > a.groups[j].power()
}

// Swap swaps the elements with indexes i and j.
func (a army) Swap(i, j int) {
	a.groups[i], a.groups[j] = a.groups[j], a.groups[i]
}

type army struct {
	groups []group
}

func (g group) power() int {
	if len(g.units) == 0 {
		return 0
	}

	return len(g.units) * g.damage
}

func (g group) possible(t group) int {
	if t.isImmune(g.attack) {
		return 0
	}
	if t.isWeak(g.attack) {
		return g.power() * 2
	}
	return g.power()
}

func (g group) isImmune(a attackType) bool {
	for _, v := range g.immune {
		if v == a {
			return true
		}
	}
	return false
}

func (g group) isWeak(a attackType) bool {
	for _, v := range g.weak {
		if v == a {
			return true
		}
	}
	return false
}

func main() {
	is := army{}
	team := army{}
	fmt.Println(is, team)
}

func (a *army) target(t *army) {
	sort.Sort(a)
	fmt.Println(a)
//	options := t.groups
	//	for i, v := range a.groups {
	//
	//	}

}
