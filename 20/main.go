package main

import "fmt"

func main() {

}

type pos struct {
	x, y int
	path string
}

func (p pos) north() pos { return pos{p.x, p.y - 1, p.path + "N"} }
func (p pos) south() pos { return pos{p.x, p.y + 1, p.path + "S"} }
func (p pos) west() pos  { return pos{p.x - 1, p.y, p.path + "W"} }
func (p pos) east() pos  { return pos{p.x + 1, p.y, p.path + "E"} }

func (p pos) change(c rune) pos {
	switch c {
	case 'N':
		return p.north()
	case 'S':
		return p.south()
	case 'W':
		return p.west()
	case 'E':
		return p.east()
	default:
		panic("unreachable")
	}
}

type kind int

const (
	DNODE kind = iota
	ENODE
)

//rooms are nodes
type node struct {
	children []node
	value    string
}

var seen = make(map[pos]bool)

func walk(s string, parent *node) string {

	if len(s) == 0 {
		return ""
	}

	for {
		fmt.Println(s)
		switch h(s) {
		case '(', '^':
			//start of a child node
			cnode := node{}
			s = walk(t(s), &cnode)
			parent.children = append(parent.children, node(cnode))

		case 'N', 'W', 'S', 'E':
			//start of dnode
			var d string
			s, d = processDnode(s)
			dn := node{}
			dn.value = d
			parent.children = append(parent.children, node(dn))
		case '|':
			if h(t(s)) == ')' {
				//empty node
				en := node{}
				parent.children = append(parent.children, node(en))
			}
			s = t(s)
			//next child node
		case ')':
			return t(s)
			// end of a node
		case '$':
			return ""
		}
	}

	panic("unreachable code")
}

func h(s string) byte {
	if len(s) == 0 {
		return '$'
	}
	return s[0]
}

func t(s string) string {
	return s[1:]
}

func processDnode(s string) (string, string) {
	var d []byte
	for {
		switch h(s) {
		case 'N', 'E', 'W', 'S':
			d = append(d, h(s))
			s = t(s)
		default:
			return s, string(d)
		}
	}
}

//stepcounter
func parse(n node, p pos) {
	queue := n.children
	paths := []string{}
	for _, child := range queue {
		possible := ""
		for _, c := range child.value {
			np := p.change(c)
			possible = possible + np.path
		}
		paths = append(paths, possible)

	}

	fmt.Println("# paths", len(n.children))

	for i, p := range paths {
		fmt.Println(i, p)
	}
}

// ^ENWWW(NEEE|SSE(EE|N))$
// (ENWWW(NEEE|SSE(EE|N)))

///         NODE
//  NODE NODE NODE NODE
