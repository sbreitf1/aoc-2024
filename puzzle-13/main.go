package main

// https://adventofcode.com/2024/day/13

import (
	"aoc/helper"
	"fmt"
	"strings"

	"github.com/alex-ant/gomath/gaussian-elimination"
	"github.com/alex-ant/gomath/rational"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")
	machines := ParseClawMachines(lines)

	solution1 := SumMinTokens(machines)
	fmt.Println("-> part 1:", solution1)

	machines2 := ConvertMachinesToPart2(machines)
	solution2 := SumMinTokens(machines2)
	fmt.Println("-> part 2:", solution2)
}

type ClawMachine struct {
	ButtonA, ButtonB helper.Vec2D[int64]
	Prize            helper.Vec2D[int64]
}

func ParseClawMachines(lines []string) []ClawMachine {
	machines := make([]ClawMachine, 0)
	var buttonA, buttonB, prize helper.Vec2D[int64]
	for _, l := range lines {
		ints := helper.ExtractInts[int64](l)
		p := helper.Vec2D[int64]{X: ints[0], Y: ints[1]}
		if strings.HasPrefix(l, "Button A:") {
			buttonA = p
		} else if strings.HasPrefix(l, "Button B:") {
			buttonB = p
		} else if strings.HasPrefix(l, "Prize:") {
			prize = p
			machines = append(machines, ClawMachine{
				ButtonA: buttonA,
				ButtonB: buttonB,
				Prize:   prize,
			})
		}
	}
	return machines
}

func SumMinTokens(machines []ClawMachine) int64 {
	var sum int64
	for _, m := range machines {
		cost, _ := m.GetMinTokens()
		sum += cost
	}
	return sum
}

func (m ClawMachine) GetMinTokens() (int64, bool) {
	// https://github.com/alex-ant/gomath

	nr := func(i int64) rational.Rational {
		return rational.New(i, 1)
	}

	equations := make([][]rational.Rational, 2)
	equations[0] = []rational.Rational{nr(m.ButtonA.X), nr(m.ButtonB.X), nr(m.Prize.X)}
	equations[1] = []rational.Rational{nr(m.ButtonA.Y), nr(m.ButtonB.Y), nr(m.Prize.Y)}

	res, err := gaussian.SolveGaussian(equations, false)
	if err != nil {
		panic(err)
	}

	if len(res) != 2 {
		return 0, false
	}
	for _, v := range res {
		if len(v) != 1 || !v[0].IsNatural() {
			return 0, false
		}
	}

	a := res[0][0].GetNumerator() / res[0][0].GetDenominator()
	b := res[1][0].GetNumerator() / res[1][0].GetDenominator()

	return 3*a + b, false
}

func ConvertMachinesToPart2(m []ClawMachine) []ClawMachine {
	m2 := helper.Clone(m)
	for i := range m2 {
		m2[i].Prize = m2[i].Prize.Add(helper.Vec2D[int64]{X: 10000000000000, Y: 10000000000000})
	}
	return m2
}
