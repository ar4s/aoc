package y2024

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ar4s/aoc/types"
	"github.com/samber/lo"
)

type Node struct {
	ID    int
	Value int
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
	return Node{ID: n.ID, Value: n.Value}
}

func findFreeSpace(nodes Nodes, size int, lessThan int) int {
	start := -1

	for i, n := range nodes {
		if n.IsFreeSpace() && start == -1 {
			start = i

		}
		if !n.IsFreeSpace() && start != -1 {
			if i-start >= size && start < lessThan {
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

func findLastBlock(nodes Nodes, offset int) FindBlockResult {
	start := findLastFileIndex(nodes, offset)
	lastNode := nodes[start]
	fmt.Println(nodes[start])
	fileId := nodes[start].ID
	for i := start; i >= 0; i-- {
		if nodes[i].IsFreeSpace() {
			continue
		}
		if !nodes[i].IsFreeSpace() && fileId == -1 {
			fileId = nodes[i].ID
		}
		if !nodes[i].IsFreeSpace() && fileId != nodes[i].ID {
			fmt.Println(">>> change", fileId, nodes[i].ID)
			return FindBlockResult{
				Node:  lastNode,
				Start: i + lastNode.Value - 1,
			}
		}
	}
	return FindBlockResult{
		Node:  Node{ID: -1, Value: -1},
		Start: -1,
	}
}

func NewPuzzle_09() *types.Puzzle {
	day := 9

	FreeSpace := Node{
		ID:    -1,
		Value: -1,
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
					ID:    order,
					Value: c,
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
		for _, n := range nodes {
			if isFreeSpace(n) {
				fmt.Print(".")
				continue
			}
			fmt.Print(n.ID)
		}
		println("")
	}
	_ = renderLayout

	return &types.Puzzle{
		Example: Example(day),
		Input:   Input(day),

		ExampleAExpected: 1928,
		SolutionA: func(lines []string) int {
			// only one line
			parsed := parseLine(lines[0])
			l := generateLayout(parsed)

			// renderLayout(l)

			endIndex := len(l) - 1
			for i := 0; i <= len(l)-1 || i == endIndex; i++ {
				if endIndex == 0 || i == endIndex {
					break
				}
				if !l[i].IsFreeSpace() {
					continue
				}
				for j := endIndex; j >= 0; j-- {
					if !l[j].IsFreeSpace() {
						endIndex = j
						break
					}
				}
				l.Swap(i, endIndex)
				endIndex--
			}
			// renderLayout(l)

			return l.Checksum()
		},

		ExampleBExpected: 2858,
		SolutionB: func(lines []string) int {
			fmt.Println("Part 2")
			parsed := parseLine(lines[0])
			l := generateLayout(parsed)

			i := len(l) - 1

			for {
				if i == 0 {
					break
				}
				i--
			}

			// renderLayout(l)
			// var freeSpaceAvailable = func(i int, nodes Nodes) int {
			// 	for x, y := range nodes[i:] {
			// 		if !y.IsFreeSpace() {
			// 			return x
			// 		}
			// 	}
			// 	return 0
			// }
			// endIndex := len(l) - 1
			// i := 0
			//			for i := 0; i < len(l); i++ {
			// 	if l[i].IsFreeSpace() {
			// 		j := len(l) - 1
			// 		for l[i].Value > 0 && j > i {
			// 			if !l[j].IsFreeSpace() && l[j].Value <= l[i].Value {
			// 				newLayout = append(newLayout, l[j])
			// 				l[i].Value -= l[j].Value
			// 				l[j].ID = -1
			// 				j = len(l) - 1
			// 			}
			// 			j--
			// 		}

			// 		// If gap still exists, reflect it in reduced
			// 		if l[i].IsFreeSpace() {
			// 			newLayout = append(newLayout, l[i])
			// 		}
			// 	} else if l[i].Value != 0 {
			// 	}
			// }
			// for {
			// 	if i >= len(l)-1 {
			// 		break
			// 	}

			// 	if endIndex == 0 || i == endIndex {
			// 		break
			// 	}
			// 	if !l[i].IsFreeSpace() {
			// 		i++
			// 		continue
			// 	}
			// 	freeSpaceSize := freeSpaceAvailable(i, l)

			// 	fmt.Println("freeSpace", freeSpaceSize)
			// 	for j := endIndex; j >= 0; j-- {
			// 		if !l[j].IsFreeSpace() && freeSpaceSize >= l[j].Value {
			// 			org := l[j].Clone()
			// 			fmt.Println("!!!", j, freeSpaceSize, org.Value, org.ID)
			// 			for x := 0; x <= org.Value-1; x++ {
			// 				fmt.Println("???")
			// 				l.Swap(i, j-x)
			// 				i++
			// 				endIndex--
			// 			}
			// 			break
			// 		}
			// 	}
			// 	renderLayout(l)
			// 	// l.Swap(i, endIndex)
			// 	// endIndex--
			// 	//todo
			// 	if i >= 1 {
			// 		break
			// 	}
			// 	i++
			// }
			i = 0
			// orig := l.Clone()
			maxIndex := len(l)
			lastFileIndexFromEnd := 0
			wasMoved := []int{}
			for {
				if i == 5 {
					break
				}

				fmt.Println()
				renderLayout(l)
				last := findLastBlock(l, lastFileIndexFromEnd)
				if slices.Contains(wasMoved, last.Node.ID) {
					i++
					continue
				}
				startFreeSpace := findFreeSpace(l, last.Node.Value, last.Start)
				fmt.Printf(">!!!!!!!!! last %+v %+v %+v\n", last.Node, last.Start, startFreeSpace)

				if startFreeSpace == -1 {
					i++
					lastFileIndexFromEnd += maxIndex - last.Start
					continue
				}
				if last.Node.IsFreeSpace() && startFreeSpace == -1 {
					break
				}
				fmt.Printf("last %+v %+v\n", last, startFreeSpace)
				for offset := 0; offset <= last.Node.Value-1; offset++ {
					l.Swap(last.Start+offset, startFreeSpace+offset)
					lastFileIndexFromEnd = maxIndex - last.Start + offset
					fmt.Println("!!!! lastFileIndex", lastFileIndexFromEnd)
				}
				wasMoved = append(wasMoved, last.Node.ID)
				i++

			}

			renderLayout(l)

			return l.Checksum()
		},
	}
}
