package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestPower(t *testing.T) {
	g := group{}
	g.units = 18
	g.damage = 8

	want := 144
	got := g.power()
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}

func TestSort(t *testing.T) {
	a := army{}
	g0 := group{}
	g0.units = 18
	g0.damage = 1

	g1 := group{}
	g1.units = 1
	g1.damage = 8

	g2 := group{}
	g2.units = 180

	g2.damage = 8

	a.groups = []group{g0, g1, g2}
	sort.Sort(a)
	if a.groups[0].power() != 1440 {
		t.Errorf("want 8 got %d", a.groups[0].power())
	}

	if a.groups[1].power() != 18 {
		t.Errorf("want 18 got %d", a.groups[1].power())
	}

	if a.groups[2].power() != 8 {
		t.Errorf("want 1440 got %d", a.groups[2].power())
	}

}

//Immune System:
//17 units each with 5390 hit points (weak to radiation, bludgeoning) with
// an attack that does 4507 fire damage at initiative 2
//989 units each with 1274 hit points (immune to fire; weak to bludgeoning,
// slashing) with an attack that does 25 slashing damage at initiative 3

//Infection:
//801 units each with 4706 hit points (weak to radiation) with an attack
// that does 116 bludgeoning damage at initiative 1
//4485 units each with 2961 hit points (immune to radiation; weak to fire,
// cold) with an attack that does 12 slashing damage at initiative 4

func TestTarget(t *testing.T) {
	a := army{id: "Immune System"}
	b := army{id: "Infection"}

	g1 := group{}
	g1.id = "Group 1"
	g1.units = 17
	g1.hp = 5390
	g1.weak = []attackType{radiation, bludgeoning}
	g1.damage = 4507
	g1.initiative = 2
	g1.attack = fire

	g2 := group{}
	g2.id = "Group 2"
	g2.units = 989
	g2.hp = 1274
	g2.immune = []attackType{fire}
	g2.weak = []attackType{bludgeoning, slashing}
	g2.damage = 25
	g2.initiative = 3
	g2.attack = slashing

	g3 := group{}
	g3.id = "Group 1"
	g3.units = 801
	g3.hp = 4706

	g3.weak = []attackType{radiation}
	g3.damage = 116
	g3.initiative = 1
	g3.attack = bludgeoning

	g4 := group{}
	g4.id = "Group 2"
	g4.units = 4485
	g4.hp = 2961
	g4.immune = []attackType{radiation}
	g4.weak = []attackType{fire, cold}
	g4.damage = 12
	g4.initiative = 4
	g4.attack = slashing

	a.groups = []group{g1, g2}
	b.groups = []group{g3, g4}

	// FIGHT!!!
	for len(a.groups) > 0 && len(b.groups) > 0 {
		atargets := target(&a, &b)
		btargets := target(&b, &a)
		sort.Sort(a)
		sort.Sort(b)
		fmt.Println()
		fmt.Println(a.id)
		for _, v := range a.groups {
			fmt.Printf("%s contains %d units\n", v.id, v.units)
		}
		fmt.Println()
		fmt.Println(b.id)
		fmt.Println()
		for _, v := range b.groups {
			fmt.Printf("%s contains %d units\n", v.id, v.units)
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
		score += v.units
	}
	fmt.Println()
	fmt.Println(b.id)
	fmt.Println()
	for _, v := range b.groups {
		fmt.Printf("%s contains %d units\n", v.id, v.units)
		score += v.units
	}

	fmt.Println("final score", score)

}
