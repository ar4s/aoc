package y2023

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ar4s/aoc/types"
	"github.com/samber/lo"
)

//go:embed examples/01_a.txt
var example_a string

//go:embed examples/01_b.txt
var example_b string

//go:embed inputs/01_a.txt
var input_a string

//go:embed inputs/01_b.txt
var input_b string

type Ordered struct {
	Order   int
	Content string
}

type ByOrder []Ordered

func (a ByOrder) Len() int           { return len(a) }
func (a ByOrder) Less(i, j int) bool { return a[i].Order < a[j].Order }
func (a ByOrder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func NewPuzzle_01() *types.Puzzle {
	return &types.Puzzle{
		ExampleA:         example_a,
		InputA:           input_a,
		ExampleAExpected: 142,
		SolutionA: func(lines []string) int {
			sum := 0
			for _, line := range lines {
				first := "0"
				last := "0"
				for _, c := range strings.Split(line, "") {
					cond := strings.Contains("1234567890", c)
					if cond {
						first = c
						break
					}
				}
				for _, c := range lo.Reverse(strings.Split(line, "")) {
					cond := strings.Contains("1234567890", c)
					if cond {
						last = c
						break
					}
				}
				toSum, err := strconv.Atoi(first + last)
				if err != nil {
					panic(err)
				}
				sum += toSum
			}
			return sum
		},

		ExampleB:         example_b,
		InputB:           input_b,
		ExampleBExpected: 281,
		SolutionB: func(lines []string) int {
			sum := 0
			mapOfDigits := map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
				"four":  4,
				"five":  5,
				"six":   6,
				"seven": 7,
				"eight": 8,
				"nine":  9,
			}

			for _, line := range lines {
				first := "0"
				last := "0"

				formattedLine := line
				toReplace := []Ordered{}
				for name, digit := range mapOfDigits {
					_ = digit
					index := strings.Index(formattedLine, name)
					if index != -1 {
						toReplace = append(toReplace, Ordered{index, name})
					}
				}
				sort.Sort(ByOrder(toReplace))
				for _, r := range toReplace {
					formattedLine = strings.Replace(formattedLine, r.Content, fmt.Sprint(mapOfDigits[r.Content]), 1)
				}

				for _, c := range strings.Split(formattedLine, "") {
					cond := strings.Contains("123456789", c)
					if cond {
						first = c
						break
					}
				}
				for _, c := range lo.Reverse(strings.Split(formattedLine, "")) {
					cond := strings.Contains("123456789", c)
					if cond {
						last = c
						break
					}
				}

				toSum, err := strconv.Atoi(fmt.Sprint(first) + fmt.Sprint(last))
				if err != nil {
					panic(err)
				}

				sum += toSum
			}

			return sum
		},
	}
}
