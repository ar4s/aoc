package y2024

import (
	"embed"
	"fmt"
)

//go:embed inputs/*.txt
var inputsFS embed.FS

//go:embed examples/*.txt
var examplesFS embed.FS

func dayValidator(day int) {
	if day < 1 {
		panic("Day must be greater than 0")
	}
	if day > 25 {
		panic("Day must be less than 26")
	}
}

func Example(day int) string {
	dayValidator(day)
	path := fmt.Sprintf("examples/%02d.txt", day)
	example, err := examplesFS.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(example)
}

func Input(day int) string {
	dayValidator(day)
	path := fmt.Sprintf("inputs/%02d.txt", day)

	input, err := inputsFS.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(input)
}
