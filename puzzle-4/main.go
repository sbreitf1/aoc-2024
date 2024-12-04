package main

// https://adventofcode.com/2024/day/4

import (
	"aoc/helper"
	"fmt"
)

func main() {
	lines := helper.ReadLines("input.txt")
	field := helper.LinesToRunes(lines)

	solution1 := countOccurences(field, "XMAS")
	fmt.Println("-> part 1:", solution1)
	solution2 := countXMAS(field)
	fmt.Println("-> part 2:", solution2)
}

func countOccurences(field [][]rune, needle string) int {
	w := len(field[0])
	h := len(field)
	var count int
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if isMatchAt(field, x, y, len(needle), 1, 0, needle) {
				count++
			}
			if isMatchAt(field, x, y, len(needle), 0, 1, needle) {
				count++
			}
			if isMatchAt(field, x, y, len(needle), 1, 1, needle) {
				count++
			}
			if isMatchAt(field, x, y, len(needle), 1, -1, needle) {
				count++
			}
		}
	}
	return count
}

func isMatchAt(field [][]rune, x, y int, count int, dx, dy int, needle string) bool {
	extracted := extractRunes(field, x, y, count, dx, dy)
	return isMatch(needle, extracted)
}

func extractRunes(field [][]rune, x, y int, count int, dx, dy int) string {
	extracted := make([]rune, 0, count)
	for i := 0; i < count; i++ {
		ypos := y + i*dy
		if ypos < 0 || ypos >= len(field) {
			return ""
		}
		xpos := x + i*dx
		if xpos < 0 || xpos >= len(field[y]) {
			return ""
		}
		extracted = append(extracted, field[ypos][xpos])
	}
	return string(extracted)
}

func isMatch(needle, extracted string) bool {
	if needle == extracted {
		return true
	}
	if needle == string(helper.GetReversedSlice([]rune(extracted))) {
		return true
	}
	return false
}

func countXMAS(field [][]rune) int {
	w := len(field[0])
	h := len(field)
	var count int
	for y := 1; y < h-1; y++ {
		for x := 1; x < w-1; x++ {
			if isXMAS(field, x, y) {
				count++
			}
		}
	}
	return count
}

func isXMAS(field [][]rune, x, y int) bool {
	w1 := extractRunes(field, x-1, y-1, 3, 1, 1)
	w2 := extractRunes(field, x-1, y+1, 3, 1, -1)
	return isMatch("MAS", w1) && isMatch("MAS", w2)
}
