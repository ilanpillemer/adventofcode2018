package main

import (
	"strings"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"^WNE$", 3},
		{"^ENWWW(NEEE|SSE(EE|N))$", 10},
		{"^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$", 18},
		{"^ENNWSWW(NEWS)SSSEEN(WNSE)EE(SWEN)NNN$", 18},
		{"^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$", 23},
		{"^ESSWWN(E|NNENN(EESS(WNSE)SSS|WWWSSSSE(SW|NNNE)))$", 23},
	}

	for _, test := range tests {
		initMaze()

		test.input = strings.Replace(test.input, "|)", ")", -1)
		walk(test.input, pos{})
		//display(10, 10)
		explore(pos{}, map[pos]bool{}, 0)
		//fmt.Println("longest", longest())
		got := longest()
		if got != test.want {
			t.Errorf("%s\n want %d got %d", test.input, test.want, got)
		}

	}
}
