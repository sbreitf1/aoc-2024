package main

// https://adventofcode.com/2024/day/2

import (
	"aoc/helper"
	"fmt"
)

func main() {
	lines := helper.ReadLines("input.txt")
	levels := parseLevels(lines)

	solution1 := countSafeLevels(levels)
	fmt.Println("-> part 1:", solution1)

	solution2 := countDampenedSafeLevels(levels)
	fmt.Println("-> part 2:", solution2)
}

func parseLevels(lines []string) [][]int {
	levels := make([][]int, 0, len(lines))
	for _, line := range lines {
		levels = append(levels, helper.ExtractInts[int](line))
	}
	return levels
}

func countSafeLevels(levels [][]int) int {
	var count int
	for _, level := range levels {
		if isSafe(level) {
			count++
		}
	}
	return count
}

func isSafe(level []int) bool {
	var lastSign int
	for i := 1; i < len(level); i++ {
		step := level[i] - level[i-1]
		absStep := helper.Abs(step)
		if absStep < 1 || absStep > 3 {
			return false
		}
		if i > 1 {
			if lastSign != helper.Sign(step) {
				return false
			}
		}
		lastSign = helper.Sign(step)
	}

	return true
}

func countDampenedSafeLevels(levels [][]int) int {
	var count int
	for _, level := range levels {
		if isSafeWithDampening(level) {
			count++
		}
	}
	return count
}

func isSafeWithDampening(level []int) bool {
	for ignoredIndex := range level {
		dampenedLevel := helper.RemoveIndex(level, ignoredIndex)
		if isSafe(dampenedLevel) {
			return true
		}
	}
	return false
}
