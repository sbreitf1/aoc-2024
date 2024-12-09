package main

// https://adventofcode.com/2024/day/6

import (
	"aoc/helper"
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	lines := helper.ReadLines("input.txt")
	area := parseArea(helper.LinesToRunes(lines))

	escapePath := area.GetEscapePath()
	solution1 := len(GetDistinctPos(escapePath))
	fmt.Println("-> part 1:", solution1)

	solution2 := area.CountLoops(escapePath)
	fmt.Println("-> part 2:", solution2)
}

type Area struct {
	fields   [][]rune
	guardPos helper.Vec2D[int]
	guardDir helper.Vec2D[int]
}

func (area Area) Clone() Area {
	return Area{
		fields:   helper.Clone(area.fields),
		guardPos: area.guardPos,
		guardDir: area.guardDir,
	}
}

func (area Area) IsObstacle(p helper.Vec2D[int]) bool {
	return area.fields[p.Y][p.X] == '#'
}

func parseArea(lines [][]rune) Area {
	var guardPos helper.Vec2D[int]
outterLoop:
	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == '^' {
				lines[y][x] = '.'
				guardPos = helper.Vec2D[int]{X: x, Y: y}
				break outterLoop
			}
		}
	}
	return Area{
		fields:   lines,
		guardPos: guardPos,
		guardDir: helper.Vec2D[int]{X: 0, Y: -1},
	}
}

func (area Area) GetEscapePath() []helper.Vec2D[int] {
	path := []helper.Vec2D[int]{area.guardPos}
	dir := area.guardDir
	for {
		newPos := path[len(path)-1].Add(dir)
		if !area.IsInside(newPos) {
			break
		}

		for area.IsObstacle(newPos) {
			dir = dir.RotCW()
			newPos = path[len(path)-1].Add(dir)
		}
		path = append(path, path[len(path)-1].Add(dir))
	}
	return path
}

func (area Area) IsInside(pos helper.Vec2D[int]) bool {
	return pos.X >= 0 && pos.X < len(area.fields[0]) && pos.Y >= 0 && pos.Y < len(area.fields)
}

func GetDistinctPos(path []helper.Vec2D[int]) []helper.Vec2D[int] {
	seen := make(map[helper.Vec2D[int]]bool)
	for _, p := range path {
		seen[p] = true
	}
	return helper.GetKeySlice(seen)
}

func (area Area) CountLoops(escapePath []helper.Vec2D[int]) int {
	candidates := GetDistinctPos(escapePath)

	var wg sync.WaitGroup
	var count int32
	for _, p := range candidates {
		wg.Add(1)
		go func(p helper.Vec2D[int]) {
			defer wg.Done()

			if area.fields[p.Y][p.X] == '.' && p != area.guardPos {
				tmpArea := area.Clone()
				tmpArea.fields[p.Y][p.X] = '#'
				if tmpArea.HasLoop() {
					atomic.AddInt32(&count, 1)
				}
			}
		}(p)
	}
	wg.Wait()
	return int(count)
}

func (area Area) HasLoop() bool {
	type GuardState struct {
		Pos, Dir helper.Vec2D[int]
	}
	guardState := GuardState{
		Pos: area.guardPos,
		Dir: area.guardDir,
	}

	seenStates := make(map[GuardState]bool)
	for {
		if _, ok := seenStates[guardState]; ok {
			return true
		}
		seenStates[guardState] = true

		newPos := guardState.Pos.Add(guardState.Dir)
		if !area.IsInside(newPos) {
			return false
		}

		for area.IsObstacle(newPos) {
			guardState.Dir = guardState.Dir.RotCW()
			newPos = guardState.Pos.Add(guardState.Dir)
		}
		guardState.Pos = guardState.Pos.Add(guardState.Dir)
	}
}
