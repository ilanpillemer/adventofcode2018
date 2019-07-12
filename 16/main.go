package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var regs [4]int
var ops = make(map[string]func([4]int))
var opcodesfilter [16]map[string]func([4]int)
var opcode [16]func([4]int)

const (
	A int = iota + 1
	B
	C
)

func setregs(new [4]int) {
	for i := range regs {
		regs[i] = new[i]
	}
}

func isregs(cmp [4]int) bool {
	for i := range regs {
		if regs[i] != cmp[i] {
			return false
		}
	}
	return true
}

func init() {
	ops["addr"] = addr
	ops["addi"] = addi
	ops["mulr"] = mulr
	ops["muli"] = muli
	ops["banr"] = banr
	ops["bain"] = bain
	ops["borr"] = borr
	ops["bori"] = bori
	ops["setr"] = setr
	ops["seti"] = seti
	ops["gtir"] = gtir
	ops["gtri"] = gtri
	ops["gtrr"] = gtrr
	ops["eqir"] = equir
	ops["eqri"] = eqri
	ops["eqrr"] = eqrr

	for i := 0; i < 16; i++ {
		opcodesfilter[i] = make(map[string]func([4]int))
		opcodesfilter[i]["addr"] = addr
		opcodesfilter[i]["addi"] = addi
		opcodesfilter[i]["mulr"] = mulr
		opcodesfilter[i]["muli"] = muli
		opcodesfilter[i]["banr"] = banr
		opcodesfilter[i]["bain"] = bain
		opcodesfilter[i]["borr"] = borr
		opcodesfilter[i]["bori"] = bori
		opcodesfilter[i]["setr"] = setr
		opcodesfilter[i]["seti"] = seti
		opcodesfilter[i]["gtir"] = gtir
		opcodesfilter[i]["gtri"] = gtri
		opcodesfilter[i]["gtrr"] = gtrr
		opcodesfilter[i]["eqir"] = equir
		opcodesfilter[i]["eqri"] = eqri
		opcodesfilter[i]["eqrr"] = eqrr
	}
}

func loadMachine() {
	for i := 0; i < 16; i++ {
		for _, v := range opcodesfilter[i] {
			opcode[i] = v
		}
	}
}

func filtered() bool {
	for _, v := range opcodesfilter {
		if len(v) > 1 {
			return false
		}
	}
	return true
}

func reduce(i int, poss []string) {
	//fmt.Println("testing code", i)
	op := opcodesfilter[i]
	for k := range op {
		good := false
		for _, w := range poss {
			if w == k {
				good = true
			}
		}
		if !good {
			delete(op, k)
		}
	}
	opcodesfilter[i] = op
	if len(op) == 1 {
		//fmt.Printf("Found opcode[%d] --> %v\n", i, op)
		for k := range op {
			constrain(i, k)
		}
	}
}

func constrain(i int, code string) {
	for j := 0; j < 16; j++ {
		if i == j {
			continue
		}
		delete(opcodesfilter[j], code)
	}
}

func main() {
	fmt.Println("Counting")
	in := bufio.NewScanner(os.Stdin)
	more3 := 0
	running := false
	for in.Scan() {
		before := [4]int{}
		after := [4]int{}
		inst := [4]int{}
		line := in.Text()
		if strings.HasPrefix(line, "Before: [") {
			line = strings.TrimPrefix(line, "Before: [")
			line = strings.TrimSuffix(line, "]")
			nums := strings.Split(line, ",")
			for i, v := range nums {
				before[i], _ = strconv.Atoi(strings.TrimSpace(v))
			}
			in.Scan()
			line = in.Text()
			nums = strings.Fields(line)
			for i, v := range nums {
				inst[i], _ = strconv.Atoi(v)
			}
			in.Scan()
			line = in.Text()
			line = strings.TrimPrefix(line, "After:  [")
			line = strings.TrimSuffix(line, "]")
			nums = strings.Split(line, ",")
			//fmt.Println(nums)
			for i, v := range nums {
				after[i], _ = strconv.Atoi(strings.TrimSpace(v))
			}
			in.Scan()
			c, poss := count(before, after, inst)
			if c >= 3 {
				more3++
			}
			reduce(inst[0], poss)
			continue
		}
		if !running {
			in.Scan() //one line
			in.Scan() //two line .. program input starts now
			running = true
			setregs([4]int{0, 0, 0, 0})
		}
		if !filtered() {
			panic("machine not loaded..")
		}
		loadMachine()
		line = in.Text()
		nums := strings.Fields(line)
		for i, v := range nums {
			inst[i], _ = strconv.Atoi(v)
		}
		opcode[inst[0]](inst)

		//	fmt.Println(line)
	}
	fmt.Println("3 or more:", more3)
	fmt.Println("state", opcodesfilter)
	for i, v := range regs {
		fmt.Printf("register [%d] -> [%d]\n", i, v)
	}

}

// addr (add register) stores into register C the result of adding register A and register B.
func addr(in [4]int) {
	regs[in[C]] = regs[in[A]] + regs[in[B]]
}

// addi (add immediate) stores into register C the result of adding register A and value B.
func addi(in [4]int) {
	regs[in[C]] = regs[in[A]] + in[B]
}

// mulr (multiply register) stores into register C the result of multiplying register A and register B
func mulr(in [4]int) {
	regs[in[C]] = regs[in[A]] * regs[in[B]]
}

// muli (multiply immediate) stores into register C the result of multiplying register A and value B
func muli(in [4]int) {
	regs[in[C]] = regs[in[A]] * in[B]
}

//banr (bitwise AND register) stores into register C the result of the bitwise AND of register A and register B
func banr(in [4]int) {
	regs[in[C]] = regs[in[A]] & regs[in[B]]
}

//bani (bitwise AND immediate) stores into register C the result of the bitwise AND of register A and value B.
func bain(in [4]int) {
	regs[in[C]] = regs[in[A]] & in[B]
}

//borr (bitwise OR register) stores into register C the result of the bitwise OR of register A and register B
func borr(in [4]int) {
	regs[in[C]] = regs[in[A]] | regs[in[B]]
}

// bori (bitwise OR immediate) stores into register C the result of the bitwise OR of register A and value B
func bori(in [4]int) {
	regs[in[C]] = regs[in[A]] | in[B]
}

// setr (set register) copies the contents of register A into register C. (Input B is ignored.)
func setr(in [4]int) {
	regs[in[C]] = regs[in[A]]
}

//seti (set immediate) stores value A into register C. (Input B is ignored.)
func seti(in [4]int) {
	regs[in[C]] = in[A]
}

//gtir (greater-than immediate/register) sets register C to 1 if value A is greater than register B. Otherwise, register C is set to 0
func gtir(in [4]int) {
	if in[A] > regs[in[B]] {
		regs[in[C]] = 1
		return
	}
	regs[in[C]] = 0
}

//gtri (greater-than register/immediate) sets register C to 1 if register A is greater than value B. Otherwise, register C is set to 0.
func gtri(in [4]int) {
	if regs[in[A]] > in[B] {
		regs[in[C]] = 1
		return
	}
	regs[in[C]] = 0
}

//gtrr (greater-than register/register) sets register C to 1 if register A is greater than register B. Otherwise, register C is set to 0
func gtrr(in [4]int) {
	if regs[in[A]] > regs[in[B]] {
		regs[in[C]] = 1
		return
	}
	regs[in[C]] = 0
}

//eqir (equal immediate/register) sets register C to 1 if value A is equal to register B. Otherwise, register C is set to 0
func equir(in [4]int) {
	if in[A] == regs[in[B]] {
		regs[in[C]] = 1
		return
	}
	regs[in[C]] = 0
}

//eqri (equal register/immediate) sets register C to 1 if register A is equal to value B. Otherwise, register C is set to 0.
func eqri(in [4]int) {
	if regs[in[A]] == in[B] {
		regs[in[C]] = 1
		return
	}
	regs[in[C]] = 0
}

//eqrr (equal register/register) sets register C to 1 if register A is equal to register B. Otherwise, register C is set to 0
func eqrr(in [4]int) {
	if regs[in[A]] == regs[in[B]] {
		regs[in[C]] = 1
		return
	}
	regs[in[C]] = 0
}

func is(f func([4]int), before [4]int, after [4]int, in [4]int) bool {
	setregs(before)
	f(in)
	return isregs(after)
}

func count(before [4]int, after [4]int, in [4]int) (int, []string) {
	c := 0
	poss := make([]string, 0)
	for k, op := range ops {
		if is(op, before, after, in) {
			poss = append(poss, k)
			c++
		}
	}
	return c, poss
}
