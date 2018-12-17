package main

import (
	"container/ring"
	"fmt"
)

var r = ring.New(1)

func playMarble(m int) (int, bool) {
	score := 0
	// we have a 23 divisible marble
	if m%23 == 0 {
		//go back 7 positions...
		r = r.Move(-7)
		addToScore := r.Next().Value.(int)
		score = score + m + addToScore
		r.Unlink(1)
		return score, true
	}
	r = r.Move(2)
	s := ring.New(1)
	s.Value = m
	r.Link(s)
	return 0, false
}

func play(players int, lastMarble int) int {
	r = ring.New(1)
	r.Value = 0
	marble := 0
	scores := make([]int, players)
	for {
		marble++
		score, _ := playMarble(marble)
		player := marble % players
		scores[player] += score
		if marble == lastMarble {
			return maxScore(scores)
		}
	}
}

func printRing(player, score int) {
	fmt.Print(player, " ")
	r.Do(func(p interface{}) {
		fmt.Print(p.(int), " ")
	})
	fmt.Println("score", score)
}

func maxScore(scores []int) (max int) {
	for _, v := range scores {
		if v > max {
			max = v
		}
	}
	return
}

func main() {
	fmt.Println("max score game 1:", play(419, 72164))
	fmt.Println("max score game 2:", play(419, 72164*100))
}

//419 players; last marble is worth 72164 points
