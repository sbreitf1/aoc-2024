package main

// https://adventofcode.com/2024/day/1

import (
	"aoc/helper"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	lines := helper.ReadLines("input.txt")

	numbersL, numbersR := parseSortedNumberGroups(lines)

	solution1 := sumDistances(numbersL, numbersR)
	fmt.Println("-> part 1:", solution1)

	solution2 := computeSimilarity(numbersL, numbersR)
	fmt.Println("-> part 2:", solution2)
}

func parseSortedNumberGroups(lines []string) ([]int, []int) {
	pattern := regexp.MustCompile(`(\d+)\s+(\d+)`)
	numbersL := make([]int, 0)
	numbersR := make([]int, 0)
	for _, line := range lines {
		if m := pattern.FindStringSubmatch(line); len(m) == 3 {
			numL, _ := strconv.ParseInt(m[1], 10, 32)
			numR, _ := strconv.ParseInt(m[2], 10, 32)
			numbersL = append(numbersL, int(numL))
			numbersR = append(numbersR, int(numR))
		}
	}

	sort.Ints(numbersL)
	sort.Ints(numbersR)

	return numbersL, numbersR
}

func sumDistances(numbersL, numbersR []int) int {
	var sum int
	for i := range numbersL {
		sum += helper.Abs(numbersL[i] - numbersR[i])
	}
	return sum
}

func computeSimilarity(numbersL, numbersR []int) int {
	rightCounts := make(map[int]int)
	for _, numR := range numbersR {
		rightCounts[numR] = rightCounts[numR] + 1
	}

	var similarity int
	for _, num := range numbersL {
		countR := rightCounts[num]
		similarity += num * countR
	}

	return similarity
}
