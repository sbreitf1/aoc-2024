package main

// https://adventofcode.com/2024/day/17

import (
	"aoc/helper"
	"fmt"
	"math"
	"slices"
	"strings"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")
	computer := ParseComputer(lines)

	/*solution1 := formatOutput(computer.Exec())
	fmt.Println("-> part 1:", solution1)*/

	solution2 := findPart2(&computer)
	fmt.Println("-> part 2:", solution2)
}

type Computer struct {
	Registers          []int64
	Program            []Instruction
	InstructionPointer int
}

type Instruction struct {
	OpCode  int
	Operand int
}

func ParseComputer(lines []string) Computer {
	registers := make([]int64, 3)
	program := make([]Instruction, 0)
	for _, l := range lines {
		if strings.HasPrefix(l, "Register ") {
			registers[l[9]-'A'] = helper.ExtractInts[int64](l)[0]
		} else if strings.HasPrefix(l, "Program") {
			ints := helper.ExtractInts[int](l)
			for i := 0; i < len(ints); i += 2 {
				program = append(program, Instruction{
					OpCode:  ints[i],
					Operand: ints[i+1],
				})
			}
		}
	}
	return Computer{
		Registers: registers,
		Program:   program,
	}
}

func (computer *Computer) Exec() []int64 {
	output := make([]int64, 0)

	for computer.InstructionPointer < len(computer.Program) {
		inst := computer.Program[computer.InstructionPointer]

		increaseInstructionPointer := true

		switch inst.OpCode {
		case 0: // adv -> div to a
			numerator := computer.Registers[0]
			denominator := int64(math.Pow(2, float64(computer.EvalComboOperand(inst.Operand))))
			computer.Registers[0] = numerator / denominator

		case 1: // bxl -> bitwise xor b/op
			val1 := computer.Registers[1]
			val2 := int64(inst.Operand)
			computer.Registers[1] = val1 ^ val2

		case 2: // bst -> mod 8
			val := computer.EvalComboOperand(inst.Operand)
			computer.Registers[1] = val & 7

		case 3: // jnz -> jump if not zero
			if computer.Registers[0] != 0 {
				computer.InstructionPointer = inst.Operand
				increaseInstructionPointer = false
			}

		case 4: // bxc -> bitwise xor b/c
			val1 := computer.Registers[1]
			val2 := computer.Registers[2]
			computer.Registers[1] = val1 ^ val2

		case 5: // out -> append to output
			output = append(output, computer.EvalComboOperand(inst.Operand)&7)

		case 6: // bdv -> div to b
			numerator := computer.Registers[0]
			denominator := int64(math.Pow(2, float64(computer.EvalComboOperand(inst.Operand))))
			computer.Registers[1] = numerator / denominator

		case 7: // cdv -> div to c
			numerator := computer.Registers[0]
			denominator := int64(math.Pow(2, float64(computer.EvalComboOperand(inst.Operand))))
			computer.Registers[2] = numerator / denominator

		default:
			helper.ExitWithMessage("OpCode %d is invalid", inst.OpCode)
		}

		if increaseInstructionPointer {
			computer.InstructionPointer++
		}
	}

	return output
}

func (computer *Computer) EvalComboOperand(operand int) int64 {
	if operand >= 0 && operand <= 3 {
		return int64(operand)
	}
	if operand >= 4 && operand <= 6 {
		return computer.Registers[operand-4]
	}
	helper.ExitWithMessage("combo operand %d is invalid", operand)
	return -1
}

func (computer *Computer) ResetAndExec(registerA int64) []int64 {
	computer.Registers = []int64{registerA, 0, 0}
	computer.InstructionPointer = 0
	return computer.Exec()
}

func formatOutput(output []int64) string {
	outputStr := helper.MapValues(output, func(val int64) string { return fmt.Sprintf("%d", val) })
	return strings.Join(outputStr, ",")
}

func findPart2(computer *Computer) int64 {
	targetOutput := make([]int64, 0, len(computer.Program)*2)
	for _, inst := range computer.Program {
		targetOutput = append(targetOutput, int64(inst.OpCode), int64(inst.Operand))
	}

	// estimate range for binary search
	min, max := findBinarySearchRange(computer, targetOutput)
	fmt.Println("binary search between", min, "and", max)

	maxBruteForceIterations := int64(1000)
	for i := 0; i < 10; i++ {
		if (max - min) <= maxBruteForceIterations {
			fmt.Println("maxBruteForceIterations reached")
			break
		}

		fmt.Println("current range is", min, "to", max, "(", max-min, ")")
		val := (min + max) / 2
		valOutput := computer.ResetAndExec(val)
		if len(valOutput) < len(targetOutput) {
			min = val
			continue
		}
		if len(valOutput) > len(targetOutput) {
			max = val
			continue
		}
		matchingValueCount := countMatchingValuesAtEndForRange(computer, val-3, val+3, targetOutput)

		/*matchingValueCountUp := countMatchingValuesAtEndForRange(computer, (val+max)/2-3, (val+max)/2+3, targetOutput)
		matchingValueCountDown := countMatchingValuesAtEndForRange(computer, (min+val)/2-3, (min+val)/2+3, targetOutput)

		if matchingValueCountUp > matchingValueCountDown {
			min = val
			continue
		} else if matchingValueCountDown > matchingValueCountUp {
			max = val
			continue
		}

		for v := val - 3; v < val+3; v++ {
			fmt.Println(computer.ResetAndExec(v))
		}*/

		fmt.Println(valOutput)
		fmt.Println(targetOutput)
		//fmt.Println(matchingValueCount)
		//fmt.Println(matchingValueCount, "->", len(targetOutput)-matchingValueCount)

		fmt.Println("compare", valOutput[len(valOutput)-matchingValueCount-1], "to", targetOutput[len(targetOutput)-matchingValueCount-1])
		if valOutput[len(valOutput)-matchingValueCount-1] < targetOutput[len(targetOutput)-matchingValueCount-1] {
			min = val
			continue
		} else {
			max = val
			continue
		}
	}
	/*for i := 0; i < 10; i++ {
		if (max - min) <= maxBruteForceIterations {
			fmt.Println("maxBruteForceIterations reached")
			break
		}

		fmt.Println("current range is", min, "to", max, "(", max-min, ")")
		val := (min + max) / 2

		for v := min; v <= max; v++ {
			fmt.Println(computer.ResetAndExec(v))
		}
		os.Exit(0)

		valOutput := computer.ResetAndExec(val)
		if len(valOutput) < len(targetOutput) {
			min = val
			continue
		}
		if len(valOutput) > len(targetOutput) {
			max = val
			continue
		}
		matchingValueCount := countMatchingValuesAtEndForRange(computer, val-3, val+3, targetOutput)
		fmt.Println(valOutput)
		fmt.Println(targetOutput)
		//fmt.Println(matchingValueCount)
		//fmt.Println(matchingValueCount, "->", len(targetOutput)-matchingValueCount)

		fmt.Println("compare", valOutput[len(valOutput)-matchingValueCount-1], "to", targetOutput[len(targetOutput)-matchingValueCount-1])
		if valOutput[len(valOutput)-matchingValueCount-1] < targetOutput[len(targetOutput)-matchingValueCount-1] {
			min = val
			continue
		} else {
			max = val
			continue
		}
	}*/

	fmt.Println("brute-force search between", min, "and", max)
	if (max - min) <= maxBruteForceIterations {
		for val := min; val <= max; val++ {
			output := computer.ResetAndExec(val)
			if slices.Equal(output, targetOutput) {
				return val
			}
		}
	}

	helper.ExitWithMessage("no solution for part 2 found!")
	return -1
}

func findBinarySearchRange(computer *Computer, targetOutput []int64) (int64, int64) {
	var max int64 = 1
	for {
		output := computer.ResetAndExec(max)
		if len(output) > len(targetOutput) {
			break
		}
		max *= 10
	}
	min := max / 100
	return min, max
}

func countMatchingValuesAtEndForRange(computer *Computer, min, max int64, targetOutput []int64) int {
	minMatchingValues := len(targetOutput)
	for val := min; val <= max; val++ {
		output := computer.ResetAndExec(val)
		minMatchingValues = helper.Min(minMatchingValues, countMatchingValuesAtEnd(output, targetOutput))
	}
	return minMatchingValues
}

func countStableValuesAtEnd(computer *Computer, min, max int64) int {
	cmpOutput := computer.ResetAndExec(min)
	return countMatchingValuesAtEndForRange(computer, min+1, max, cmpOutput)
}

func countMatchingValuesAtEnd(output, targetOutput []int64) int {
	for i := 0; i < helper.Min(len(output), len(targetOutput)); i++ {
		if output[len(output)-i-1] != targetOutput[len(targetOutput)-i-1] {
			return i
		}
	}
	return helper.Min(len(output), len(targetOutput))
}
