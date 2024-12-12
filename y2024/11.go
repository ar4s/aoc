package y2024

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/ar4s/aoc/types"
	lru "github.com/hashicorp/golang-lru/v2"
)

type A sync.Map

type Stone uint
type Stones []Stone
type StoneIndex map[Stone]uint

type CacheKey string

func (s Stone) EvenNumberOfDigits() bool {
	a := int(math.Log10(float64(s))) + 1
	return a%2 == 0
}
func (s Stone) SplitInHalf() (Stone, Stone) {
	// split numer into two halves
	// e.g. 1234 -> 12, 34
	// 27 -> 2, 7
	// TODO: do it better
	str := strconv.Itoa(int(s))
	mid := len(str) / 2
	a, _ := strconv.Atoi(str[:mid])
	b, _ := strconv.Atoi(str[mid:])
	return Stone(a), Stone(b)
}

func parseLine(line string) Stones {
	var stones Stones
	for _, s := range strings.Split(line, " ") {
		a, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		stones = append(stones, Stone(a))
	}
	return stones
}

type IndexOperation func(Stone, int)

func transformStone(l *lru.Cache[CacheKey, uint], s Stone, blinks int) uint {
	if blinks == 0 {
		return 1
	}

	value, ok := l.Get(CacheKey(fmt.Sprintf("%d-%d", s, blinks)))
	if ok {
		return value
	}

	v := uint(0)

	if s == 0 {
		stone := Stone(1)
		v = transformStone(l, stone, blinks-1)

	} else if s.EvenNumberOfDigits() {
		a, b := s.SplitInHalf()
		v = transformStone(l, a, blinks-1) + transformStone(l, b, blinks-1)

	} else {
		stone := Stone(int(s) * 2024)
		v = transformStone(l, stone, blinks-1)
	}
	if !ok {
		l.Add(CacheKey(fmt.Sprintf("%d-%d", s, blinks)), v)
	}
	return v
}

func addFactory(stonesIndex StoneIndex) IndexOperation {
	return func(s Stone, value int) {
		if _, ok := stonesIndex[s]; !ok {
			stonesIndex[s] = 0
		}
		stonesIndex[s] += uint(value)
	}
}

func NewPuzzle_11() *types.Puzzle {
	day := 11

	return &types.Puzzle{
		Example: Example(day),
		Input:   Input(day),

		ExampleAExpected: 217443,
		SolutionA: func(lines []string) int {
			stones := parseLine(lines[0])

			l, _ := lru.New[CacheKey, uint](1024 * 10)

			maxBlinks := 75
			sum := uint(0)
			for _, s := range stones {
				sum += transformStone(l, s, maxBlinks)
			}

			return int(sum)
		},

		ExampleBExpected: 0,
		SolutionB: func(lines []string) int {
			// stones := parseLine(lines[0])
			// fmt.Println(stones)
			sum := 0
			// for _, s := range stones {
			// 	sum += len(transformStoneRec(0, s))
			// }
			return sum
		},
	}
}
