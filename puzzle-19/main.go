package main

// https://adventofcode.com/2024/day/19

import (
	"aoc/helper"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")
	pattern, designs := ParseTowelsAndDesigns(lines)

	solution1 := CountValidDesigns(pattern, designs)
	fmt.Println("-> part 1:", solution1)

	solution2 := 0
	fmt.Println("-> part 2:", solution2)
}

func ParseTowelsAndDesigns(lines []string) (*regexp.Regexp, []string) {
	patterns := helper.SplitAndTrim(lines[0], ",")
	pattern := regexp.MustCompile("^(" + strings.Join(patterns, "|") + ")+$")
	designs := lines[1:]
	return pattern, designs
}

func CountValidDesigns(pattern *regexp.Regexp, designs []string) int {
	var count int
	for _, design := range designs {
		if pattern.MatchString(design) {
			count++
		}
	}
	return count
}
