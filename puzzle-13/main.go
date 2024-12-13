package main

// https://adventofcode.com/2024/day/13

import (
	"aoc/helper"
	"fmt"
	"strings"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")
	machines := ParseClawMachines(lines)

	solution1 := SumMinTokens(machines)
	fmt.Println("-> part 1:", solution1)
}

type ClawMachine struct {
	ButtonA, ButtonB helper.Vec2D[int]
	Prize            helper.Vec2D[int]
}

func ParseClawMachines(lines []string) []ClawMachine {
	machines := make([]ClawMachine, 0)
	var buttonA, buttonB, prize helper.Vec2D[int]
	for _, l := range lines {
		ints := helper.ExtractInts[int](l)
		p := helper.Vec2D[int]{X: ints[0], Y: ints[1]}
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

func SumMinTokens(machines []ClawMachine) int {
	var sum int
	for _, m := range machines {
		cost, _ := m.GetMinTokens()
		sum += cost
	}
	return sum
}

func (m ClawMachine) GetMinTokens() (int, bool) {
	type ButtonPresses struct {
		A, B int
		Cost int
		Pos  helper.Vec2D[int]
	}

	queue := helper.NewPriorityQueue[int, ButtonPresses]()
	queue.Push(0, ButtonPresses{A: 0, B: 0, Cost: 0, Pos: helper.Vec2D[int]{X: 0, Y: 0}})
	visited := make(map[ButtonPresses]bool)

	for queue.Len() > 0 {
		bp, _ := queue.Pop()
		if visited[bp] {
			continue
		}
		visited[bp] = true

		if bp.Pos == m.Prize {
			return bp.Cost, true
		}

		if bp.A > 100 || bp.B > 100 {
			continue
		}
		if bp.Pos.X > m.Prize.X || bp.Pos.Y > m.Prize.Y {
			continue
		}

		bpa := ButtonPresses{A: bp.A + 1, B: bp.B, Cost: bp.Cost + 3, Pos: helper.Vec2D[int]{X: bp.Pos.X + m.ButtonA.X, Y: bp.Pos.Y + m.ButtonA.Y}}
		queue.Push(bpa.Cost, bpa)
		bpb := ButtonPresses{A: bp.A, B: bp.B + 1, Cost: bp.Cost + 1, Pos: helper.Vec2D[int]{X: bp.Pos.X + m.ButtonB.X, Y: bp.Pos.Y + m.ButtonB.Y}}
		queue.Push(bpb.Cost, bpb)
	}
	return 0, false
}
