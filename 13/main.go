package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type pos struct {
	y int
	x int
}

type turn int

const (
	LeftTurn turn = iota
	Straight
	RightTurn
)

type direction int

const (
	Up direction = iota
	Down
	Left
	Right
)

type cart struct {
	Next      turn
	Direction direction
}

func (c cart) String() string {
	switch c.Direction {
	case Up:
		return "^"
	case Down:
		return "v"
	case Left:
		return "<"
	case Right:
		return ">"
	}
	return "?"
}

func main() {
	in := bufio.NewScanner(os.Stdin)

	// create grid.
	// 1) load lines into an array
	raw := make([]string, 0)
	for in.Scan() != false {
		line := in.Text()
		raw = append(raw, line)
	}

	//	display(raw)
	carts := extractCarts(raw)
	grid := clear(raw)
	//display(grid)
	for {
		//dualDisplay(grid, carts)
		carts = tick(grid, carts)
	}

	for k, v := range carts {
		fmt.Println(k, v)
	}

}

func clear(raw []string) []string {
	grid := make([]string, len(raw))
	for i, line := range raw {
		line = strings.Replace(line, ">", "-", -1)
		line = strings.Replace(line, "<", "-", -1)
		line = strings.Replace(line, "^", "|", -1)
		line = strings.Replace(line, "v", "|", -1)
		grid[i] = line

	}
	return grid
}

func extractCarts(raw []string) map[pos]cart {
	carts := make(map[pos]cart)
	for y, line := range raw {
		for x, p := range line {
			switch p {
			case '>':
				carts[pos{x: x, y: y}] = cart{LeftTurn, Right}
			case '<':
				carts[pos{x: x, y: y}] = cart{LeftTurn, Left}
			case 'v':
				carts[pos{x: x, y: y}] = cart{LeftTurn, Down}
			case '^':
				carts[pos{x: x, y: y}] = cart{LeftTurn, Up}
			}
		}
	}
	return carts
}

func display(grid []string) {
	for _, l := range grid {
		fmt.Println(l)
	}
}

func dualDisplay(grid []string, carts map[pos]cart) {
	for y, line := range grid {
		for x, p := range line {
			if cart, ok := carts[pos{x: x, y: y}]; ok {
				fmt.Printf("%s", cart)
			} else {
				fmt.Printf("%c", p)
			}
		}
		fmt.Println()
	}
}

// Carts all move at the same speed; they take turns moving a single
// step at a time. They do this based on their current location: carts
// on the top row move first (acting from left to right), then carts
// on the second row move (again from left to right), then carts on
// the third row, and so on. Once each cart has moved one step, the
// process repeats; each of these loops is called a tick.
func tick(grid []string, carts map[pos]cart) map[pos]cart {
	next := make(map[pos]cart)
	for y, line := range grid {
		for x, p := range line {
			if c, ok := carts[pos{x: x, y: y}]; ok {
				switch p {
				case '-':
					switch c.Direction {
					case Right:
						if _, ok = next[pos{x: x + 1, y: y}]; ok {
							fmt.Printf("Collision at %d, %d\n", x+1, y)
							os.Exit(1)
						}
						next[pos{x: x + 1, y: y}] = cart{c.Next, c.Direction}
					case Left:
						if _, ok = next[pos{x: x - 1, y: y}]; ok {
							fmt.Printf("Collision at %d, %d\n", x-1, y)
							os.Exit(1)
						}
						next[pos{x: x - 1, y: y}] = cart{c.Next, c.Direction}
					}

				case '|':
					switch c.Direction {
					case Up:
						if _, ok = next[pos{x: x, y: y - 1}]; ok {
							fmt.Printf("Collision at %d, %d\n", x, y-1)
							os.Exit(1)
						}
						next[pos{x: x, y: y - 1}] = cart{c.Next, c.Direction}
					case Down:
						if _, ok = next[pos{x: x, y: y + 1}]; ok {
							fmt.Printf("Collision at %d, %d\n", x, y+1)
							os.Exit(1)
						}
						next[pos{x: x, y: y + 1}] = cart{c.Next, c.Direction}
					}

				case '\\':
					switch c.Direction {
					case Right:
						if _, ok = next[pos{x: x, y: y + 1}]; ok {
							fmt.Printf("Collision at %d, %d\n", x, y+1)
							os.Exit(1)
						}
						next[pos{x: x, y: y + 1}] = cart{c.Next, Down}
					case Left:
						if _, ok = next[pos{x: x, y: y - 1}]; ok {
							fmt.Printf("Collision at %d, %d\n", x, y-1)
							os.Exit(1)
						}
						next[pos{x: x, y: y - 1}] = cart{c.Next, Up}
					case Up:
						if _, ok = next[pos{x: x - 1, y: y}]; ok {
							fmt.Printf("Collision at %d, %d\n", x-1, y)
							os.Exit(1)
						}
						next[pos{x: x - 1, y: y}] = cart{c.Next, Left}
					case Down:
						if _, ok = next[pos{x: x + 1, y: y}]; ok {
							fmt.Printf("Collision at %d, %d\n", x+1, y)
							os.Exit(1)
						}
						next[pos{x: x + 1, y: y}] = cart{c.Next, Right}
					}
				case '/':
					switch c.Direction {
					case Left:
						if _, ok = next[pos{x: x, y: y + 1}]; ok {
							fmt.Printf("Collision at %d, %d\n", x, y+1)
							os.Exit(1)
						}
						next[pos{x: x, y: y + 1}] = cart{c.Next, Down}
					case Right:
						if _, ok = next[pos{x: x, y: y - 1}]; ok {
							fmt.Printf("Collision at %d, %d\n", x, y-1)
							os.Exit(1)
						}
						next[pos{x: x, y: y - 1}] = cart{c.Next, Up}
					case Up:
						if _, ok = next[pos{x: x + 1, y: y}]; ok {
							fmt.Printf("Collision at %d, %d\n", x+1, y)
							os.Exit(1)
						}
						next[pos{x: x + 1, y: y}] = cart{c.Next, Right}
					case Down:
						if _, ok = next[pos{x: x - 1, y: y}]; ok {
							fmt.Printf("Collision at %d, %d\n", x-1, y)
							os.Exit(1)
						}
						next[pos{x: x - 1, y: y}] = cart{c.Next, Left}
					}
				case '+':
					switch c.Next {
					case LeftTurn:
						switch c.Direction {
						case Up:
							if _, ok = next[pos{x: x - 1, y: y}]; ok {
								fmt.Printf("Collision at %d, %d\n", x-1, y)
								os.Exit(1)
							}
							next[pos{x: x - 1, y: y}] = cart{Straight, Left}
						case Down:
							if _, ok = next[pos{x: x + 1, y: y}]; ok {
								fmt.Printf("Collision at %d, %d\n", x+1, y)
								os.Exit(1)
							}
							next[pos{x: x + 1, y: y}] = cart{Straight, Right}
						case Left:
							if _, ok = next[pos{x: x, y: y + 1}]; ok {
								fmt.Printf("Collision at %d, %d\n", x, y+1)
								os.Exit(1)
							}
							next[pos{x: x, y: y + 1}] = cart{Straight, Down}
						case Right:
							if _, ok = next[pos{x: x, y: y - 1}]; ok {
								fmt.Printf("Collision at %d, %d\n", x, y-1)
								os.Exit(1)
							}
							next[pos{x: x, y: y - 1}] = cart{Straight, Up}
						}
					case RightTurn:
						switch c.Direction {
						case Up:
							if _, ok = next[pos{x: x + 1, y: y}]; ok {
								fmt.Printf("Collision at %d, %d\n", x+1, y)
								os.Exit(1)
							}
							next[pos{x: x + 1, y: y}] = cart{LeftTurn, Right}
						case Down:
							if _, ok = next[pos{x: x - 1, y: y}]; ok {
								fmt.Printf("Collision at %d, %d\n", x-1, y)
								os.Exit(1)
							}
							next[pos{x: x - 1, y: y}] = cart{LeftTurn, Left}
						case Left:
							if _, ok = next[pos{x: x, y: y - 1}]; ok {
								fmt.Printf("Collision at %d, %d\n", x, y-1)
								os.Exit(1)
							}
							next[pos{x: x, y: y - 1}] = cart{LeftTurn, Up}
						case Right:
							if _, ok = next[pos{x: x, y: y + 1}]; ok {
								fmt.Printf("Collision at %d, %d\n", x, y+1)
								os.Exit(1)
							}
							next[pos{x: x, y: y + 1}] = cart{LeftTurn, Down}
						}
					case Straight:
						switch c.Direction {
						case Up:
							if _, ok = next[pos{x: x, y: y - 1}]; ok {
								fmt.Printf("Collision at %d, %d", x, y-1)
								os.Exit(1)
							}
							next[pos{x: x, y: y - 1}] = cart{RightTurn, Up}
						case Down:
							if _, ok = next[pos{x: x, y: y + 1}]; ok {
								fmt.Printf("Collision at %d, %d", x, y+1)
								os.Exit(1)
							}
							next[pos{x: x, y: y + 1}] = cart{RightTurn, Down}
						case Left:
							if _, ok = next[pos{x: x - 1, y: y}]; ok {
								fmt.Printf("Collision at %d, %d", x-1, y)
								os.Exit(1)
							}
							next[pos{x: x - 1, y: y}] = cart{RightTurn, Left}
						case Right:
							if _, ok = next[pos{x: x + 1, y: y}]; ok {
								fmt.Printf("Collision at %d, %d", x+1, y)
								os.Exit(1)
							}
							next[pos{x: x + 1, y: y}] = cart{RightTurn, Right}
						}
					}

				}
			}
		}
	}
	return next
}
