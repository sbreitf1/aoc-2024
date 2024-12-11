package main

// https://adventofcode.com/2024/day/11

import (
	"aoc/helper"
	"fmt"
)

func main() {
	str := helper.ReadString("input.txt")
	stones := parseStones(str)

	solution1 := stones.Blink(25)
	fmt.Println("-> part 1:", solution1)

	solution2 := stones.Blink(75)
	fmt.Println("-> part 2:", solution2)
}

type Stones []Stone
type Stone int64

func parseStones(str string) Stones {
	return Stones(helper.ExtractInts[Stone](str))
}

func (stones Stones) Blink(n int) int64 {
	var count int64
	for _, s := range stones {
		count += s.Blink(n)
	}
	return count
}

type CacheKey struct {
	Stone      Stone
	BlinkCount int
}

var cache map[CacheKey]int64 = make(map[CacheKey]int64)

func (s Stone) Blink(n int) int64 {
	if n == 0 {
		return 1
	}

	key := CacheKey{Stone: s, BlinkCount: n}
	if count, ok := cache[key]; ok {
		return count
	}

	var result int64
	if s == 0 {
		result = Stone(1).Blink(n - 1)
	} else if str := fmt.Sprintf("%v", s); len(str)%2 == 0 {
		stone1 := helper.ParseInt[Stone](str[:len(str)/2])
		stone2 := helper.ParseInt[Stone](str[len(str)/2:])
		result = stone1.Blink(n-1) + stone2.Blink(n-1)
	} else {
		result = Stone(s * 2024).Blink(n - 1)
	}
	cache[key] = result
	return result
}
