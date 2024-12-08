package y2024

import (
	_ "embed"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/ar4s/aoc/types"
	"github.com/samber/lo"
)

func NewPuzzle_05() *types.Puzzle {
	day := 5

	type Rules = map[int][]int

	var parseRule = func(line string, res Rules) {
		var start, end int
		_, err := fmt.Sscanf(line, "%d|%d", &start, &end)
		if err != nil {
			panic(err)
		}
		res[start] = append(res[start], end)
	}

	var isInRules = func(start, end int, rules Rules) bool {
		rule := rules[start]
		return slices.Contains(rule, end)
	}
	var parseLine = func(line string, rules Rules, shouldFix bool) int {
		chunks := lo.Map(strings.Split(line, ","), func(item string, _ int) int {
			a, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			return a
		})

		outerOk := true

		for index, chunk := range chunks {
			for _, sub := range chunks[index+1:] {
				ok := isInRules(chunk, sub, rules)
				if !ok && !shouldFix {
					return 0
				}
				if !ok && shouldFix {
					sort.Slice(chunks, func(i, j int) bool {
						a := rules[chunks[i]]
						fmt.Println("a", a)
						fmt.Println("j", chunks[j])
						return slices.Contains(a, chunks[j])
					})
					return chunks[len(chunks)/2]
				}
			}
		}
		if outerOk && shouldFix {
			return 0
		}
		return chunks[len(chunks)/2]
	}
	_ = parseLine

	return &types.Puzzle{
		Example:          Example(day),
		ExampleAExpected: 143,
		Input:            Input(day),
		SolutionA: func(lines []string) int {
			rules := make(map[int][]int)
			lo.ForEach(lines, func(line string, _ int) {
				if !strings.Contains(line, "|") {
					return
				}
				parseRule(line, rules)
			})
			sum := 0
			lo.ForEach(lines, func(line string, _ int) {
				if strings.Contains(line, ",") {
					sum += parseLine(line, rules, false)
				}
			})
			return sum
		},

		ExampleBExpected: 123,
		SolutionB: func(lines []string) int {
			rules := make(map[int][]int)
			fmt.Println("______")
			lo.ForEach(lines, func(line string, _ int) {
				if !strings.Contains(line, "|") {
					return
				}
				parseRule(line, rules)
			})
			sum := 0
			lo.ForEach(lines, func(line string, _ int) {
				if strings.Contains(line, ",") {
					sum += parseLine(line, rules, true)
				}
			})
			return sum
		},
	}
}
