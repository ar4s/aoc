package y2024

import (
	_ "embed"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/ar4s/aoc/types"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

var keepInstructionPointerError = fmt.Errorf("keep instruction pointer")

type Base int

type InstructionPointer Base
type Opcode Base
type Operand Base

const (
	ADV Opcode = iota
	BXL
	BST
	JNZ
	BXC
	OUT
	BDV
	CDV
)

const (
	RA Operand = iota + 4
	RB
	RC
	RESERVED
)

type Registers struct {
	A Base
	B Base
	C Base
}

type Program []Base
type Output []Base

type Instruction struct {
	opcode  Opcode
	operand Operand
}

func (i Instruction) String() string {
	return fmt.Sprintf("%v %v", i.opcode, i.operand)
}

func (o Operand) Raw() Base {
	return Base(o)
}

func (o Operand) Value(r Registers) Base {
	switch o {
	case 0, 1, 2, 3, 7:
		return Base(o)
	case RA:
		return Base(r.A)
	case RB:
		return Base(r.B)
	case RC:
		return Base(r.C)
	}
	return -1
}

func (o Base) Modulo() Base {
	return o & 7
}

func (i Instruction) Execute(cpu *CPU) error {
	switch i.opcode {
	case ADV:
		cpu.Registers.A >>= i.operand.Value(cpu.Registers)
	case BDV:
		cpu.Registers.B = cpu.Registers.A >> i.operand.Value(cpu.Registers)
	case CDV:
		cpu.Registers.C = cpu.Registers.A >> i.operand.Value(cpu.Registers)
	case BXL:
		cpu.Registers.B ^= i.operand.Raw()
	case BST:
		cpu.Registers.B = i.operand.Value(cpu.Registers).Modulo()
	case JNZ:
		if cpu.Registers.A == 0 {
			return nil
		}
		cpu.IP = int(i.operand)
		return keepInstructionPointerError
	case BXC:
		cpu.Registers.B ^= cpu.Registers.C
	case OUT:
		cpu.Output = append(cpu.Output, Base(i.operand.Value(cpu.Registers).Modulo()))
	}
	return nil
}

type CPU struct {
	IP        int
	Registers Registers
	Program   Program
	Output    Output
}

func (o Output) String() string {
	outputStrings := make([]string, len(o))
	for i, v := range o {
		outputStrings[i] = fmt.Sprint(v)
	}
	return strings.Join(outputStrings, ",")
}

func NewComputer() CPU {
	return CPU{
		Registers: Registers{},
		Program:   Program{},
	}
}
func (c CPU) String() string {
	return fmt.Sprintf("IP: %v, Registers: %v, Program: %v, Output: %v", c.IP, c.Registers, c.Program, c.Output)
}

func (c *CPU) Reset() {
	c.IP = 0
	c.Output = Output{}
	c.Registers = Registers{}
}

func (c *CPU) Run() {
	for c.IP < len(c.Program)-1 {
		logger.Info("Run", "IP", c.IP, "opcode", c.Program[c.IP], "operand", c.Program[c.IP+1], "Registers", c.Registers)
		instruction := Instruction{
			opcode:  Opcode(c.Program[c.IP]),
			operand: Operand(c.Program[c.IP+1]),
		}
		err := instruction.Execute(c)
		if err != nil {
			if errors.Is(err, keepInstructionPointerError) {
				continue
			}
		}
		c.IP += 2
	}
}

func (p *CPU) POST(pr Program, a, b, c int) {
	p.Program = pr
	p.Run()
	fmt.Println(p)
	p.Registers.ShouldContains(Base(a), Base(b), Base(c))
	fmt.Println("POST OK!\n==================")
	p.Reset()
}

func (r Registers) ShouldContains(a, b, c Base) {
	if r.A != a {
		panic("A should be " + fmt.Sprint(a) + " but is " + fmt.Sprint(r.A))
	}
	if r.B != b {
		panic("B should be " + fmt.Sprint(b) + " but is " + fmt.Sprint(r.B))
	}
	if r.C != c {
		panic("C should be " + fmt.Sprint(c) + " but is " + fmt.Sprint(r.C))
	}
}

func NewPuzzle_17() *types.Puzzle {
	day := 17
	return &types.Puzzle{
		Example: Example(day),
		Input:   Input(day),

		ExampleAExpected: 0,
		SolutionA: func(lines []string) int {
			cpu := NewComputer()

			// A power-on self-test
			// case 1
			cpu.Registers.C = 9
			cpu.POST(Program{2, 6}, 0, 1, 9)

			// case 2
			cpu.Registers.A = 10
			cpu.POST(Program{5, 0, 5, 1, 5, 4}, 10, 0, 0)

			// case 3
			cpu.Registers.A = 2024
			cpu.POST(Program{0, 1, 5, 4, 3, 0}, 0, 0, 0)

			// case 4
			cpu.Registers.B = 29
			cpu.POST(Program{1, 7}, 0, 26, 0)

			// case 5
			cpu.Registers.B = 2024
			cpu.Registers.C = 43690
			cpu.POST(Program{4, 0}, 0, 44354, 43690)

			// example
			cpu.Reset()
			cpu.Registers.A = 729
			cpu.Registers.B = 0
			cpu.Registers.C = 0
			cpu.Program = Program{0, 1, 5, 4, 3, 0}
			cpu.Run()
			fmt.Println("Example:", cpu.Output.String())

			cpu.Reset()
			cpu.Registers.A = 0 // input
			cpu.Registers.B = 0
			cpu.Registers.C = 0
			cpu.Program = Program{} // input
			cpu.Run()
			fmt.Println("Solution:", cpu.Output.String())

			return -1
		},

		ExampleBExpected: 0,
		SolutionB: func(lines []string) int {
			fmt.Println("================\n================\nSolution B")
			pr := Program{} // input
			cpu := NewComputer()
			cpu.Registers.A = 0
			cpu.Registers.B = 0
			cpu.Registers.C = 0
			cpu.Program = pr
			cpu.Run()
			fmt.Println(cpu.Output)
			return -1
		},
	}
}
