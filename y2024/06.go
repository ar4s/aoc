package y2024

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/ar4s/aoc/types"
	"github.com/samber/lo"
)

func NewPuzzle_06() *types.Puzzle {
	day := 6

	var findStart = func(lines []string) *types.Cord2D {
		for y, line := range lines {
			x := strings.Index(line, "^")
			if x != -1 {
				return &types.Cord2D{X: x, Y: y}
			}
		}
		return nil
	}

	var getNextPos = func(pos, dir types.Cord2D) types.Cord2D {
		return types.Cord2D{X: pos.X + dir.X, Y: pos.Y + dir.Y}
	}
	var getLine = func(currPos, currDir types.Cord2D, lines []string) (string, error) {
		if currPos.Y >= len(lines) {
			return "", fmt.Errorf("MAX")
		}
		return lines[currPos.Y], nil
	}
	var isObstacle = func(line string, pos types.Cord2D) bool {
		return line[pos.X] == '#'
	}
	var rotate = func(dir types.Cord2D) types.Cord2D {
		switch dir {
		case types.DIR_UP:
			return types.DIR_RIGHT
		case types.DIR_RIGHT:
			return types.DIR_DOWN
		case types.DIR_DOWN:
			return types.DIR_LEFT
		case types.DIR_LEFT:
			return types.DIR_UP
		}
		panic("Is not exeustive")
	}

	return &types.Puzzle{
		Example:          Example(day),
		ExampleAExpected: 41,
		Input:            Input(day),
		SolutionA: func(lines []string) int {
			pos := findStart(lines)
			maxIndex := len(lines[0])
			_ = pos
			if pos == nil {
				panic("Start position not found")
			}
			dir := types.DIR_UP
			visited := []types.Cord2D{*pos}
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

		ExampleBExpected: 123,
		SolutionB: func(lines []string) int {
			pos := findStart(lines)
			maxIndex := len(lines[0])
			_ = pos
			if pos == nil {
				panic("Start position not found")
			}
			dir := types.DIR_UP
			visited := []types.Cord2D{*pos}
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
	}
}
