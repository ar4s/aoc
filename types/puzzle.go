package types

import (
	"strings"

	"github.com/samber/lo"
)

type Puzzle struct {
	ExampleA         string
	ExampleAExpected int
	SolutionA        func([]string) int

	ExampleB         string
	ExampleBExpected int
	SolutionB        func([]string) int

	InputA string
	InputB string
}

func SplitLines(input string) []string {
	return lo.Filter(strings.Split(input, "\n"), func(item string, index int) bool {
		return item != ""
	})
}

func (p *Puzzle) RunExampleA() int {
	return p.SolutionA(SplitLines(p.ExampleA))
}

func (p *Puzzle) RunExampleB() int {
	return p.SolutionB(SplitLines(p.ExampleB))
}

func (p *Puzzle) RunSolutionA() int {
	return p.SolutionA(SplitLines(p.InputA))
}

func (p *Puzzle) RunSolutionB() int {
	return p.SolutionB(SplitLines(p.InputB))
}
