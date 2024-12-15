package main

// https://adventofcode.com/2024/day/15

import (
	"aoc/helper"
	"fmt"
	"strings"
)

func main() {
	lines := helper.ReadNonEmptyLines("example-1.txt")
	level, moves := ParseLevelAndMoves(lines)

	level.MoveSequence(moves)

	solution1 := level.ComputeGPSCoordinates()
	fmt.Println("-> part 1:", solution1)
}

type Level struct {
	Fields   [][]rune
	RobotPos helper.Vec2D[int]
}

func (l *Level) Print() {
	for y := range l.Fields {
		fmt.Println(string(l.Fields[y]))
	}
	fmt.Println(l.RobotPos)
}

func ParseLevelAndMoves(lines []string) (Level, []rune) {
	fields := make([][]rune, 0)
	var robotPos helper.Vec2D[int]
	moves := make([]rune, 0)
	for y, l := range lines {
		if strings.HasPrefix(l, "#") {
			fields = append(fields, []rune(strings.TrimSpace(l)))
			for x := range l {
				if l[x] == '@' {
					robotPos = helper.Vec2D[int]{X: x, Y: y}
				}
			}
		} else {
			moves = append(moves, []rune(strings.TrimSpace(l))...)
		}
	}
	return Level{Fields: fields, RobotPos: robotPos}, moves
}

func (l *Level) MoveSequence(moves []rune) {
	dirs := map[rune]helper.Vec2D[int]{
		'<': {X: -1, Y: 0},
		'>': {X: 1, Y: 0},
		'^': {X: 0, Y: -1},
		'v': {X: 0, Y: 1},
	}

	for _, r := range moves {
		if d, ok := dirs[r]; ok {
			l.Move(d)
		}
	}
}

func (l *Level) Move(dir helper.Vec2D[int]) bool {
	p := l.RobotPos.Add(dir)
	for ; l.Fields[p.Y][p.X] == 'O'; p = p.Add(dir) {
	}
	if l.Fields[p.Y][p.X] == '#' {
		return false
	}
	dist := l.RobotPos.Dist(p)

	for i := dist; i > 0; i-- {
		pSrc := p.Sub(dir)
		l.Fields[p.Y][p.X] = l.Fields[pSrc.Y][pSrc.X]
		p = pSrc
	}
	l.Fields[p.Y][p.X] = '.'
	l.RobotPos = p.Add(dir)

	return true
}

func (l *Level) ComputeGPSCoordinates() int {
	var sum int
	for y := range l.Fields {
		for x := range l.Fields[y] {
			if l.Fields[y][x] == 'O' {
				sum += 100*y + x
			}
		}
	}
	return sum
}
