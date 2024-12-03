package main

// https://adventofcode.com/2024/day/3

import (
	"aoc/helper"
	"fmt"
	"regexp"
)

func main() {
	input := helper.ReadString("input.txt")

	instructions := parseInstructions(input)
	solution1 := execAndSumAllMulInstructions(instructions)
	fmt.Println("-> part 1:", solution1)
	solution2 := execAndSumConditionalMulInstructions(instructions)
	fmt.Println("-> part 2:", solution2)
}

const (
	InstructionTypeMul  InstructionType = "mul"
	InstructionTypeDo   InstructionType = "do"
	InstructionTypeDont InstructionType = "dont"
)

type InstructionType string

type Instruction struct {
	Type InstructionType
	A, B int
}

func parseInstructions(str string) []Instruction {
	patternInstruction := regexp.MustCompile(`(mul|do|don't)\((|\d+,\d+)\)`)
	patternMulArgs := regexp.MustCompile(`^(\d+),(\d+)$`)
	matches := patternInstruction.FindAllStringSubmatch(str, -1)
	instructions := make([]Instruction, 0)
	for _, m := range matches {
		if m[1] == "mul" {
			mArgs := patternMulArgs.FindStringSubmatch(m[2])
			if len(mArgs) == 3 {
				instructions = append(instructions, Instruction{
					Type: InstructionTypeMul,
					A:    helper.ParseInt[int](mArgs[1]),
					B:    helper.ParseInt[int](mArgs[2]),
				})
			}
		} else if m[1] == "do" {
			if len(m[2]) == 0 {
				instructions = append(instructions, Instruction{
					Type: InstructionTypeDo,
				})
			}
		} else if m[1] == "don't" {
			if len(m[2]) == 0 {
				instructions = append(instructions, Instruction{
					Type: InstructionTypeDont,
				})
			}
		}
	}
	return instructions
}

func execAndSumAllMulInstructions(instructions []Instruction) int {
	var sum int
	for _, inst := range instructions {
		if inst.Type == InstructionTypeMul {
			sum += inst.A * inst.B
		}
	}
	return sum
}

func execAndSumConditionalMulInstructions(instructions []Instruction) int {
	var sum int
	doExec := true
	for _, inst := range instructions {
		switch inst.Type {
		case InstructionTypeMul:
			if doExec {
				sum += inst.A * inst.B
			}
		case InstructionTypeDo:
			doExec = true
		case InstructionTypeDont:
			doExec = false
		}
	}
	return sum
}
