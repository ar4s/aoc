package types

import (
	"strings"

	"github.com/samber/lo"
)

type Puzzle struct {
	Example string
	Input   string

	ExampleAExpected int
	SolutionA        func([]string) int

	ExampleBExpected int
	SolutionB        func([]string) int
}

func SplitLines(input string) []string {
	return lo.Filter(strings.Split(input, "\n"), func(item string, index int) bool {
		return item != ""
	})
}

func (p *Puzzle) RunExampleA() int {
	return p.SolutionA(SplitLines(p.Example))
}

func (p *Puzzle) RunExampleB() int {
	return p.SolutionB(SplitLines(p.Example))
}

func (p *Puzzle) RunSolutionA() int {
	return p.SolutionA(SplitLines(p.Input))
}

func (p *Puzzle) RunSolutionB() int {
	return p.SolutionB(SplitLines(p.Input))
}
