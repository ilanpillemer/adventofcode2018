package main

import (
	"fmt"
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
		{"^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$", 23},
	}

	for _, test := range tests {
		root := node{}
		walk(test.input, &root)
		fmt.Println(root)
		parse(root.children[0], pos{})
	}
}
