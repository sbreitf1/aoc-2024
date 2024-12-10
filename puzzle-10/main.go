package main

// https://adventofcode.com/2024/day/10

import (
	"aoc/helper"
	"fmt"
)

func main() {
	str := helper.ReadNonEmptyLines("input.txt")
	world := parseWorld(str)

	solution1 := world.SumAllTrailheadScores()
	fmt.Println("-> part 1:", solution1)

	solution2 := world.CountAllTrailheadPaths()
	fmt.Println("-> part 2:", solution2)
}

type World struct {
	Width, Height int
	Fields        [][]int
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
		Width:  len(fields[0]),
		Height: len(fields),
		Fields: fields,
	}
}

func (w World) GetTrailheads() []helper.Vec2D[int] {
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

func (w World) FindAllPathsFrom(trailhead helper.Vec2D[int]) []Path {
	return w.gatherAllPathsFrom(Path{trailhead})
}

func (w World) gatherAllPathsFrom(currentPath Path) []Path {
	currentPos := currentPath[len(currentPath)-1]
	currentVal := w.Fields[currentPos.Y][currentPos.X]
	if currentVal == 9 {
		return []Path{helper.Clone(currentPath)}
	}

	paths := make([]Path, 0)
	nextPath := make(Path, len(currentPath)+1)
	copy(nextPath[:len(currentPath)], currentPath)
	for _, d := range []helper.Vec2D[int]{{X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}} {
		nextPos := currentPos.Add(d)
		if nextPos.X >= 0 && nextPos.Y >= 0 && nextPos.X < w.Width && nextPos.Y < w.Height {
			if w.Fields[nextPos.Y][nextPos.X] == currentVal+1 {
				nextPath[len(nextPath)-1] = currentPos.Add(d)
				paths = append(paths, w.gatherAllPathsFrom(nextPath)...)
			}
		}
	}
	return paths
}

func (w World) SumAllTrailheadScores() int {
	trailheads := w.GetTrailheads()
	var sum int
	for _, p := range trailheads {
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

func (w World) CountAllTrailheadPaths() int {
	trailheads := w.GetTrailheads()
	var count int
	for _, p := range trailheads {
		paths := w.FindAllPathsFrom(p)
		count += len(paths)
	}
	return count
}
