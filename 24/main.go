package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var boost = flag.Int("b", 0, "boost")

type attackType int

const (
	radiation attackType = iota
	bludgeoning
	fire
	slashing
	cold
)

const maxInit = 25

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
	flag.Parse()
	a := army{id: "Immune System"}
	b := army{id: "Infection"}
	a.groups = []group{}
	b.groups = []group{}
	g1 := 1
	g2 := 1

	in := bufio.NewScanner(os.Stdin)
	state := ""
	for in.Scan() {
		line := in.Text()
		//fmt.Println(line)
		switch {
		case strings.Contains(line, "Immune System"):
			state = "immune"
			continue
		case strings.Contains(line, "Infection"):
			state = "infect"
			continue
		}
		if len(line) == 0 {
			continue
		}
		switch state {
		case "immune":
			gname := fmt.Sprintf("Group %d", g1)
			g := group{id: gname}
			g1++
			fields := strings.Fields(line)
			substate := ""
			for i, field := range fields {
				if i == 0 {
					g.units, _ = strconv.Atoi(field)
				}
				if i == 4 {
					g.hp, _ = strconv.Atoi(field)
				}
				switch {
				case strings.Contains(field, "weak"):
					substate = "weak"
				case strings.Contains(field, "immune"):
					substate = "immune"
				case strings.Contains(field, "attack"):
					substate = "attack"
				case strings.Contains(field, "damage"):
					substate = "damage"
				case strings.Contains(field, "initiative"):
					substate = "initiative"
				}
				//fmt.Println("substate", substate)
				switch substate {
				case "immune":
					//fmt.Println(field)
					g = addImmune(g, field)
				case "weak":
					//fmt.Println(field)
					g = addWeak(g, field)
				case "attack":
					g = addAttack(g, field)


				case "initiative":
					g = addInit(g, field)
				}
			}
				g.damage = g.damage + *boost
					fmt.Println(g.damage)
			a.groups = append(a.groups, g)
		case "infect":
			gname := fmt.Sprintf("Group %d", g2)
			g := group{id: gname}
			g2++
			fields := strings.Fields(line)
			substate := ""
			for i, field := range fields {
				if i == 0 {
					g.units, _ = strconv.Atoi(field)
				}
				if i == 4 {
					g.hp, _ = strconv.Atoi(field)
				}

				switch {
				case strings.Contains(field, "weak"):
					substate = "weak"
				case strings.Contains(field, "immune"):
					substate = "immune"
				case strings.Contains(field, "attack"):
					substate = "attack"
				case strings.Contains(field, "damage"):
					substate = "damage"
				case strings.Contains(field, "initiative"):
					substate = "initiative"
				}
				//fmt.Println("substate", substate)
				switch substate {
				case "immune":
					//	fmt.Println(field)
					g = addImmune(g, field)
				case "weak":
					//	fmt.Println(field)
					g = addWeak(g, field)
				case "attack":
					g = addAttack(g, field)
				case "initiative":
					g = addInit(g, field)
				}

			}
			b.groups = append(b.groups, g)
		}

	}

	for len(a.groups) > 0 && len(b.groups) > 0 {
		sort.Sort(a)
		sort.Sort(b)
		atargets := target(&a, &b)
		btargets := target(&b, &a)

		fmt.Println()
		fmt.Println(a.id)
		for _, v := range a.groups {
			fmt.Printf("%s contains %d units\n", v.id, v.units)
			if v.units < 0 {
				panic("units should not be less than zero")
			}
		}
		fmt.Println()
		fmt.Println(b.id)
		fmt.Println()
		for _, v := range b.groups {
			fmt.Printf("%s contains %d units\n", v.id, v.units)
			if v.units < 0 {
				panic("units should not be less than zero")
			}
		}
		for k, v := range atargets {
			fmt.Printf("%s %s a would deal defending %s %s %d damage\n", a.id, a.groups[k].id, b.id, b.groups[v].id, a.groups[k].possible(b.groups[v]))
		}
		for k, v := range btargets {
			fmt.Printf("%s %s a would deal defending %s %s %d damage\n", b.id, b.groups[k].id, a.id, a.groups[v].id, b.groups[k].possible(a.groups[v]))
		}

		attack(&a, &b, atargets, btargets)
	}

	score := 0
	fmt.Println()
	fmt.Println(a.id)
	for _, v := range a.groups {
		fmt.Printf("%s contains %d units\n", v.id, v.units)
		if v.units < 0 {
			panic("units should not be less than zero")
		}
		score += v.units
	}
	fmt.Println("score", score)
	fmt.Println()
	fmt.Println(b.id)
	fmt.Println()
	for _, v := range b.groups {
		fmt.Printf("%s contains %d units\n", v.id, v.units)
		if v.units < 0 {
			panic("units should not be less than zero")
		}
		score += v.units
	}

	fmt.Println("final score", score)
}

//	bludgeoning
//	fire
//	slashing
//	cold

func addImmune(g group, s string) group {
	if g.immune == nil {
		g.immune = []attackType{}
	}
	switch {
	case strings.Contains(s, "fire"):
		g.immune = append(g.immune, fire)
	case strings.Contains(s, "radiation"):
		g.immune = append(g.immune, radiation)
	case strings.Contains(s, "bludgeoning"):
		g.immune = append(g.immune, bludgeoning)
	case strings.Contains(s, "slashing"):
		g.immune = append(g.immune, slashing)
	case strings.Contains(s, "cold"):
		g.immune = append(g.immune, cold)

	}
	return g
}

func addWeak(g group, s string) group {
	if g.weak == nil {
		g.weak = []attackType{}
	}
	switch {
	case strings.Contains(s, "fire"):
		g.weak = append(g.weak, fire)
	case strings.Contains(s, "radiation"):
		g.weak = append(g.weak, radiation)
	case strings.Contains(s, "bludgeoning"):
		g.weak = append(g.weak, bludgeoning)
	case strings.Contains(s, "slashing"):
		g.weak = append(g.weak, slashing)
	case strings.Contains(s, "cold"):
		g.weak = append(g.weak, cold)

	}
	return g
}

func addInit(g group, s string) group {
	if init, err := strconv.Atoi(s); err == nil {
		g.initiative = init
	}
	return g
}

func addAttack(g group, s string) group {

	if damage, err := strconv.Atoi(s); err == nil {
		g.damage = damage
	}

	switch {
	case strings.Contains(s, "fire"):
		g.attack = fire
	case strings.Contains(s, "radiation"):
		g.attack = radiation
	case strings.Contains(s, "bludgeoning"):
		g.attack = bludgeoning
	case strings.Contains(s, "slashing"):
		g.attack = slashing
	case strings.Contains(s, "cold"):
		g.attack = cold

	}
	return g
}

func target(a *army, t *army) map[int]int {
	targets := map[int]int{}
	taken := map[int]bool{}
	for i, v := range a.groups {
		if v.units == 0 {
			continue
		}
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
			} else if damage == maxDamage {
				//				fmt.Println("power", v2.power(), "vs", selected.power())
				//				fmt.Println("init", v2.initiative, "vs", selected.initiative)

				switch {
				case v2.power() > selected.power():
					fmt.Println(v2.power(), "vs", selected.power())
					selectedIndex = i2
					selected = v2
				case v2.power() < selected.power():
				// no-op
				case v2.initiative > selected.initiative:
					selectedIndex = i2
					selected = v2
				}
			}
		}
		if maxDamage > 0 {
			taken[selectedIndex] = true
			targets[i] = selectedIndex
		}
		maxDamage = -1
		selectedIndex = -1
	}
	fmt.Println(targets)
	return targets
}

func attack(a *army, b *army, at map[int]int, bt map[int]int) (*army, *army) {
	fmt.Println()

	for i := maxInit; i > 0; i-- {
		//fmt.Println("init", i)
		for j, v := range a.groups {
			//dead
			if v.units == 0 {
				continue
			}
			//no target
			if _, ok := at[j]; !ok {
				continue
			}
			if v.initiative == i {
				damage := a.groups[j].possible(b.groups[at[j]])
				hp := b.groups[at[j]].hp
				units := b.groups[at[j]].units
				fmt.Println("damage", damage, "mod hp", hp, "=", damage/hp)
				tot := damage / hp
				if tot >= units {
					tot = units
				}
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
			if _, ok := bt[j]; !ok {
				continue
			}
			if v.initiative == i {
				damage := b.groups[j].possible(a.groups[bt[j]])
				hp := a.groups[bt[j]].hp
				units := a.groups[bt[j]].units
				fmt.Println("damage", damage, "mod hp", hp, "=", damage/hp)
				tot := damage / hp
				if tot >= units {
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
