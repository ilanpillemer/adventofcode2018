package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	sum := sumMetadata(line)
	fmt.Println("sum", sum)
}

type node struct {
	children []node
	metadata []int
}

func (n *node) sum() int {
	sum := 0
	for _, v := range n.metadata {
		sum += v
	}
	for _, child := range n.children {
		sum += child.sum()
	}
	return sum
}

func sumMetadata(lic string) int {
	lic = strings.TrimSpace(lic)
	root, _ := build(strings.Fields(lic), 0)
	fmt.Println(lic, root)
	return root.sum()
}

func build(fields []string, pointer int) (node, int) {
	n := node{
		children: make([]node, 0),
		metadata: make([]int, 0),
	}
	childCount, metaDataCount := mustAtoi(fields[pointer]), mustAtoi(fields[pointer+1])
	pointer = pointer + 2
	log.Printf("node has %d children\n", childCount)
	for i := 0; i < childCount; i++ {
		child, shiftedPointer := build(fields, pointer)
		pointer = shiftedPointer
		n.children = append(n.children, child)
	}

	for i := 0; i < metaDataCount; i++ {
		n.metadata = append(n.metadata, mustAtoi(fields[pointer]))
		pointer++
	}
	return n, pointer
}

func mustAtoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err.Error)
	}
	return i
}
