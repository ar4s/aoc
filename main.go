package main

import (
	"fmt"

	"github.com/ar4s/aoc/types"
	"github.com/ar4s/aoc/y2024"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func runExamples(puzzle *types.Puzzle) {
	outputExampleA := puzzle.RunExampleA()
	outputExampleB := puzzle.RunExampleB()

	resultA := "❌"
	if outputExampleA == puzzle.ExampleAExpected {
		resultA = "✅"
	}

	resultB := "❌"
	if outputExampleB == puzzle.ExampleBExpected {
		resultB = "✅"
	}

	rows := [][]string{
		{"Example A", fmt.Sprintf("%d", puzzle.ExampleAExpected), fmt.Sprintf("%d", outputExampleA), fmt.Sprintf("%s", resultA)},
		{"Example B", fmt.Sprintf("%d", puzzle.ExampleBExpected), fmt.Sprintf("%d", outputExampleB), fmt.Sprintf("%s", resultB)},
	}
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		Headers("Task", "Expected", "Result", "Check").
		Rows(rows...)

	fmt.Println(t.Render())
}

func runProblems(puzzle *types.Puzzle) {
	outputA := puzzle.RunSolutionA()
	outputB := puzzle.RunSolutionB()
	rows := [][]string{
		{"Solution A", fmt.Sprintf("%d", outputA)},
		{"Solution B", fmt.Sprintf("%d", outputB)},
	}
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("9"))).
		Headers("Task", "Result").
		Rows(rows...)

	fmt.Println(t.Render())
}

func main() {
	puzzle := y2024.NewPuzzle_07()
	runExamples(puzzle)
	runProblems(puzzle)
}
