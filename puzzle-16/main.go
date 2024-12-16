package main

// https://adventofcode.com/2024/day/16

import (
	"aoc/helper"
	"fmt"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")
	level := ParseLevel(lines)

	solution1 := level.FindBestPathScore()
	fmt.Println("-> part 1:", solution1)

	fmt.Println(level.CountAllBestPathTiles(solution1))
}

type Level struct {
	Fields     [][]rune
	Start, End helper.Vec2D[int]
	StartDir   helper.Vec2D[int]
}

func ParseLevel(lines []string) Level {
	fields := helper.LinesToRunes(lines)
	var start, end helper.Vec2D[int]
	for y := range fields {
		for x := range fields {
			if fields[y][x] == 'S' {
				start = helper.Vec2D[int]{X: x, Y: y}
			} else if fields[y][x] == 'E' {
				end = helper.Vec2D[int]{X: x, Y: y}
			}
		}
	}
	return Level{
		Fields:   fields,
		Start:    start,
		StartDir: helper.Vec2D[int]{X: 1, Y: 0},
		End:      end,
	}
}

func (l Level) Print(path []helper.Vec2D[int]) {
	runes := helper.Clone(l.Fields)
	for _, p := range path {
		runes[p.Y][p.X] = 'X'
	}
	for y := range runes {
		fmt.Println(string(runes[y]))
	}
}

func (l Level) FindBestPathScore() int {
	type Crumb struct {
		Pos   helper.Vec2D[int]
		Dir   helper.Vec2D[int]
		Score int
		Path  []helper.Vec2D[int]
	}
	type Key struct {
		Pos, Dir helper.Vec2D[int]
	}

	queue := helper.NewPriorityQueue[int, Crumb]()
	queue.Push(0, Crumb{Pos: l.Start, Dir: l.StartDir, Score: 0, Path: []helper.Vec2D[int]{l.Start}})
	seen := make(map[Key]bool)
	for queue.Len() > 0 {
		c, _ := queue.Pop()
		key := Key{Pos: c.Pos, Dir: c.Dir}
		if seen[key] {
			continue
		}
		seen[key] = true

		if c.Pos == l.End {
			return c.Score
		}

		{
			d := c.Dir
			p := c.Pos.Add(d)
			if l.Fields[p.Y][p.X] != '#' {
				next := Crumb{Pos: p, Dir: d, Score: c.Score + 1, Path: helper.Combine(c.Path, p)}
				queue.Push(next.Score, next)
			}
		}
		{
			d := c.Dir.RotCW()
			p := c.Pos.Add(d)
			if l.Fields[p.Y][p.X] != '#' {
				next := Crumb{Pos: p, Dir: d, Score: c.Score + 1001, Path: helper.Combine(c.Path, p)}
				queue.Push(next.Score, next)
			}
		}
		{
			d := c.Dir.RotCCW()
			p := c.Pos.Add(d)
			if l.Fields[p.Y][p.X] != '#' {
				next := Crumb{Pos: p, Dir: d, Score: c.Score + 1001, Path: helper.Combine(c.Path, p)}
				queue.Push(next.Score, next)
			}
		}
	}

	helper.ExitWithMessage("no path found!")
	return -1
}

func (l Level) CountAllBestPathTiles(maxScore int) int {
	type Key struct {
		Pos, Dir helper.Vec2D[int]
	}
	type Crumb struct {
		Pos      helper.Vec2D[int]
		Dir      helper.Vec2D[int]
		Score    int
		IsStart  bool
		Previous Key
	}
	type Seen struct {
		Score   int
		Parents []Key
	}

	queue := helper.NewPriorityQueue[int, Crumb]()
	queue.Push(0, Crumb{Pos: l.Start, Dir: l.StartDir, Score: 0, IsStart: true})
	seen := make(map[Key]Seen)
	for queue.Len() > 0 {
		c, _ := queue.Pop()
		key := Key{Pos: c.Pos, Dir: c.Dir}
		if s, ok := seen[key]; ok {
			if s.Score < c.Score {
				seen[key] = Seen{
					Score:   c.Score,
					Parents: []Key{c.Previous},
				}
			} else if s.Score == c.Score {
				seen[key] = Seen{
					Score:   c.Score,
					Parents: append(seen[key].Parents, c.Previous),
				}
			}
			continue
		}
		if c.IsStart {
			seen[key] = Seen{Score: c.Score}
		} else {
			seen[key] = Seen{
				Score:   c.Score,
				Parents: []Key{c.Previous},
			}
		}

		if c.Pos == l.End {
			return c.Score
		}

		{
			d := c.Dir
			p := c.Pos.Add(d)
			if l.Fields[p.Y][p.X] != '#' {
				next := Crumb{Pos: p, Dir: d, Score: c.Score + 1, Path: helper.Combine(c.Path, p)}
				queue.Push(next.Score, next)
			}
		}
		{
			d := c.Dir.RotCW()
			p := c.Pos.Add(d)
			if l.Fields[p.Y][p.X] != '#' {
				next := Crumb{Pos: p, Dir: d, Score: c.Score + 1001, Path: helper.Combine(c.Path, p)}
				queue.Push(next.Score, next)
			}
		}
		{
			d := c.Dir.RotCCW()
			p := c.Pos.Add(d)
			if l.Fields[p.Y][p.X] != '#' {
				next := Crumb{Pos: p, Dir: d, Score: c.Score + 1001, Path: helper.Combine(c.Path, p)}
				queue.Push(next.Score, next)
			}
		}
	}

	helper.ExitWithMessage("no path found!")
	return -1
}
