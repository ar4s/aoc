package y2024

import (
	_ "embed"
	"fmt"
	"math"
	"strings"

	"github.com/ar4s/aoc/types"
	"github.com/mowshon/iterium"
	"github.com/samber/lo"
)

func NewPuzzle_08() *types.Puzzle {
	day := 8
	type Freq string
	type Antena struct {
		X    int
		Y    int
		Freq Freq
	}

	type Antynode struct {
		X int
		Y int
	}

	var newEmptyAntyNode = func() Antynode {
		return Antynode{-1, -1}
	}

	var parseLine = func(line string, y int) []Antena {
		return lo.Map(strings.Split(line, ""), func(s string, index int) Antena {
			if s == "." {
				return Antena{}
			}
			return Antena{
				X:    index,
				Y:    y,
				Freq: Freq(s),
			}
		})
	}

	var calculateAntynodes = func(a, b Antena) []Antynode {
		dx, dy := b.X-a.X, b.Y-a.Y
		if dx == 0 || dy == 0 {
			return []Antynode{newEmptyAntyNode()}
		}
		n1 := Antynode{X: a.X - dx, Y: a.Y - dy}
		n2 := Antynode{X: b.X + dx, Y: b.Y + dy}
		return []Antynode{n1, n2}
	}
	var calculateAntynodesPart2 = func(a, b Antena, maxSize int) []Antynode {
		dx, dy := b.X-a.X, b.Y-a.Y
		if dx == 0 || dy == 0 {
			return []Antynode{{X: a.X, Y: a.Y}, {X: b.X, Y: b.Y}}
		}
		res := []Antynode{}
		for i := 0; i <= int(
			math.Max(
				math.Abs(float64(maxSize/dx)),
				math.Abs(float64(maxSize/dy)),
			),
		); i++ {
			res = append(res, Antynode{X: a.X - dx*i, Y: a.Y - dy*i})
			res = append(res, Antynode{X: b.X + dx*i, Y: b.Y + dy*i})
		}
		return res
	}

	var renderMap = func(lines []string, antinodes []Antynode) {

		var antAtLine = func(y int) []Antynode {
			return lo.Filter(antinodes, func(a Antynode, _ int) bool {
				return y == a.Y
			})
		}
		for y, line := range lines {
			as := antAtLine(y)
			o := []rune(line)
			for _, a := range as {
				o[a.X] = '#'
			}
			fmt.Printf("%d\t%s\n", y, string(o))
		}
	}

	return &types.Puzzle{
		Example: Example(day),
		Input:   Input(day),

		ExampleAExpected: 14,
		SolutionA: func(lines []string) int {
			maxSize := len(lines) - 1
			antenas := map[Freq][]Antena{}
			antinodes := map[Freq][]Antynode{}
			for y, line := range lines {
				for _, a := range parseLine(line, y) {
					if a.Freq == "" {
						continue
					}
					antenas[a.Freq] = append(antenas[a.Freq], a)
				}
			}

			for k, v := range antenas {
				a := iterium.Combinations(v, 2)
				for b := range a.Chan() {
					n := calculateAntynodes(b[0], b[1])
					for _, m := range n {
						if m.X <= maxSize && m.Y <= maxSize && m.X >= 0 && m.Y >= 0 {
							antinodes[k] = append(antinodes[k], m)
						}
					}
				}
			}
			allAntiNodes := []Antynode{}
			for _, v := range antinodes {
				allAntiNodes = append(allAntiNodes, v...)
			}

			renderMap(lines, lo.FlatMap(lo.Values(antinodes), func(item []Antynode, _ int) []Antynode {
				return item
			}))
			return len(lo.Uniq(allAntiNodes))
		},

		ExampleBExpected: 34,
		SolutionB: func(lines []string) int {
			maxSize := len(lines) - 1
			antenas := map[Freq][]Antena{}
			antinodes := map[Freq][]Antynode{}
			for y, line := range lines {
				for _, a := range parseLine(line, y) {
					if a.Freq == "" {
						continue
					}
					antenas[a.Freq] = append(antenas[a.Freq], a)
				}
			}

			for k, v := range antenas {
				a := iterium.Combinations(v, 2)
				for b := range a.Chan() {
					n := calculateAntynodesPart2(b[0], b[1], maxSize)
					for _, m := range n {
						if m.X <= maxSize && m.Y <= maxSize && m.X >= 0 && m.Y >= 0 {
							antinodes[k] = append(antinodes[k], m)
						}
					}

				}
			}
			allAntiNodes := []Antynode{}
			for _, v := range antinodes {
				allAntiNodes = append(allAntiNodes, v...)
			}

			renderMap(lines, lo.FlatMap(lo.Values(antinodes), func(item []Antynode, _ int) []Antynode {
				return item
			}))
			return len(lo.Uniq(allAntiNodes))
		},
	}
}
