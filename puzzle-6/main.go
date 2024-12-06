package main

// https://adventofcode.com/2024/day/6

import (
	"aoc/helper"
	"fmt"
)

func main() {
	lines := helper.ReadLines("input.txt")
	area := parseArea(helper.LinesToRunes(lines))

	escapePath := area.GetEscapePath()
	solution1 := len(GetDistinctPos(escapePath))
	fmt.Println("-> part 1:", solution1)

	solution2 := area.CountLoops(escapePath)
	fmt.Println("-> part 2:", solution2, "(possibly wrong!)")
	// is NOT 1564 !
	// and NOT 1563 !
	// and NOT 1565 !
}

type Area struct {
	fields   [][]rune
	guardPos helper.Vec2D[int]
	guardDir helper.Vec2D[int]
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

		if area.IsObstacle(newPos) {
			dir = dir.RotCW()
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

	var count int
	for _, p := range candidates {
		if area.fields[p.Y][p.X] == '.' && p != area.guardPos {
			area.fields[p.Y][p.X] = '#'
			if area.HasLoop() {
				count++
			}
			area.fields[p.Y][p.X] = '.'
		}
	}
	return count
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

		if area.IsObstacle(newPos) {
			guardState.Dir = guardState.Dir.RotCW()
		}
		guardState.Pos = guardState.Pos.Add(guardState.Dir)
	}
}
