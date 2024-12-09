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

type NodeV2 struct {
	ID    int
	Value int
	Size  int
}

var isFreeSpace = func(n Node) bool {
	return n.ID == -1
}

type Nodes []Node
type NodesV2 []NodeV2

func (n Nodes) Len() int {
	return len(n)
}

func (n Node) IsFreeSpace() bool {
	return n.ID == -1
}
func (n NodeV2) IsFreeSpace() bool {
	return n.ID == -1
}

func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n NodesV2) Swap(i, j int) {
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

func (n Node) Clone() Node {
	return Node{ID: n.ID, Value: n.Value}
}
func (n NodeV2) Clone() NodeV2 {
	return NodeV2{ID: n.ID, Value: n.Value, Size: n.Size}
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
	var generateLayoutV2 = func(input []int) NodesV2 {
		r := NodesV2{}
		order := 0
		for i, c := range input {
			if i%2 == 0 {
				r = append(r, NodeV2{
					ID:    order,
					Value: c,
					Size:  c,
				})
				order++
				continue
			}
			r = append(r, NodeV2{
				ID:    -1,
				Value: -1,
				Size:  c,
			})
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
	var renderLayoutV2 = func(nodes NodesV2) {
		for _, n := range nodes {
			if n.IsFreeSpace() {
				fmt.Print(strings.Repeat(".", n.Size))
				continue
			}
			fmt.Print(strings.Repeat(strconv.Itoa(n.ID), n.Size))

		}
		println("")
	}

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
			l := generateLayoutV2(parsed)
			renderLayoutV2(l)

			endIndex := len(l) - 1
			_ = endIndex
			i := 0
		outer:
			for {
				if i >= len(l) {
					break
				}
				if !l[i].IsFreeSpace() {
					i++
					continue
				}

				for j := endIndex; j >= 0; j-- {
					if !l[j].IsFreeSpace() && l[i].Size >= l[j].Size {
						endIndex = j
						v := l[j].Clone()
						x := l[i].Clone()
						diff := x.Size - v.Size
						fmt.Println(v, x, diff)
						l[i] = v
						l[j] = x

						if diff != 0 {
							l = slices.Insert(l, i+1, NodeV2{
								ID:    -1,
								Value: -1,
								Size:  x.Size - v.Size,
							})
							i += 2
						}
						renderLayoutV2(l)
						continue outer
					}
				}
				i++
			}

			renderLayoutV2(l)
			// tmp := []int{1, 2, 3, 4, 5}
			// fmt.Println(slices.Insert(tmp, 1, 0))
			return 2858
		},
	}
}
