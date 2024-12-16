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

	solution2 := level.CountAllBestPathTiles(solution1)
	fmt.Println("-> part 2:", solution2)
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
				start = helper.NewVec2D(x, y)
			} else if fields[y][x] == 'E' {
				end = helper.NewVec2D(x, y)
			}
		}
	}
	return Level{
		Fields:   fields,
		Start:    start,
		StartDir: helper.NewVec2D(1, 0),
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
		if c.Score > maxScore {
			break
		}

		key := Key{Pos: c.Pos, Dir: c.Dir}
		if s, ok := seen[key]; ok {
			if c.Score < s.Score {
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
			continue
		}

		{
			d := c.Dir
			p := c.Pos.Add(d)
			if l.Fields[p.Y][p.X] != '#' {
				next := Crumb{Pos: p, Dir: d, Score: c.Score + 1, Previous: key}
				queue.Push(next.Score, next)
			}
		}
		{
			d := c.Dir.RotCW()
			p := c.Pos.Add(d)
			if l.Fields[p.Y][p.X] != '#' {
				next := Crumb{Pos: p, Dir: d, Score: c.Score + 1001, Previous: key}
				queue.Push(next.Score, next)
			}
		}
		{
			d := c.Dir.RotCCW()
			p := c.Pos.Add(d)
			if l.Fields[p.Y][p.X] != '#' {
				next := Crumb{Pos: p, Dir: d, Score: c.Score + 1001, Previous: key}
				queue.Push(next.Score, next)
			}
		}
	}

	keyQueue := []Key{
		{Pos: l.End, Dir: helper.Vec2D[int]{X: 1, Y: 0}},
		{Pos: l.End, Dir: helper.Vec2D[int]{X: -1, Y: 0}},
		{Pos: l.End, Dir: helper.Vec2D[int]{X: 0, Y: 1}},
		{Pos: l.End, Dir: helper.Vec2D[int]{X: 0, Y: -1}},
	}
	distinctPos := make(map[helper.Vec2D[int]]bool)
	for len(keyQueue) > 0 {
		k := keyQueue[0]
		keyQueue = keyQueue[1:]
		if s, ok := seen[k]; ok {
			distinctPos[k.Pos] = true
			keyQueue = append(keyQueue, s.Parents...)
		}
	}
	return len(distinctPos)
}
