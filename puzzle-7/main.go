package main

// https://adventofcode.com/2024/day/7

import (
	"aoc/helper"
	"fmt"
	"strconv"
)

func main() {
	lines := helper.ReadLines("input.txt")
	calibrations := parseCalibrations(lines)

	solution1 := sumSolvableCalibrations(calibrations, []rune{'+', '*'})
	fmt.Println("-> part 1:", solution1)

	solution2 := sumSolvableCalibrations(calibrations, []rune{'+', '*', '|'})
	fmt.Println("-> part 2:", solution2)
}

type Calibration struct {
	TestValue int64
	Values    []int64
}

func parseCalibrations(lines []string) []Calibration {
	calibrations := make([]Calibration, 0, len(lines))
	for _, l := range lines {
		ints := helper.ParseInts[int64](l)
		if len(ints) >= 2 {
			calibrations = append(calibrations, Calibration{
				TestValue: ints[0],
				Values:    ints[1:],
			})
		}
	}
	return calibrations
}

func sumSolvableCalibrations(calibrations []Calibration, allowedOperators []rune) int64 {
	var sum int64
	for _, c := range calibrations {
		if c.IsSolvable(allowedOperators) {
			sum += c.TestValue
		}
	}
	return sum
}

func (c Calibration) IsSolvable(allowedOperators []rune) bool {
	return c.isSolvable(c.Values[0], 1, allowedOperators)
}

func (c Calibration) isSolvable(currentVal int64, pos int, allowedOperators []rune) bool {
	if pos >= len(c.Values) {
		return currentVal == c.TestValue
	}
	if currentVal > c.TestValue {
		return false
	}

	for _, op := range allowedOperators {
		nextVal := applyOperator(op, currentVal, c.Values[pos])
		if c.isSolvable(nextVal, pos+1, allowedOperators) {
			return true
		}
	}
	return false
}

func applyOperator(op rune, val1, val2 int64) int64 {
	switch op {
	case '+':
		return val1 + val2
	case '*':
		return val1 * val2
	case '|':
		val, _ := strconv.ParseInt(fmt.Sprintf("%v%v", val1, val2), 10, 64)
		return val
	default:
		panic(fmt.Sprintf("unexpected operator %q", op))
	}
}
