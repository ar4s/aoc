package y2024

import (
	"fmt"
	"math"
	"sort"

	"github.com/ar4s/aoc/types"
	"github.com/samber/lo"
)

func NewPuzzle_01() *types.Puzzle {
	day := 1

	var parseLine = func(line string, left, right *[]int) {
		var l, r int
		fmt.Sscanf(line, "%d %d", &l, &r)
		*left = append(*left, l)
		*right = append(*right, r)
	}

	return &types.Puzzle{
		Example:          Example(day),
		Input:            Input(day),
		ExampleAExpected: 11,
		SolutionA: func(lines []string) int {
			var left []int
			var right []int
			lo.ForEach(lines, func(line string, _ int) {
				parseLine(line, &left, &right)
			})
			sort.Ints(left)
			sort.Ints(right)

			return lo.Reduce(left, func(acc, item int, index int) int {
				return acc + int(math.Abs(float64((item - right[index]))))
			}, 0)

		},

		ExampleBExpected: 31,
		SolutionB: func(lines []string) int {
			var left []int
			var right []int
			lo.ForEach(lines, func(line string, _ int) {
				parseLine(line, &left, &right)
			})
			groups := lo.CountValues(right)

			return lo.Reduce(left, func(acc, item int, _ int) int {
				return acc + item*groups[item]
			}, 0)
		},
	}
}
