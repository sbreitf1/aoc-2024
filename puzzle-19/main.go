package main

// https://adventofcode.com/2024/day/19

import (
	"aoc/helper"
	"fmt"
	"strings"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")
	patterns, designs := ParseTowelsAndDesigns(lines)

	solution1 := CountValidDesigns(patterns, designs)
	fmt.Println("-> part 1:", solution1)

	solution2 := SumAllArrangements(patterns, designs)
	fmt.Println("-> part 2:", solution2)
}

func ParseTowelsAndDesigns(lines []string) ([]string, []string) {
	patterns := helper.SplitAndTrim(lines[0], ",")
	designs := lines[1:]
	return patterns, designs
}

func CountValidDesigns(patterns []string, designs []string) int {
	var count int
	for _, design := range designs {
		if CountPossibleArrangements(patterns, design) > 0 {
			count++
		}
	}
	return count
}

func SumAllArrangements(patterns []string, designs []string) int {
	var sum int
	for _, design := range designs {
		sum += CountPossibleArrangements(patterns, design)
	}
	return sum
}

var (
	cache map[string]int = make(map[string]int)
)

func CountPossibleArrangements(patterns []string, design string) int {
	if len(design) == 0 {
		return 1
	}

	if val, ok := cache[design]; ok {
		return val
	}

	var count int
	for _, p := range patterns {
		if strings.HasPrefix(design, p) {
			count += CountPossibleArrangements(patterns, design[len(p):])
		}
	}
	cache[design] = count
	return count
}
