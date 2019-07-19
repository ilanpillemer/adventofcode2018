package main

import (
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
		walk(test.input, pos{})
		explore(pos{}, map[pos]bool{}, 0)
		got := longest()
		if got != test.want {
			t.Errorf("%s\n want %d got %d", test.input, test.want, got)
		}

	}
}
