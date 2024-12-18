package main

// https://adventofcode.com/2024/day/18

import (
	"aoc/helper"
	"fmt"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")
	snowflakes := ParseSnowflakes(lines)
	size, steps := getParams(snowflakes)

	mem := NewMemory(size, snowflakes, steps)
	path, ok := mem.FindPath(helper.NewVec2D(0, 0), helper.NewVec2D(size.X-1, size.Y-1))
	if !ok {
		helper.ExitWithMessage("no path found!")
	}
	solution1 := len(path) - 1
	fmt.Println("-> part 1:", solution1)

	solution2 := findFirstSnowflakeThatBlocks(helper.NewVec2D(0, 0), helper.NewVec2D(size.X-1, size.Y-1), size, snowflakes)
	fmt.Println("-> part 2:", fmt.Sprintf("%d,%d", solution2.X, solution2.Y))
}

func ParseSnowflakes(lines []string) []helper.Vec2D[int] {
	snowflakes := make([]helper.Vec2D[int], 0)
	for _, l := range lines {
		ints := helper.ExtractInts[int](l)
		if len(ints) == 2 {
			snowflakes = append(snowflakes, helper.NewVec2D(ints[0], ints[1]))
		}
	}
	return snowflakes
}

func getParams(snowflakes []helper.Vec2D[int]) (helper.Vec2D[int], int) {
	for _, s := range snowflakes {
		if s.X > 6 || s.Y > 6 {
			return helper.NewVec2D(71, 71), 1024
		}
	}
	return helper.NewVec2D(7, 7), 12
}

type Memory struct {
	Width, Height int
	Fields        [][]rune
}

func (mem Memory) Print(Path []helper.Vec2D[int]) {
	runes := helper.Clone(mem.Fields)
	for _, p := range Path {
		runes[p.Y][p.X] = 'O'
	}
	for _, l := range helper.RunesToLines(runes) {
		fmt.Println(l)
	}
}

func NewMemory(size helper.Vec2D[int], snowflakes []helper.Vec2D[int], steps int) Memory {
	fields := helper.InitSlice2D(size.Y, size.X, '.')
	for i := 0; i < steps; i++ {
		fields[snowflakes[i].Y][snowflakes[i].X] = '#'
	}
	return Memory{
		Width:  len(fields[0]),
		Height: len(fields),
		Fields: fields,
	}
}

func (mem Memory) FindPath(from, to helper.Vec2D[int]) ([]helper.Vec2D[int], bool) {
	type Crumb struct {
		Pos      helper.Vec2D[int]
		IsStart  bool
		Previous helper.Vec2D[int]
	}
	queue := helper.NewPriorityQueue[int, Crumb]()
	queue.Push(0, Crumb{Pos: from, IsStart: true})
	seen := make(map[helper.Vec2D[int]]Crumb)
	for queue.Len() > 0 {
		c, pathLen := queue.Pop()
		if c.Pos == to {
			seen[c.Pos] = c
			path := make([]helper.Vec2D[int], 0)
			p := c.Pos
			for {
				crumb := seen[p]
				path = append(path, crumb.Pos)
				if crumb.IsStart {
					break
				}
				p = crumb.Previous
			}
			helper.ReverseSlice(path)
			return path, true
		}
		if _, ok := seen[c.Pos]; ok {
			continue
		}
		seen[c.Pos] = c

		for _, dir := range []helper.Vec2D[int]{helper.NewVec2D(1, 0), helper.NewVec2D(-1, 0), helper.NewVec2D(0, 1), helper.NewVec2D(0, -1)} {
			p := c.Pos.Add(dir)
			if p.X >= 0 && p.Y >= 0 && p.X < mem.Width && p.Y < mem.Height {
				if mem.Fields[p.Y][p.X] != '#' {
					queue.Push(pathLen+1, Crumb{Pos: p, Previous: c.Pos})
				}
			}
		}
	}

	return nil, false
}

func findFirstSnowflakeThatBlocks(from, to helper.Vec2D[int], size helper.Vec2D[int], snowflakes []helper.Vec2D[int]) helper.Vec2D[int] {
	cache := make(map[int]bool)
	hasPath := func(steps int) bool {
		if val, ok := cache[steps]; ok {
			return val
		}
		mem := NewMemory(size, snowflakes, steps+1)
		_, ok := mem.FindPath(from, to)
		cache[steps] = ok
		return ok
	}

	min := 0
	max := len(snowflakes) - 1
	for {
		if (max - min) <= 1 {
			break
		}
		steps := (min + max) / 2
		if hasPath(steps) {
			min = steps
		} else {
			max = steps
		}
	}

	if hasPath(min-1) && !hasPath(min) {
		return snowflakes[min]
	}
	if hasPath(max-1) && !hasPath(max) {
		return snowflakes[max]
	}

	helper.ExitWithMessage("no snowflake is blocking!")
	return helper.NewVec2D(0, 0)
}
