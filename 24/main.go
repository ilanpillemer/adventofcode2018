package main

import (
	"fmt"
)

type attackType int

const (
	radiation attackType = iota
	bludgeoning
	fire
	slashing
	cold
	love
)

const maxInit = 4

type group struct {
	id         string
	units      int
	hp         int
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
	id     string
	groups []group
}

func (g group) power() int {
	if g.units == 0 {
		return 0
	}

	return g.units * g.damage
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

func target(a *army, t *army) map[int]int {
	targets := map[int]int{}
	taken := map[int]bool{}
	for i, v := range a.groups {
		maxDamage := -1
		var selected group
		selectedIndex := -1

		for i2, v2 := range t.groups {
			if taken[i2] {
				continue
			}
			damage := v.possible(v2)
			if damage > maxDamage {
				maxDamage = damage
				selectedIndex = i2
				selected = v2
			}
			if damage == maxDamage && v2.initiative < selected.initiative {
				selectedIndex = i2
				selected = v2
			}
		}
		if maxDamage > 0 {
			taken[selectedIndex] = true
			targets[i] = selectedIndex
		}
		maxDamage = -1
	}
	fmt.Println(targets)
	return targets
}

func attack(a *army, b *army, at map[int]int, bt map[int]int) (*army, *army) {
	fmt.Println()
	//top:
	for i := maxInit; i > 0; i-- {
		//fmt.Println("init", i)
		for j, v := range a.groups {
			//dead
			if v.units == 0 {
				continue
			}
			//no target
			if _,ok := at[j]; !ok {
			continue
			}
			if v.initiative == i {
				damage := a.groups[j].possible(b.groups[at[j]])
				hp := b.groups[at[j]].hp
				units := b.groups[at[j]].units
				//fmt.Println("damage mod hp", damage, hp)
				tot := damage / hp
				units = units - tot
				b.groups[at[j]].units = units
				fmt.Printf("%s %s attacks defending %s %s, killing %d units\n", a.id, a.groups[j].id, b.id, b.groups[at[j]].id, tot)
			}
		}
		for j, v := range b.groups {
			if v.units == 0 {
				continue
			}
						//no target
			if _,ok := bt[j]; !ok {
			continue
			}
			if v.initiative == i {
				damage := b.groups[j].possible(a.groups[bt[j]])
				hp := a.groups[bt[j]].hp
				units := a.groups[bt[j]].units
				//fmt.Println("damage mod hp", damage, hp)
				tot := damage / hp
				if tot > units {
					tot = units
				}
				units = units - tot
				a.groups[bt[j]].units = units
				fmt.Printf("%s %s attacks defending %s %s, killing %d units\n", b.id, b.groups[j].id, a.id, a.groups[bt[j]].id, tot)
			}
		}
	}
	for i := len(a.groups) - 1; i > -1; i-- {
		if a.groups[i].units == 0 {
			a.groups = append(a.groups[:i], a.groups[i+1:]...)
		}
	}
	for i := len(b.groups) - 1; i > -1; i-- {
		if b.groups[i].units == 0 {
			b.groups = append(b.groups[:i], b.groups[i+1:]...)
		}
	}
	return a, b

}
