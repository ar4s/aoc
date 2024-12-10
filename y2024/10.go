package y2024

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/ar4s/aoc/types"
	"github.com/samber/lo"
)

type TrailPoint struct {
	X      int
	Y      int
	Height int
}

func (t TrailPoint) Cord2D() types.Cord2D {
	return types.Cord2D{
		t.X, t.Y,
	}
}

type TrailPointsMap map[types.Cord2D]TrailPoint

func findStartingPoints(m TrailPointsMap) []TrailPoint {
	r := []TrailPoint{}

	for _, point := range m {
		if point.Height == 0 {
			r = append(r, point)
		}
	}
	return r
}

func getNextPoint(dir types.Direction2D, f TrailPoint, m TrailPointsMap) (TrailPoint, error) {
	n := f.Cord2D().ApplyDirection(dir)
	if n.X < 0 || n.Y < 0 {
		return TrailPoint{}, fmt.Errorf("out of bound")
	}
	return m[n], nil
}

func traverse(start TrailPoint, dir types.Direction2D, m TrailPointsMap, level int) ([]TrailPoint, error) {
	if level > 9 {
		return []TrailPoint{}, fmt.Errorf("reached limit")
	}

	target, err := getNextPoint(dir, start, m)
	if err != nil {
		return []TrailPoint{}, err
	}
	fmt.Printf("%s%s ", strings.Repeat(" ", level+1), ">")
	fmt.Printf("start=%+v target=%+v dir=%+v\n", start, target.Height, dir)

	if start.Height-target.Height != -1 {
		return []TrailPoint{}, fmt.Errorf("wrong direction")
	}

	if target.Height == 9 {
		fmt.Println("found ", target)
		return []TrailPoint{target}, nil
	}
	sum := []TrailPoint{}
	for _, d := range types.DIR_4 {
		// dont look back
		if d == dir.Negative() {
			continue
		}
		// fmt.Printf("%s%s ", strings.Repeat(" ", level+1), ">")
		// fmt.Printf("start=%+v target=%+v\n", start, target)
		a, _ := traverse(target, d, m, level+1)
		sum = append(sum, a...)
	}
	return sum, nil
}

func traverseTrail(start TrailPoint, m TrailPointsMap, uniq bool) int {
	// score how many different trails from 0 to 9
	// only 1 diff and always incresing
	all := []TrailPoint{}
	for _, dir := range types.DIR_4 {
		a, _ := traverse(start, dir, m, 0)
		all = append(all, a...)

	}

	if uniq {
		return len(lo.Uniq(all))
	}
	return len(all)
}

func NewPuzzle_10() *types.Puzzle {
	day := 10
	var parseLine = func(m TrailPointsMap, line string, y int) {
		for x, point := range line {

			r, err := strconv.Atoi(string(point))
			if err != nil {
				continue
			}
			p := types.Cord2D{
				X: x,
				Y: y,
			}
			m[p] = TrailPoint{
				X:      x,
				Y:      y,
				Height: r,
			}
		}
	}

	return &types.Puzzle{
		Example: Example(day),
		Input:   Input(day),

		ExampleAExpected: 36,
		SolutionA: func(lines []string) int {
			pointsMap := TrailPointsMap{}
			for y, line := range lines {
				parseLine(pointsMap, line, y)
			}
			startingPoints := findStartingPoints(pointsMap)
			fmt.Println("starting points", startingPoints)
			sum := 0
			for _, s := range startingPoints {
				sum += traverseTrail(s, pointsMap, true)
			}
			return sum
		},

		ExampleBExpected: 0,
		SolutionB: func(lines []string) int {
			pointsMap := TrailPointsMap{}
			for y, line := range lines {
				parseLine(pointsMap, line, y)
			}
			startingPoints := findStartingPoints(pointsMap)
			fmt.Println("starting points", startingPoints)
			sum := 0
			for _, s := range startingPoints {
				sum += traverseTrail(s, pointsMap, false)
			}
			return sum
		},
	}
}
