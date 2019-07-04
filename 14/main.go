package main

import "flag"
import "fmt"
import "strconv"

// You finally have a chance to look at all of the produce moving
// around. Chocolate, cinnamon, mint, chili peppers, nutmeg,
// vanilla... the Elves must be growing these plants to make hot
// chocolate! As you realize this, you hear a conversation in the
// distance. When you go to investigate, you discover two Elves in
// what appears to be a makeshift underground kitchen/laboratory.

// The Elves are trying to come up with the ultimate hot chocolate
// recipe; they're even maintaining a scoreboard which tracks the
// quality score (0-9) of each recipe.

// Only two recipes are on the board: the first recipe got a score of
// 3, the second, 7. Each of the two Elves has a current recipe: the
// first Elf starts with the first recipe, and the second Elf starts
// with the second recipe.

// To create new recipes, the two Elves combine their current recipes.
// This creates new recipes from the digits of the sum of the current
// recipes' scores. With the current recipes' scores of 3 and 7, their
// sum is 10, and so two new recipes would be created: the first with
// score 1 and the second with score 0. If the current recipes' scores
// were 2 and 3, the sum, 5, would only create one recipe (with a
// score of 5) with its single digit.

// The new recipes are added to the end of the scoreboard in the order
// they are created. So, after the first round, the scoreboard is 3,
// 7, 1, 0.

// After all new recipes are added to the scoreboard, each Elf picks a
// new current recipe. To do this, the Elf steps forward through the
// scoreboard a number of recipes equal to 1 plus the score of their
// current recipe. So, after the first round, the first Elf moves
// forward 1 + 3 = 4 times, while the second Elf moves forward 1 + 7 =
// 8 times. If they run out of recipes, they loop back around to the
// beginning. After the first round, both Elves happen to loop around
// until they land on the same recipe that they had in the beginning;
// in general, they will move to different recipes.

var gen = flag.Int("gen", 5, "default number of iterations of the recipes")

type elf struct {
	current int
	backing *[]int
}

func (e *elf) next() {
	l := len(*(e.backing))
	e.current = (e.current + e.score() + 1) % l
}

func (e *elf) score() int {
	b := e.backing
	return (*b)[e.current]
}

func main() {
	flag.Parse()
	recipes := []int{3, 7}
	e0 := elf{0, &recipes}
	e1 := elf{1, &recipes}
	fmt.Println("initialising")
	fmt.Println(recipes)
	for len(recipes) < *gen+10+2 {
		e0.next()
		e1.next()
		sum := strconv.Itoa(e0.score() + e1.score())
		for _, v := range sum {
			recipes = append(recipes, int(v-'0'))
		}
	}
	fmt.Printf("%v\n", recipes[*gen:*gen+10])
}
