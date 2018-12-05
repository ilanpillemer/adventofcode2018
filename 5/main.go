package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type polymer string

var explosions = makeExplosions()
var alphabeta = "abcdefghijklmnopqrstuvwxyz"

func (p polymer) react() polymer {
	original := p
	for _, e := range explosions {
		p = polymer(strings.Replace(string(p), e, "", -1))
	}
	if original == p {
		return p
	}
	return p.react()
}

func (p polymer) remove(b string) polymer {
	l, u := strings.ToLower(b), strings.ToUpper(b)
	// I know, I KNOW
	p = polymer(strings.Replace(string(p), l, "", -1))
	return polymer(strings.Replace(string(p), u, "", -1))
}

func (p polymer) String() string {
	return string(p)
}

func (p polymer) units() int {
	return len(p)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	input, _ := r.ReadString('\n')
	p := polymer(strings.TrimSpace(input)).react()
	fmt.Printf("it has when fully reacted [%d] units\n", p.units())
	for _, c := range alphabeta {
		s := fmt.Sprintf("%c", c)
		units := polymer(strings.TrimSpace(input)).remove(s).react().units()
		fmt.Println(s, units)
	}
}

func makeExplosions() [52]string {
	var explosions [52]string
	lower, upper := alphabeta, strings.ToUpper(alphabeta)
	for i := 0; i < 26; i++ {
		explosions[i] = fmt.Sprintf("%c%c", lower[i], upper[i])
		explosions[i+26] = fmt.Sprintf("%c%c", upper[i], lower[i])
	}
	return explosions
}