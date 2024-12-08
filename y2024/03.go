package y2024

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ar4s/aoc/types"
	"github.com/samber/lo"
)

func NewPuzzle_03() *types.Puzzle {
	day := 3
	var parseLine = func(line string) int {

		r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

		a := r.FindAllStringSubmatch(line, -1)
		sum := 0
		for _, v := range a {
			b, _ := strconv.Atoi(v[1])
			c, _ := strconv.Atoi(v[2])
			sum += b * c
		}
		return sum
	}

	var parseLineB = func(line string) int {
		r := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
		res := r.FindAllStringSubmatch(line, -1)
		fmt.Printf("res: %v\n", res)
		do := true
		sum := 0
		for _, v := range res {
			if strings.HasPrefix(v[0], "do") {
				do = true
			}
			if strings.HasPrefix(v[0], "don't") {
				do = false
			}
			if strings.HasPrefix(v[0], "mul") {
				if do {
					b, _ := strconv.Atoi(v[1])
					c, _ := strconv.Atoi(v[2])
					fmt.Printf("b: %d, c: %d %d\n", b, c, b*c)
					sum += b * c
				}
			}
		}
		fmt.Printf("sum: %d\n", sum)
		return sum
	}

	return &types.Puzzle{
		Example:          Example(day),
		ExampleAExpected: 161,
		Input:            Input(day),
		SolutionA: func(lines []string) int {
			parsed := lo.Map(lines, func(line string, _ int) int {
				return parseLine(line)
			})
			sum := lo.Sum(parsed)
			return sum
		},

		ExampleBExpected: 48,
		SolutionB: func(lines []string) int {
			memory := strings.Join(lines, "")
			return parseLineB(memory)
		},
	}
}
