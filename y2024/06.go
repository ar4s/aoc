package y2024

import (
	_ "embed"
	"errors"
	"fmt"
	"slices"

	"github.com/ar4s/aoc/types"
	"github.com/samber/lo"
)

func printPath(path []types.Cord2D, start types.Cord2D, maxSize int, extra []types.Cord2D) {
	fmt.Printf(" ")
	for y := 0; y <= maxSize-1; y++ {
		fmt.Printf("%d", y)
	}
	fmt.Println("")
	for y := 0; y <= maxSize-1; y++ {
	a:
		for x := 0; x <= maxSize-1; x++ {
			if x == 0 {
				fmt.Printf("%d", y)
			}
			if len(extra) >= 0 {
				for _, o := range extra {
					if o.X == x && o.Y == y {
						fmt.Printf("x")
						continue a
					}
				}
			}
			if start.X == x && start.Y == y {
				fmt.Printf("^")
				continue
			}
			if slices.Contains(path, types.Cord2D{X: x, Y: y}) {
				fmt.Printf("*")
				continue
			}
			fmt.Printf(".")
		}
		fmt.Println("")
	}
	println()
}

var errLoop = fmt.Errorf("loop detected")

type MapItem rune

const (
	OBSTACLE MapItem = '#'
	START            = '^'
	EMPTY            = 0
)

type Map map[types.Cord2D]MapItem

func (i MapItem) IsObstacle() bool {
	return i == OBSTACLE
}

func (i MapItem) IsEmpty() bool {
	return i == EMPTY
}

func parseInput(lines []string) (Map, types.Cord2D) {
	var m Map = Map{}
	var start types.Cord2D
	for y, line := range lines {
		for x, c := range line {
			switch MapItem(c) {
			case OBSTACLE:
				m[types.Cord2D{
					X: x, Y: y,
				}] = OBSTACLE
			case START:
				start = types.Cord2D{
					X: x, Y: y,
				}
			}
		}
	}
	return m, start
}

func getGuardPath(start types.Cord2D, maxIndex int, obstaclePredicate func(types.Cord2D) bool) ([]types.Cord2D, error) {
	dir := types.DIR_UP
	pos := start
	visited := []types.Cord2D{start}
	i := 0
	for {
		nextItemCord := pos.ApplyDirection(dir)
		if i > maxIndex*maxIndex {
			return visited, errLoop
		}
		if nextItemCord.IsOutOfBound(maxIndex) {
			break
		}

		if obstaclePredicate(nextItemCord) {
			dir = dir.RotateCW()
			continue
		}
		pos = nextItemCord
		visited = append(visited, pos)
		i++
	}
	return visited, nil
}

type MapPredicate func(types.Cord2D) bool

func defaultObstaclePredicateFactory(m Map) MapPredicate {
	return func(c types.Cord2D) bool {
		return m[c].IsObstacle()
	}
}

func NewPuzzle_06() *types.Puzzle {
	day := 6

	return &types.Puzzle{
		Example:          Example(day),
		ExampleAExpected: 41,
		Input:            Input(day),
		SolutionA: func(lines []string) int {
			m, start := parseInput(lines)
			maxIndex := len(lines[0])
			visited, err := getGuardPath(start, maxIndex, defaultObstaclePredicateFactory(m))
			if err != nil {
				panic("Oh no! This part should not contain a loop!")
			}
			return len(lo.Uniq(visited))
		},

		ExampleBExpected: 6,
		SolutionB: func(lines []string) int {
			m, start := parseInput(lines)
			maxIndex := len(lines[0])
			defaultPredicate := defaultObstaclePredicateFactory(m)
			visited, err := getGuardPath(start, maxIndex, defaultPredicate)
			if err != nil {
				panic("Oh no! This part should not contain a loop!")
			}
			loops := 0
			for _, v := range lo.Uniq(visited) {
				var obstaclePredicate = func(c types.Cord2D) bool {
					return defaultPredicate(c) || v == c
				}
				_, err := getGuardPath(start, maxIndex, obstaclePredicate)
				if errors.Is(err, errLoop) {
					loops++
				}
			}
			return loops
		},
	}
}
