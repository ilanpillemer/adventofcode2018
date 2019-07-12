package main

import "testing"

func Test(t *testing.T) {
	want := 3
	got := count([4]int{3, 2, 1, 1}, [4]int{3, 2, 2, 1}, [4]int{9, 2, 1, 2})
	if want != got {
		t.Errorf("want %d got %d", want, got)
	}

}

// Before: [3, 2, 1, 1]
// 9 2 1 2
// After:  [3, 2, 2, 1]

// Opcode 9 could be mulr: register 2 (which has a value of 1) times register 1 (which has a value of 2) produces 2, which matches the value stored in the output register, register 2.
// Opcode 9 could be addi: register 2 (which has a value of 1) plus value 1 produces 2, which matches the value stored in the output register, register 2.
// Opcode 9 could be seti: value 2 matches the value stored in the output register, register 2; the number given for B is irrelevant.
