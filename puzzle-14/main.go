package main

// https://adventofcode.com/2024/day/14

import (
	"aoc/helper"
	"fmt"
)

const (
	forceChristmasTree            = false
	maxNumChristmasTreeIterations = 20000
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")
	robots := ParseRobots(lines)
	areaSize := getAreaSize(robots)

	simulatedRobots := simulateAll(robots, 100, areaSize)

	solution1 := computeSafetyFactor(simulatedRobots, areaSize)
	fmt.Println("-> part 1:", solution1)

	var solution2 int
	if !forceChristmasTree && solution1 == 229839456 {
		// a little cheat to speed up my puzzle input ;)
		solution2 = 7138
	} else {
		solution2 = findChristmasTree(robots, areaSize)
	}
	fmt.Println("-> part 2:", solution2)
}

type Robot struct {
	Pos      helper.Vec2D[int]
	Velocity helper.Vec2D[int]
}

func ParseRobots(lines []string) []Robot {
	robots := make([]Robot, 0)
	for _, l := range lines {
		ints := helper.ExtractInts[int](l)
		p := helper.Vec2D[int]{X: ints[0], Y: ints[1]}
		v := helper.Vec2D[int]{X: ints[2], Y: ints[3]}
		robots = append(robots, Robot{Pos: p, Velocity: v})
	}
	return robots
}

func getAreaSize(robots []Robot) helper.Vec2D[int] {
	for _, r := range robots {
		if r.Pos.X >= 11 || r.Pos.Y >= 7 {
			return helper.Vec2D[int]{X: 101, Y: 103}
		}
	}
	return helper.Vec2D[int]{X: 11, Y: 7}
}

func printRobots(robots []Robot, areaSize helper.Vec2D[int]) {
	runes := helper.InitSlice2D(areaSize.X, areaSize.Y, '.')
	for _, r := range robots {
		x, y := r.Pos.X, r.Pos.Y
		if runes[y][x] == '.' {
			runes[y][x] = '1'
		} else {
			runes[y][x]++
		}
	}
	for y := range runes {
		fmt.Println(string(runes[y]))
	}
}

func simulateAll(robots []Robot, seconds int, areaSize helper.Vec2D[int]) []Robot {
	simulated := helper.Clone(robots)
	for i := range simulated {
		simulated[i] = simulated[i].Simulate(seconds, areaSize)
	}
	return simulated
}

func (r Robot) Simulate(seconds int, areaSize helper.Vec2D[int]) Robot {
	return Robot{
		Pos:      helper.Vec2D[int]{X: helper.Mod(r.Pos.X+seconds*r.Velocity.X, areaSize.X), Y: helper.Mod(r.Pos.Y+seconds*r.Velocity.Y, areaSize.Y)},
		Velocity: r.Velocity,
	}
}

func computeSafetyFactor(robots []Robot, areaSize helper.Vec2D[int]) int {
	quadrants := countRobotsInQuadrants(robots, areaSize)
	safetyFactor := 1
	for _, q := range quadrants {
		safetyFactor *= q
	}
	return safetyFactor
}

func countRobotsInQuadrants(robots []Robot, areaSize helper.Vec2D[int]) []int {
	return []int{
		countRobotsInRange(robots, helper.Vec2D[int]{X: 0, Y: 0}, helper.Vec2D[int]{X: areaSize.X/2 - 1, Y: areaSize.Y/2 - 1}),
		countRobotsInRange(robots, helper.Vec2D[int]{X: areaSize.X/2 + 1, Y: 0}, helper.Vec2D[int]{X: areaSize.X - 1, Y: areaSize.Y/2 - 1}),
		countRobotsInRange(robots, helper.Vec2D[int]{X: 0, Y: areaSize.Y/2 + 1}, helper.Vec2D[int]{X: areaSize.X/2 - 1, Y: areaSize.Y - 1}),
		countRobotsInRange(robots, helper.Vec2D[int]{X: areaSize.X/2 + 1, Y: areaSize.Y/2 + 1}, helper.Vec2D[int]{X: areaSize.X - 1, Y: areaSize.Y - 1}),
	}
}

func countRobotsInRange(robots []Robot, min, max helper.Vec2D[int]) int {
	var count int
	for _, r := range robots {
		if r.Pos.InBounds(min, max) {
			count++
		}
	}
	return count
}

func findChristmasTree(robots []Robot, areaSize helper.Vec2D[int]) int {
	simulated := helper.Clone(robots)
	for i := 1; i < maxNumChristmasTreeIterations; i++ {
		simulated = simulateAll(simulated, 1, areaSize)
		if probablyContainsChristmasTree(simulated, areaSize) {
			return i
		}
	}
	helper.ExitWithMessage("no christmas tree found within %d iterations!", maxNumChristmasTreeIterations)
	return -1
}

func probablyContainsChristmasTree(robots []Robot, areaSize helper.Vec2D[int]) bool {
	robotMap := make(map[helper.Vec2D[int]]bool)
	for _, r := range robots {
		robotMap[r.Pos] = true
	}

	// looking for something like that:
	/*
		1.1....................1..1..1..........................1............................................
		.........1......................................1......................................1.............
		...................................................................1.................................
		.........................1..1111111111111111111111111111111...............1...............1..........
		............................1.............................1...1......................................
		............................1.............................1..........................................
		............................1.............................1.......1.............1....................
		......................1.....1.............................1.....................1....................
		............................1..............1..............1......1....................1..............
		............................1.............111.............1..........................................
		.1..........................1............11111............1..............................1...........
		............................1...........1111111...........1...............1..........................
		............................1..........111111111..........1..........................................
		............................1............11111............1...................................1......
		............................1...........1111111...........1..........................................
		.............1..............1..........111111111..........1..........................................
		............................1.........11111111111.........1...1......................................
		............................1........1111111111111........1.........................................1
		............................1..........111111111..........1..........................................
		............................1.........11111111111.........1.1........................................
		............................1........1111111111111........1.........................1................
		............................1.......111111111111111.......1..........................................
		.....................1......1......11111111111111111......1..........................................
		............................1........1111111111111........1..........................................
		............................1.......111111111111111.......1......1...................................
		........................1...1......11111111111111111......1............................1.............
		......1.........1...........1.....1111111111111111111.....1..........1...............1...............
		............................1....111111111111111111111....1..........................................
		............................1.............111.............1..........................................
		............................1.............111.............1...................................1......
		............................1.............111.............1....................................1....1
		............................1.............................1......................1...................
		.1..........................1.............................1..1.......................................
		............................1.............................1..........................................
		........1...................1.............................1.......1..................................
		............................1111111111111111111111111111111..................1....1..................
		.............................................................................1.......................
		.....................................................................................................
	*/

	var count int
	for _, r := range robots {
		// count robots that are completely surrounded by other robots
		if robotMap[helper.Vec2D[int]{X: r.Pos.X + 1, Y: r.Pos.Y}] &&
			robotMap[helper.Vec2D[int]{X: r.Pos.X - 1, Y: r.Pos.Y}] &&
			robotMap[helper.Vec2D[int]{X: r.Pos.X, Y: r.Pos.Y + 1}] &&
			robotMap[helper.Vec2D[int]{X: r.Pos.X, Y: r.Pos.Y - 1}] {
			count++
		}
	}
	// 20 surrounded robots should be a safe hint:
	if count > 20 {
		printRobots(robots, areaSize)
		return true
	}
	return false
}
