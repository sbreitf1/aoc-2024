package main

// https://adventofcode.com/2024/day/11

import (
	"aoc/helper"
	"fmt"
)

func main() {
	str := helper.ReadString("example-1.txt")
	stones := parseStones(str)

	solution1 := len(stones.Blink(25))
	fmt.Println("-> part 1:", solution1)

	/*solution2 := len(stones.Blink(75))
	fmt.Println("-> part 2:", solution2)*/
}

type Stones []int64

func parseStones(str string) Stones {
	return Stones(helper.ExtractInts[int64](str))
}

func (stones Stones) Blink(n int) Stones {
	next := helper.Clone(stones)
	//fmt.Println(next)
	for i := 0; i < n; i++ {
		next = next.blinkOneTime()
		//fmt.Println(next)
	}
	return next
}

func (stones Stones) blinkOneTime() Stones {
	next := make(Stones, 0, len(stones))
	for _, s := range stones {
		str := fmt.Sprintf("%v", s)
		if s == 0 {
			next = append(next, 1)
		} else if len(str)%2 == 0 {
			next = append(next, helper.ParseInt[int64](str[:len(str)/2]), helper.ParseInt[int64](str[len(str)/2:]))
		} else {
			next = append(next, s*2024)
		}
	}
	return next
}
