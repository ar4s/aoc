package y2024

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/ar4s/aoc/types"
	"github.com/samber/lo"
)

type Cord struct {
	X int
	Y int
}

var DIR_UP Cord = Cord{X: 0, Y: -1}
var DIR_DOWN Cord = Cord{X: 0, Y: 1}
var DIR_LEFT Cord = Cord{X: -1, Y: 0}
var DIR_RIGHT Cord = Cord{X: 1, Y: 0}

func NewPuzzle_06() *types.Puzzle {
	day := 6

	var findStart = func(lines []string) *Cord {
		for y, line := range lines {
			x := strings.Index(line, "^")
			if x != -1 {
				return &Cord{X: x, Y: y}
			}
		}
		return nil
	}

	var getNextPos = func(pos, dir Cord) Cord {
		return Cord{X: pos.X + dir.X, Y: pos.Y + dir.Y}
	}
	var getLine = func(currPos, currDir Cord, lines []string) (string, error) {
		if currPos.Y >= len(lines) {
			return "", fmt.Errorf("MAX")
		}
		return lines[currPos.Y], nil
	}
	var isObstacle = func(line string, pos Cord) bool {
		return line[pos.X] == '#'
	}
	var rotate = func(dir Cord) Cord {
		switch dir {
		case DIR_UP:
			return DIR_RIGHT
		case DIR_RIGHT:
			return DIR_DOWN
		case DIR_DOWN:
			return DIR_LEFT
		case DIR_LEFT:
			return DIR_UP
		}
		panic("Is not exeustive")
	}

	return &types.Puzzle{
		ExampleA:         Example(day),
		ExampleAExpected: 41,
		InputA:           Input(day),
		SolutionA: func(lines []string) int {
			pos := findStart(lines)
			maxIndex := len(lines[0])
			_ = pos
			if pos == nil {
				panic("Start position not found")
			}
			dir := DIR_UP
			visited := []Cord{*pos}
			for {
				nextPos := getNextPos(*pos, dir)
				line, err := getLine(nextPos, dir, lines)
				if err != nil {
					break
				}
				if nextPos.X == maxIndex || nextPos.Y == maxIndex {
					break
				}
				if isObstacle(line, nextPos) {
					dir = rotate(dir)
					continue
				}
				pos = &nextPos
				visited = append(visited, *pos)
			}
			return len(lo.Uniq(visited))
		},

		ExampleB:         Example(day),
		ExampleBExpected: 123,
		InputB:           Input(day),
		SolutionB: func(lines []string) int {

			return -1
		},
	}
}
