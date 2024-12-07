package y2024

import (
	_ "embed"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/ar4s/aoc/types"
	"github.com/samber/lo"
)

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

func NewPuzzle_02() *types.Puzzle {
	day := 2
	var parseLine = func(line string) []int {
		return lo.Map(strings.Split(line, " "), func(s string, _ int) int {
			r, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			return r
		})
	}

	var lineIsSafe = func(toCheck []int) []int {
		prev := 0
		errors := []int{}
		reversed := reverseInts(toCheck)
		rightOrder := sort.IntsAreSorted(toCheck) || sort.IntsAreSorted(reversed)

		if !rightOrder {
			return []int{0}
		}

		for i, num := range toCheck {
			if i == 0 {
				prev = num
				continue
			}
			diff := math.Abs(float64(num - prev))
			if diff > 3 || diff == 0 {
				errors = append(errors, i)

			}
			if prev == num {
				errors = append(errors, i)
			}
			prev = num
		}
		return errors
	}
	return &types.Puzzle{
		ExampleA:         Example(day),
		ExampleAExpected: 2,
		InputA:           Input(day),
		SolutionA: func(lines []string) int {
			ok := 0
			parsed := lo.Map(lines, func(line string, _ int) []int {
				return parseLine(line)
			})

			for _, line := range parsed {
				isLineOk := len(lineIsSafe(line)) == 0
				if isLineOk {
					ok++
				}
			}
			return ok
		},

		ExampleB:         Example(day),
		ExampleBExpected: 4,
		InputB:           Input(day),
		SolutionB: func(lines []string) int {
			ok := 0
			parsed := lo.Map(lines, func(line string, _ int) []int {
				return parseLine(line)
			})

			for _, line := range parsed {
				errors := lineIsSafe(line)
				if len(errors) == 0 {
					ok++
					continue
				}
				lineLen := len(line)
				for i := 0; i < lineLen; i++ {
					newLine := lo.Filter(line, func(_, j int) bool {
						return j != i
					})
					if len(lineIsSafe(newLine)) == 0 {
						ok++
						break
					}
				}
			}
			return ok
		},
	}
}
