package main

// https://adventofcode.com/2024/day/10

import (
	"aoc/helper"
	"fmt"
)

func main() {
	str := helper.ReadNonEmptyLines("example-2.txt")
	world := parseWorld(str)

	solution1 := world.SumAllTrailheadScores()
	fmt.Println("-> part 1:", solution1)
}

type World struct {
	Fields [][]int
}

func parseWorld(lines []string) World {
	fields := make([][]int, len(lines))
	for y := range lines {
		fields[y] = make([]int, len(lines[y]))
		for x := range lines[y] {
			fields[y][x] = int(lines[y][x] - '0')
		}
	}
	return World{
		Fields: fields,
	}
}

func (w World) GetAllZeroPos() []helper.Vec2D[int] {
	zeroPos := make([]helper.Vec2D[int], 0)
	for y := range w.Fields {
		for x := range w.Fields[y] {
			if w.Fields[y][x] == 0 {
				zeroPos = append(zeroPos, helper.Vec2D[int]{X: x, Y: y})
			}
		}
	}
	return zeroPos
}

type Path []helper.Vec2D[int]

func (w World) FindAllPathsFrom(start helper.Vec2D[int]) []Path {
	return nil
}

func (w World) SumAllTrailheadScores() int {
	zeroPos := w.GetAllZeroPos()
	var sum int
	for _, p := range zeroPos {
		paths := w.FindAllPathsFrom(p)
		sum += CountDistinctEndPos(paths)
	}
	return sum
}

func CountDistinctEndPos(paths []Path) int {
	endPosMap := make(map[helper.Vec2D[int]]bool)
	for _, path := range paths {
		endPosMap[path[len(path)-1]] = true
	}
	return len(endPosMap)
}
