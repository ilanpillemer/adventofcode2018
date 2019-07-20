package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var part2 = flag.Bool("part2", false, "part2")
var regs = [6]int{0, 0, 0, 0, 0, 0} // 6 registers
var ops = make(map[string]func([3]int))
var ip int //instruction pointer register

type instruction struct {
	op string
	A  int
	B  int
	C  int
}

var prog = make(map[int]instruction)

const (
	A int = iota
	B
	C
)

func init() {
	ops["addr"] = addr
	ops["addi"] = addi
	ops["mulr"] = mulr
	ops["muli"] = muli
	ops["banr"] = banr
	ops["bani"] = bani
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
}
func main() {
	flag.Parse()
	if *part2 {
		regs = [6]int{0, 0, 0, 0, 0, 0}
	}

	in := bufio.NewScanner(os.Stdin)
	instLine := 0
	for in.Scan() {
		line := in.Text()
		if strings.HasPrefix(line, "#") { //set ip
			p := strings.Fields(line)[1]
			ip, _ = strconv.Atoi(p)
			continue
		}
		//load
		parts := strings.Fields(line)
		op, strA, strB, strC := parts[0], parts[1], parts[2], parts[3]
		a, _ := strconv.Atoi(strA)
		b, _ := strconv.Atoi(strB)
		c, _ := strconv.Atoi(strC)
		prog[instLine] = instruction{
			op: op,
			A:  a,
			B:  b,
			C:  c,
		}
		instLine++
	}

	display()
	fmt.Println("executing")
	for execute() {
		//display executing registers
		fmt.Println(regs) //
	}
}

func display() {
	for i := 0; i < len(prog); i++ {
		fmt.Printf("%d: %v\n", i, prog[i])
	}
	fmt.Println("ip:", ip)
	fmt.Println("regs", regs)

}

func execute() bool {
	if regs[ip] >= len(prog) {
		return false //halt
	}
	inst := prog[regs[ip]]
	ops[inst.op]([3]int{inst.A, inst.B, inst.C})

	if regs[ip] == 28 {
		fmt.Println(regs[ip], inst)
		fmt.Println(regs)
		fmt.Println("-------------")
		os.Exit(0)
	}

	regs[ip]++

	return true
}

// addr (add register) stores into register C the result of adding register A and register B.
func addr(in [3]int) {
	regs[in[C]] = regs[in[A]] + regs[in[B]]
}

// addi (add immediate) stores into register C the result of adding register A and value B.
func addi(in [3]int) {
	regs[in[C]] = regs[in[A]] + in[B]
}

// mulr (multiply register) stores into register C the result of multiplying register A and register B
func mulr(in [3]int) {
	regs[in[C]] = regs[in[A]] * regs[in[B]]
}

// muli (multiply immediate) stores into register C the result of multiplying register A and value B
func muli(in [3]int) {
	regs[in[C]] = regs[in[A]] * in[B]
}

//banr (bitwise AND register) stores into register C the result of the bitwise AND of register A and register B
func banr(in [3]int) {
	regs[in[C]] = regs[in[A]] & regs[in[B]]
}

//bani (bitwise AND immediate) stores into register C the result of the bitwise AND of register A and value B.
func bani(in [3]int) {
	regs[in[C]] = regs[in[A]] & in[B]
}

//borr (bitwise OR register) stores into register C the result of the bitwise OR of register A and register B
func borr(in [3]int) {
	regs[in[C]] = regs[in[A]] | regs[in[B]]
}

// bori (bitwise OR immediate) stores into register C the result of the bitwise OR of register A and value B
func bori(in [3]int) {
	regs[in[C]] = regs[in[A]] | in[B]
}

// setr (set register) copies the contents of register A into register C. (Input B is ignored.)
func setr(in [3]int) {
	regs[in[C]] = regs[in[A]]
}

//seti (set immediate) stores value A into register C. (Input B is ignored.)
func seti(in [3]int) {
	regs[in[C]] = in[A]
}

//gtir (greater-than immediate/register) sets register C to 1 if value A is greater than register B. Otherwise, register C is set to 0
func gtir(in [3]int) {
	if in[A] > regs[in[B]] {
		regs[in[C]] = 1
		return
	}
	regs[in[C]] = 0
}

//gtri (greater-than register/immediate) sets register C to 1 if register A is greater than value B. Otherwise, register C is set to 0.
func gtri(in [3]int) {
	if regs[in[A]] > in[B] {
		regs[in[C]] = 1
		return
	}
	regs[in[C]] = 0
}

//gtrr (greater-than register/register) sets register C to 1 if register A is greater than register B. Otherwise, register C is set to 0
func gtrr(in [3]int) {
	if regs[in[A]] > regs[in[B]] {
		regs[in[C]] = 1
		return
	}
	regs[in[C]] = 0
}

//eqir (equal immediate/register) sets register C to 1 if value A is equal to register B. Otherwise, register C is set to 0
func equir(in [3]int) {
	if in[A] == regs[in[B]] {
		regs[in[C]] = 1
		return
	}
	regs[in[C]] = 0
}

//eqri (equal register/immediate) sets register C to 1 if register A is equal to value B. Otherwise, register C is set to 0.
func eqri(in [3]int) {
	if regs[in[A]] == in[B] {
		regs[in[C]] = 1
		return
	}
	regs[in[C]] = 0
}

//eqrr (equal register/register) sets register C to 1 if register A is equal to register B. Otherwise, register C is set to 0
func eqrr(in [3]int) {
	if regs[in[A]] == regs[in[B]] {
		regs[in[C]] = 1
		return
	}
	regs[in[C]] = 0
}
