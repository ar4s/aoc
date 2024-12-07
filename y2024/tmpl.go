package y2024

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/ar4s/aoc/types"
	"github.com/samber/lo"
)

func NewPuzzle_0x() *types.Puzzle {
	day := 3
	var parseLine = func(line string) []int {
		return lo.Map(strings.Split(line, " "), func(s string, _ int) int {
			r, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			return r
		})
	}

	return &types.Puzzle{
		ExampleA:         Example(day),
		ExampleAExpected: 0,
		InputA:           Input(day),
		SolutionA: func(lines []string) int {
			parsed := lo.Map(lines, func(line string, _ int) []int {
				return parseLine(line)
			})
			_ = parsed
			return -1
		},

		ExampleB:         Example(day),
		ExampleBExpected: 0,
		InputB:           Input(day),
		SolutionB: func(lines []string) int {
			return -1
		},
	}
}
