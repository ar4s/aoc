package y2023

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/ar4s/aoc/types"
	"github.com/samber/lo"
)

//go:embed examples/02_a.txt
var example_2a string

//go:embed examples/02_b.txt
var example_2b string

//go:embed inputs/02_a.txt
var input_2a string

//go:embed inputs/02_b.txt
var input_2b string

type Game struct {
	Number   int
	Possible bool
}

var conditions = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func ParseGame(line string) Game {
	chunks := strings.Split(line, ":")
	if len(chunks) != 2 {
		panic("Invalid input")
	}
	g := Game{Number: 0, Possible: true}
	fmt.Sscanf(chunks[0], "Game %d", &g.Number)

	sets := strings.Split(chunks[1], ";")
	for _, set := range sets {
		c := strings.Split(set, ",")
		c = lo.Map(c, func(item string, index int) string {
			return strings.TrimSpace(item)
		})
		for _, item := range c {
			v := 0
			color := ""
			fmt.Sscanf(item, "%d %s", &v, &color)
			if v > conditions[color] {
				g.Possible = false
			}
		}
	}
	return g
}

func NewPuzzle_02() *types.Puzzle {
	return &types.Puzzle{
		ExampleA:         example_2a,
		InputA:           input_2a,
		ExampleAExpected: 8,
		SolutionA: func(lines []string) int {
			games := lo.Map(lines, func(item string, _ int) Game {
				return ParseGame(item)
			})
			possibleGames := lo.Filter(games, func(item Game, _ int) bool {
				return item.Possible
			})
			return lo.Sum(lo.Map(possibleGames, func(item Game, _ int) int {
				return item.Number
			}))
		},

		ExampleB:         example_2b,
		InputB:           input_2b,
		ExampleBExpected: -1,
		SolutionB: func(lines []string) int {

			return 0
		},
	}
}
