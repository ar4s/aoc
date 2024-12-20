package y2024

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"

	"github.com/ar4s/aoc/types"
)

type Node struct {
	ID   int
	Size int
}

var isFreeSpace = func(n Node) bool {
	return n.ID == -1
}

type Nodes []Node

func (n Nodes) Len() int {
	return len(n)
}

func (n Node) IsFreeSpace() bool {
	return n.ID == -1
}

func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n Nodes) Checksum() int {
	sum := 0
	for i, m := range n {
		if m.IsFreeSpace() {
			continue
		}
		sum += i * m.ID
	}
	return sum
}

func (n Nodes) Clone() Nodes {
	return lo.Map(n, func(i Node, _ int) Node {
		return i.Clone()
	})
}

func (n Node) Clone() Node {
	return Node{ID: n.ID, Size: n.Size}
}

func findFreeSpace(nodes Nodes, size int, maxIndex int) int {
	start := -1

	for i, n := range nodes {
		if i > maxIndex {
			return -1
		}
		if n.IsFreeSpace() && start == -1 {
			start = i
		}
		if !n.IsFreeSpace() && start != -1 {
			if i-start >= size {
				return start
			} else {
				start = -1
			}
		}
	}
	return -1
}

func findLastFileIndex(nodes Nodes, offset int) int {
	for i := len(nodes) - (1 + offset); i >= 0; i-- {
		if !nodes[i].IsFreeSpace() {
			return i
		}
	}
	return -1
}

type FindBlockResult struct {
	Node  Node
	Start int
}

func findLastBlock(nodes Nodes, offsetFromEnd int) FindBlockResult {
	start := findLastFileIndex(nodes, offsetFromEnd)
	if start == -1 {
		return FindBlockResult{Node: Node{ID: -1, Size: -1}, Start: -1}
	}
	lastNode := nodes[start]
	if lastNode.IsFreeSpace() {
		return FindBlockResult{Node: Node{ID: -1, Size: -1}, Start: -1}
	}
	r := FindBlockResult{Node: lastNode, Start: start - lastNode.Size + 1}
	return r
}

func NewPuzzle_09() *types.Puzzle {
	day := 9

	FreeSpace := Node{
		ID:   -1,
		Size: -1,
	}

	_ = isFreeSpace
	var parseLine = func(line string) []int {
		return lo.Map(strings.Split(line, ""), func(s string, _ int) int {
			r, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			return r
		})
	}

	var generateLayout = func(input []int) Nodes {
		r := []Node{}
		order := 0
		for i, c := range input {
			if i%2 == 0 {
				t := lo.Repeat(c, Node{
					ID:   order,
					Size: c,
				})
				r = append(r, t...)
				order++
				continue
			}
			t := lo.Repeat(c, FreeSpace)
			r = append(r, t...)
		}
		return r
	}

	var renderLayout = func(nodes Nodes) {
		for i, _ := range nodes {
			fmt.Printf("%3d", len(nodes)-i-1)
		}
		println()
		for i, _ := range nodes {
			fmt.Printf("%3d", i)
		}
		println("")
		for _, n := range nodes {
			if isFreeSpace(n) {
				fmt.Printf("%3s", ".")
				continue
			}
			fmt.Printf("%3d", n.ID)
		}
		println("")
		println("")
	}
	_ = renderLayout

	return &types.Puzzle{
		Example: Example(day),
		Input:   Input(day),

		ExampleAExpected: 1928,
		SolutionA: func(lines []string) int {
			return 0
			// only one line
			// parsed := parseLine(lines[0])
			// l := generateLayout(parsed)

			// //renderLayout(l)

			// endIndex := len(l) - 1
			// for i := 0; i <= len(l)-1 || i == endIndex; i++ {
			// 	if endIndex == 0 || i == endIndex {
			// 		break
			// 	}
			// 	if !l[i].IsFreeSpace() {
			// 		continue
			// 	}
			// 	for j := endIndex; j >= 0; j-- {
			// 		if !l[j].IsFreeSpace() {
			// 			endIndex = j
			// 			break
			// 		}
			// 	}
			// 	l.Swap(i, endIndex)
			// 	endIndex--
			// }

			// return l.Checksum()
		},

		ExampleBExpected: 8193,
		SolutionB: func(lines []string) int {
			fmt.Println("Part 2")
			parsed := parseLine(lines[0])
			memory := generateLayout(parsed)
			maxIndex := len(memory)
			fmt.Printf("max index: %d\n", maxIndex)
			lastFileIndexFromEnd := 0
			wasMoved := []int{}
			for {
				var last FindBlockResult
				for {
					last = findLastBlock(memory, lastFileIndexFromEnd)
					lastFileIndexFromEnd = maxIndex - last.Start
					if !slices.Contains(wasMoved, last.Node.ID) {
						break
					}
				}
				if last.Node.IsFreeSpace() {
					break
				}
				wasMoved = append(wasMoved, last.Node.ID)

				startFreeSpace := findFreeSpace(memory, last.Node.Size, last.Start)
				if startFreeSpace == -1 {
					continue
				}

				for offset := 0; offset <= last.Node.Size-1; offset++ {
					memory.Swap(last.Start+offset, startFreeSpace+offset)
				}
			}
			return memory.Checksum()
		},
	}
}
