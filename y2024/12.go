package y2024

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/ar4s/aoc/types"
)

type Region struct {
	Plots []Plot
}
type Regions map[string]Region

func regionKey(value string, id int) string {
	return fmt.Sprintf("%s-%d", value, id)
}

type Plot struct {
	Pos   types.Cord2D
	Value string
	Area  int
}
type Garden map[types.Cord2D]Plot

func (g Garden) PlotFromDirection(c types.Cord2D, d types.Direction2D) (Plot, error) {
	n := c.ApplyDirection(d)
	if n.X < 0 || n.Y < 0 {
		return Plot{}, fmt.Errorf("out of bound")
	}
	return g[n], nil
}

func detectRegion(g Garden, target Plot, dir types.Direction2D) ([]Plot, error) {
	candidate, err := g.PlotFromDirection(target.Pos, dir)
	if err != nil {
		return []Plot{}, fmt.Errorf("not found")
	}
	if target.Value != candidate.Value {
		return []Plot{}, fmt.Errorf("not found")
	}
	res := []Plot{}
	for _, d := range types.DIR_4 {
		_ = d
		fmt.Println(d)
		// _, err := detectRegion(g, target, d)
		if err != nil {
			continue
		}
		// res = append(res, r...)
	}
	return res, nil
}

func NewPuzzle_12() *types.Puzzle {
	day := 12
	var parseLine = func(line string, y int, plots Garden) {
		for x, l := range strings.Split(line, "") {
			plots[types.Cord2D{X: x, Y: y}] = Plot{
				Pos:   types.Cord2D{X: x, Y: y},
				Value: l,
				Area:  1,
			}
		}
	}

	return &types.Puzzle{
		Example: Example(day),
		Input:   Input(day),

		ExampleAExpected: 0,
		SolutionA: func(lines []string) int {
			garden := Garden{}

			for y, line := range lines {
				parseLine(line, y, garden)
			}
			// visited := make([]Plot, maxIndex*maxIndex)

			maxIndex := len(lines)
			for y := 0; y <= maxIndex; y++ {
				for x := 0; x <= maxIndex; x++ {
					cord := types.Cord2D{X: x, Y: y}

					p := garden[cord]
					if p.Area == 0 {
						continue
					}
					for _, dir := range types.DIR_4 {
						c, err := garden.PlotFromDirection(p.Pos, dir)
						cd := p.Pos.ApplyDirection(dir)
						if err != nil {
							continue
						}
						if c.Value != p.Value {
							continue
						}
						p.Area++
						garden[cord] = p
						delete(garden, cd)
					}
				}
			}
			fmt.Printf("%+v\n", garden)
			fmt.Println(len(garden))

			return -1
		},

		ExampleBExpected: 0,
		SolutionB: func(lines []string) int {
			return -1
		},
	}
}
