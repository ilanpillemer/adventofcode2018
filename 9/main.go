package main

import (
	"fmt"
)

type circle struct {
	circle  []int
	current int
}

func (c *circle) playMarble(m int) (int, bool) {
	score := 0
	if m%23 == 0 {

		delpos := c.current - 7
		if delpos < 0 {
			delpos = len(c.circle) + delpos
		}
		c.current = (delpos)
		score = score + m + c.circle[c.current]
		c.circle = append(c.circle[:c.current], c.circle[c.current+1:]...)
		return score, true
	}
	c.current = (c.current + 2) % len(c.circle)
	if c.current == m {
		fmt.Println("wooooooo")
		c.circle = append(c.circle, m)
		return 0, false
	}
	c.circle = append(c.circle[:c.current], append([]int{m}, c.circle[c.current:]...)...)
	return 0, false
}

func play(players int, lastMarble int) int {
	c := circle{
		circle:  make([]int, 1),
		current: 0,
	}
	marble := 0

	fmt.Println(c.circle)

	for {
		marble++
		score, _ := c.playMarble(marble)
		//		scoreString := ""
		//		if ok {
		//			scoreString = fmt.Sprintf("scored %d", score)
		//			//fmt.Println(scoreString)
		//		}

		//		if len(c.circle) < 20 {
		//			fmt.Println(c.circle, scoreString)
		//		} else {
		//			fmt.Println(c.circle[:20], scoreString)
		//		}

		if score == lastMarble {
			fmt.Println("game is over with last marble worth", score)
		}

		//fmt.Println("playing marble", marble)

		//		if marble == 25 {
		//			os.Exit(0)
		//		}

	}

	return 0
}

func main() {

}