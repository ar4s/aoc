package y2024

import (
	_ "embed"
	"fmt"

	"github.com/ar4s/aoc/types"
	"github.com/samber/lo"
)

func countXMASOccurrences(grid [][]rune, word string) int {
	rows := len(grid)
	cols := len(grid[0])
	wordLen := len(word)
	count := 0

	// Define all 8 directions
	directions := [][2]int{
		{0, 1},   // Right
		{0, -1},  // Left
		{1, 0},   // Down
		{-1, 0},  // Up
		{1, 1},   // Down-Right
		{-1, -1}, // Up-Left
		{1, -1},  // Down-Left
		{-1, 1},  // Up-Right
	}

	checkWord := func(r, c, dr, dc int) bool {
		for i := 0; i < wordLen; i++ {
			nr, nc := r+i*dr, c+i*dc
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] != rune(word[i]) {
				return false
			}
		}
		return true
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, d := range directions {
				dr, dc := d[0], d[1]
				if checkWord(r, c, dr, dc) {
					count++
				}
			}
		}
	}
	fmt.Println(count)
	return count
}
func NewPuzzle_04() *types.Puzzle {
	day := 4

	// reg := regexp.MustCompile(`XMAS|SAMX`)
	// var search = func(line string) int {
	// 	return len(reg.FindAllStringIndex(line, -1))
	// }

	return &types.Puzzle{
		ExampleA:         Example(day),
		ExampleAExpected: 18,
		InputA:           Input(day),
		SolutionA: func(lines []string) int {
			dimension := len(lines[0])
			fmt.Println(dimension)
			r := lo.Map(lines, func(line string, _ int) []rune {
				return []rune(line)
			})
			return countXMASOccurrences(r, "XMAS")
		},

		ExampleB:         Example(day),
		ExampleBExpected: 9,
		InputB:           Input(day),
		SolutionB: func(lines []string) int {
			type Cord struct {
				X int
				Y int
			}
			candidates := []Cord{}
			for y, line := range lines {
				if y == 0 || y == len(lines)-1 {
					continue
				}
				for x, char := range line {
					if x == 0 || x == len(line)-1 {
						continue
					}
					if char == 'A' {
						candidates = append(candidates, Cord{X: x, Y: y})
					}
				}
			}
			count := 0
			for _, c := range candidates {
				if lines[c.Y-1][c.X-1] == 'M' && lines[c.Y-1][c.X+1] == 'S' &&
					lines[c.Y+1][c.X-1] == 'M' && lines[c.Y+1][c.X+1] == 'S' {
					count++
				}
				if lines[c.Y-1][c.X-1] == 'S' && lines[c.Y-1][c.X+1] == 'M' &&
					lines[c.Y+1][c.X-1] == 'S' && lines[c.Y+1][c.X+1] == 'M' {
					count++
				}
				if lines[c.Y-1][c.X-1] == 'S' && lines[c.Y-1][c.X+1] == 'S' &&
					lines[c.Y+1][c.X-1] == 'M' && lines[c.Y+1][c.X+1] == 'M' {
					count++
				}
				if lines[c.Y-1][c.X-1] == 'M' && lines[c.Y-1][c.X+1] == 'M' &&
					lines[c.Y+1][c.X-1] == 'S' && lines[c.Y+1][c.X+1] == 'S' {
					count++
				}
			}
			return count
		},
	}
}
