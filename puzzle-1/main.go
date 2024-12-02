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

	solution1 := sumDistances(numbersL, numbersR)
	fmt.Println("-> part 1:", solution1)

	rightCounts := make(map[int]int)
	for _, numR := range numbersR {
		rightCounts[numR] = rightCounts[numR] + 1
	}

	var solution2 int
	for _, num := range numbersL {
		countR := rightCounts[num]
		solution2 += num * countR
	}
	fmt.Println("-> part 2:", solution2)
}

func sumDistances(numbers1, numbers2 []int) int {
	var sum int
	for i := range numbers1 {
		sum += helper.Abs(numbers1[i] - numbers2[i])
	}
	return sum
}
