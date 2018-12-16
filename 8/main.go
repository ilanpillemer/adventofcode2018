package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	sum := sumMetadata(line)
	fmt.Println("sum", sum)
	complexSum := sumComplex(line)
	fmt.Println("complexSum", complexSum)
}

type node struct {
	children []node
	metadata []int
}

func (n *node) isLeaf() bool {
	if len(n.children) == 0 {
		return true
	}
	return false
}

func (n *node) isValidIndex(i int) bool {
	return i >= 0 && i <= len(n.children)
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

func (n *node) complexSum() int {
	sum := 0
	if n.isLeaf() {
		for _, v := range n.metadata {
			sum += v
		}
		return sum
	}
	for _, childIndex := range n.metadata {
		if !n.isValidIndex(childIndex) {
			continue
		}
		sum += n.children[childIndex-1].complexSum() // metadata is index 1 based
	}
	return sum
}

func sumComplex(lic string) int {
	lic = strings.TrimSpace(lic)
	root, _ := build(strings.Fields(lic), 0)
	return root.complexSum()
}

func sumMetadata(lic string) int {
	lic = strings.TrimSpace(lic)
	root, _ := build(strings.Fields(lic), 0)
	return root.sum()
}

func build(fields []string, pointer int) (node, int) {
	n := node{
		children: make([]node, 0),
		metadata: make([]int, 0),
	}
	childCount, metaDataCount := mustAtoi(fields[pointer]), mustAtoi(fields[pointer+1])
	pointer = pointer + 2
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
