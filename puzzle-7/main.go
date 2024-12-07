package main

// https://adventofcode.com/2024/day/7

import (
	"aoc/helper"
	"fmt"
	"math"
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
	ops := prepareOperatorConstellations(len(c.Values)-1, allowedOperators)
	for _, op := range ops {
		if val, ok := c.Compute(op); ok && val == c.TestValue {
			return true
		}
	}
	return false
}

func (c Calibration) Compute(ops []rune) (int64, bool) {
	val := c.Values[0]
	for i := range ops {
		if val > c.TestValue {
			return 0, false
		}

		if ops[i] == '+' {
			val += c.Values[i+1]
		} else if ops[i] == '*' {
			val *= c.Values[i+1]
		} else if ops[i] == '|' {
			val, _ = strconv.ParseInt(fmt.Sprintf("%v%v", val, c.Values[i+1]), 10, 64)
		}
	}
	return val, true
}

func prepareOperatorConstellations(count int, allowedOperators []rune) [][]rune {
	ops := make([][]rune, int(math.Pow(float64(len(allowedOperators)), float64(count))))
	for i := range ops {
		ops[i] = make([]rune, count)
		for j := 0; j < count; j++ {
			ops[i][j] = allowedOperators[(i/int(math.Pow(float64(len(allowedOperators)), float64(j))))%len(allowedOperators)]
		}
	}
	return ops
}
