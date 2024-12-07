package y2024

import (
	_ "embed"
	"math"
	"strconv"
	"strings"

	"github.com/ar4s/aoc/types"
	"github.com/mowshon/iterium"
	"github.com/samber/lo"
)

func NewPuzzle_07() *types.Puzzle {
	type Line struct {
		Target int
		Chunks []int
	}
	day := 7

	var parseLine = func(line string) Line {
		splited := strings.Split(line, ": ")
		target, err := strconv.Atoi(splited[0])
		if err != nil {
			panic(err)
		}
		chunks := lo.Map(strings.Split(splited[1], " "), func(item string, _ int) int {
			i, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			return i
		})
		return Line{
			Target: target,
			Chunks: chunks,
		}
	}
	type Operation func(int, int) int

	var add = func(a, b int) int {
		return a + b
	}
	var mul = func(a, b int) int {
		return a * b
	}
	var concat = func(a, b int) int {
		l := math.Log10(float64(b))
		return a*(int(math.Pow10(int(l)+1))) + b
	}
	var aimToTarget = func(line Line, operations []Operation) int {
		permutationSize := len(line.Chunks) - 1
		permutations := iterium.Product(operations, permutationSize)
		for {
			perm, err := permutations.Next()
			if err != nil {
				return 0
			}
			a := line.Chunks[0]
			for i := 0; i <= len(line.Chunks)-2; i++ {
				b := line.Chunks[i+1]
				op := perm[i]
				a = op(a, b)
				if a > line.Target {
					continue
				}
			}
			if a == line.Target {
				return a
			}
		}
	}

	return &types.Puzzle{
		ExampleA:         Example(day),
		ExampleAExpected: 3749,
		InputA:           Input(day),
		SolutionA: func(lines []string) int {
			parsed := lo.Map(lines, func(line string, _ int) Line {
				return parseLine(line)
			})
			operations := []Operation{}

			sum := 0
			for _, line := range parsed {
				sum += aimToTarget(line, operations)
			}
			return sum
		},

		ExampleB:         Example(day),
		ExampleBExpected: 11387,
		InputB:           Input(day),
		SolutionB: func(lines []string) int {
			parsed := lo.Map(lines, func(line string, _ int) Line {
				return parseLine(line)
			})
			operations := []Operation{add, mul, concat}

			sum := 0
			for _, line := range parsed {
				sum += aimToTarget(line, operations)
			}
			return sum
		},
	}
}
