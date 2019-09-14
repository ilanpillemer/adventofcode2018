package main

import "fmt"

type attackType int

const (
	radiation attackType = iota
	bludgeoning
	fire
	slashing
)

type unit struct {
	hp     int
	damage int
	attack attackType
	immune []attackType
	weak   []attackType
}

type army struct {
	group []unit
}

func power(group []unit) int {
	if len(group) == 0 {
		return 0
	}

	return len(group) * group[0].damage
}

func main() {
	is := army{}
	team := army{}
	fmt.Println(is, team)
}
