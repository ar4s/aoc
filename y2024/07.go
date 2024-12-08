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

type Line struct {
	Target int
	Chunks []int
}

type BinOperation func(int, int) int

func calc(line Line, operations []BinOperation, value int, index int) bool {
	if index == len(line.Chunks) {
		return value == line.Target
	}
	if value > line.Target {
		return false
	}
	right := line.Chunks[index]
	for _, op := range operations {
		c := op(value, right)
		if calc(line, operations, c, index+1) {
			return true
		}
	}
	return false
}

func NewPuzzle_07() *types.Puzzle {
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

	var aimToTargetRec = func(line Line, operations []BinOperation) bool {
		r := calc(line, operations, line.Chunks[0], 1)
		return r
	}

	var aimToTarget = func(line Line, operations []BinOperation) int {
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
	_ = aimToTarget

	return &types.Puzzle{
		Example: Example(day),
		Input:   Input(day),

		ExampleAExpected: 3749,
		SolutionA: func(lines []string) int {
			parsed := lo.Map(lines, func(line string, _ int) Line {
				return parseLine(line)
			})
			operations := []BinOperation{add, mul}

			sum := 0
			for _, line := range parsed {
				if aimToTargetRec(line, operations) {
					sum += line.Target
				}
			}
			return sum
		},

		ExampleBExpected: 11387,
		SolutionB: func(lines []string) int {
			parsed := lo.Map(lines, func(line string, _ int) Line {
				return parseLine(line)
			})
			operations := []BinOperation{add, mul, concat}

			sum := 0
			for _, line := range parsed {
				if aimToTargetRec(line, operations) {
					sum += line.Target
				}
			}
			return sum
		},
	}
}
