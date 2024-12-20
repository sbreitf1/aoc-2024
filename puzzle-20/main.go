package main

// https://adventofcode.com/2024/day/20

import (
	"aoc/helper"
	"aoc/helper/dijkstra"
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")
	maze := ParseMaze(lines)

	solution1 := CountWorthwileCheatPositions(maze, 100)
	fmt.Println("-> part 1:", solution1)

	solution2 := 0
	fmt.Println("-> part 2:", solution2)
}

type Maze struct {
	Field      [][]rune
	Start, End helper.Vec2D[int]
}

func ParseMaze(lines []string) Maze {
	field := helper.LinesToRunes(lines)
	var start, end helper.Vec2D[int]
	for y := range field {
		for x := range field[y] {
			if field[y][x] == 'S' {
				start = helper.NewVec2D(x, y)
			} else if field[y][x] == 'E' {
				end = helper.NewVec2D(x, y)
			}
		}
	}
	return Maze{
		Field: field,
		Start: start,
		End:   end,
	}
}

func GetPathDist(maze Maze) int {
	_, noCheatScore := dijkstra.MustFindPath(maze.Start, maze.End, dijkstra.Params[int, helper.Vec2D[int]]{
		SuccessorGenerator: dijkstra.NewDefaultFieldSuccessorGenerator(maze.Field, []rune{'.', 'S', 'E'}, []rune{'#'}),
	})
	return noCheatScore
}

func CountWorthwileCheatPositions(maze Maze, minSaving int) int {
	noCheatScore := GetPathDist(maze)

	var wg sync.WaitGroup
	var doneLines, count int32
	for y := range maze.Field {
		wg.Add(1)
		go func(maze Maze, y int) {
			defer wg.Done()
			for x := range maze.Field[y] {
				p := helper.NewVec2D(x, y)
				if maze.Field[p.Y][p.X] == '#' {
					maze.Field[p.Y][p.X] = '.'

					cheatScore := GetPathDist(maze)

					if (noCheatScore - cheatScore) >= minSaving {
						atomic.AddInt32(&count, 1)
					}

					maze.Field[p.Y][p.X] = '#'
				}
			}
			atomic.AddInt32(&doneLines, 1)
			fmt.Println(doneLines, "of", len(maze.Field), "done")
		}(helper.Clone(maze), y)
	}
	wg.Wait()
	return int(count)
}
