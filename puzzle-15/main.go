package main

// https://adventofcode.com/2024/day/15

import (
	"aoc/helper"
	"fmt"
	"strings"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")
	level, moves := ParseLevelAndMoves(lines)

	level2 := level.Widen()

	level.MoveSequence(moves)
	solution1 := level.ComputeGPSCoordinates()
	fmt.Println("-> part 1:", solution1)

	level2.MoveSequence(moves)
	solution2 := level2.ComputeGPSCoordinates()
	fmt.Println("-> part 2:", solution2)
}

type Level struct {
	Width, Height int
	Barriers      map[helper.Vec2D[int]]bool
	Boxes         []Box
	RobotPos      helper.Vec2D[int]
}

type Box struct {
	PosL   helper.Vec2D[int]
	IsWide bool
}

func (b Box) Width() int {
	if b.IsWide {
		return 2
	}
	return 1
}

func (l *Level) Print() {
	runes := helper.InitSlice2D(l.Width, l.Height, '.')
	for p := range l.Barriers {
		runes[p.Y][p.X] = '#'
	}
	for _, b := range l.Boxes {
		if b.IsWide {
			runes[b.PosL.Y][b.PosL.X] = '['
			runes[b.PosL.Y][b.PosL.X+1] = ']'
		} else {
			runes[b.PosL.Y][b.PosL.X] = 'O'
		}
	}
	runes[l.RobotPos.Y][l.RobotPos.X] = '@'
	for y := range runes {
		fmt.Println(string(runes[y]))
	}
	fmt.Println(l.RobotPos)
}

func ParseLevelAndMoves(lines []string) (Level, []rune) {
	var w, h int
	barriers := make(map[helper.Vec2D[int]]bool)
	boxes := make([]Box, 0)
	var robotPos helper.Vec2D[int]
	moves := make([]rune, 0)
	for y, l := range lines {
		if strings.HasPrefix(l, "#") {
			h++
			w = len(l)
			for x, r := range l {
				p := helper.Vec2D[int]{X: x, Y: y}
				if r == '#' {
					barriers[p] = true
				} else if r == 'O' {
					boxes = append(boxes, Box{PosL: p, IsWide: false})
				} else if r == '@' {
					robotPos = p
				}
			}
		} else {
			moves = append(moves, []rune(strings.TrimSpace(l))...)
		}
	}
	return Level{
		Width:    w,
		Height:   h,
		Barriers: barriers,
		Boxes:    boxes,
		RobotPos: robotPos,
	}, moves
}

func (l *Level) MoveSequence(moves []rune) {
	for _, r := range moves {
		switch r {
		case '<':
			l.MoveX(-1)
		case '>':
			l.MoveX(1)
		case '^':
			l.MoveY(-1)
		case 'v':
			l.MoveY(1)
		}
	}
}

func (l *Level) MoveX(dirX int) bool {
	dir := helper.Vec2D[int]{X: dirX, Y: 0}

	pushBoxes := make([]int, 0)
	p := l.RobotPos.Add(dir)
	for {
		if l.Barriers[p] {
			return false
		}
		boxIndex := l.GetBoxIndexAt(p)
		if boxIndex == -1 {
			break
		}
		pushBoxes = append(pushBoxes, boxIndex)
		p = p.Add(dir.Mul(l.Boxes[boxIndex].Width()))
	}

	for _, boxIndex := range pushBoxes {
		l.Boxes[boxIndex].PosL = l.Boxes[boxIndex].PosL.Add(dir)
	}
	l.RobotPos = l.RobotPos.Add(dir)

	return true
}

func (l *Level) MoveY(dirY int) bool {
	dir := helper.Vec2D[int]{X: 0, Y: dirY}

	p := l.RobotPos.Add(dir)
	if l.Barriers[p] {
		return false
	}

	boxIndex := l.GetBoxIndexAt(p)
	if boxIndex >= 0 {
		if !l.boxCanMoveY(boxIndex, dirY) {
			return false
		}
		l.boxMoveY(boxIndex, dirY)
	}

	l.RobotPos = p

	return true
}

func (l *Level) boxCanMoveY(boxIndex, dirY int) bool {
	dir := helper.Vec2D[int]{X: 0, Y: dirY}
	p := l.Boxes[boxIndex].PosL.Add(dir)
	if l.Barriers[p] {
		return false
	}
	if pushedBoxIndex := l.GetBoxIndexAt(p); pushedBoxIndex >= 0 {
		if !l.boxCanMoveY(pushedBoxIndex, dirY) {
			return false
		}
	}
	if l.Boxes[boxIndex].IsWide {
		p := l.Boxes[boxIndex].PosL.Add(dir).Add(helper.Vec2D[int]{X: 1, Y: 0})
		if l.Barriers[p] {
			return false
		}
		if pushedBoxIndex := l.GetBoxIndexAt(p); pushedBoxIndex >= 0 {
			if !l.boxCanMoveY(pushedBoxIndex, dirY) {
				return false
			}
		}
	}
	return true
}

func (l *Level) boxMoveY(boxIndex, dirY int) bool {
	dir := helper.Vec2D[int]{X: 0, Y: dirY}
	p := l.Boxes[boxIndex].PosL.Add(dir)
	if pushedBoxIndex := l.GetBoxIndexAt(p); pushedBoxIndex >= 0 {
		l.boxMoveY(pushedBoxIndex, dirY)
	}
	if l.Boxes[boxIndex].IsWide {
		p := l.Boxes[boxIndex].PosL.Add(dir).Add(helper.Vec2D[int]{X: 1, Y: 0})
		if pushedBoxIndex := l.GetBoxIndexAt(p); pushedBoxIndex >= 0 {
			// no double move if boxes align -> already moved by left index
			l.boxMoveY(pushedBoxIndex, dirY)
		}
	}
	l.Boxes[boxIndex].PosL = l.Boxes[boxIndex].PosL.Add(dir)
	return true
}

func (l *Level) GetBoxIndexAt(p helper.Vec2D[int]) int {
	for i, b := range l.Boxes {
		if b.PosL == p {
			return i
		}
		if b.IsWide && b.PosL.Add(helper.Vec2D[int]{X: 1, Y: 0}) == p {
			return i
		}
	}
	return -1
}

func (l *Level) ComputeGPSCoordinates() int {
	var sum int
	for _, b := range l.Boxes {
		sum += b.PosL.X + 100*b.PosL.Y
	}
	return sum
}

func (l *Level) Widen() Level {
	barriers := make(map[helper.Vec2D[int]]bool, 2*len(l.Barriers))
	boxes := make([]Box, len(l.Boxes))
	for p := range l.Barriers {
		p := helper.Vec2D[int]{X: 2 * p.X, Y: p.Y}
		barriers[p] = true
		barriers[p.Add(helper.Vec2D[int]{X: 1, Y: 0})] = true
	}
	for _, b := range l.Boxes {
		p := helper.Vec2D[int]{X: 2 * b.PosL.X, Y: b.PosL.Y}
		boxes = append(boxes, Box{PosL: p, IsWide: true})
	}
	return Level{
		Width:    2 * l.Width,
		Height:   l.Height,
		Barriers: barriers,
		Boxes:    boxes,
		RobotPos: helper.Vec2D[int]{X: 2 * l.RobotPos.X, Y: l.RobotPos.Y},
	}
}
