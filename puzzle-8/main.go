package main

// https://adventofcode.com/2024/day/8

import (
	"aoc/helper"
	"fmt"
)

func main() {
	lines := helper.ReadLines("input.txt")
	world := parseWorld(helper.LinesToRunes(lines))

	antinodes1 := world.GetAntinodes(1, 1)
	solution1 := countDistinctPos(antinodes1)
	fmt.Println("-> part 1:", solution1)

	antinodes2 := world.GetAntinodes(0, helper.Max(world.Width, world.Height))
	solution2 := countDistinctPos(antinodes2)
	fmt.Println("-> part 2:", solution2)
}

type World struct {
	Width, Height int
	Antennas      map[rune][]helper.Vec2D[int]
}

func (w World) Contains(p helper.Vec2D[int]) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < w.Width && p.Y < w.Height
}

func parseWorld(lines [][]rune) World {
	antennas := make(map[rune][]helper.Vec2D[int])
	for y := range lines {
		for x := range lines[y] {
			r := lines[y][x]
			if r != '.' {
				if _, ok := antennas[r]; !ok {
					antennas[r] = make([]helper.Vec2D[int], 0)
				}
				antennas[r] = append(antennas[r], helper.Vec2D[int]{X: x, Y: y})
			}
		}
	}
	return World{
		Width:    len(lines[0]),
		Height:   len(lines),
		Antennas: antennas,
	}
}

type Antinode struct {
	Frequency rune
	Pos       helper.Vec2D[int]
}

func (w World) GetAntinodes(min, max int) []Antinode {
	antinodes := make([]Antinode, 0)
	for f := range w.Antennas {
		for i := 0; i < len(w.Antennas[f])-1; i++ {
			for j := i + 1; j < len(w.Antennas[f]); j++ {
				a1 := w.Antennas[f][i]
				a2 := w.Antennas[f][j]
				dir := a2.Sub(a1)
				antinodes = append(antinodes, w.GenerateAntinodes(f, a1, dir.Neg(), min, max)...)
				antinodes = append(antinodes, w.GenerateAntinodes(f, a2, dir, min, max)...)
			}
		}
	}
	return antinodes
}

func (w World) GenerateAntinodes(frequency rune, p, d helper.Vec2D[int], min, max int) []Antinode {
	antinodes := make([]Antinode, 0)
	for i := min; i <= max; i++ {
		ap := p.Add(d.Mul(i))
		if w.Contains(ap) {
			antinodes = append(antinodes, Antinode{
				Frequency: frequency,
				Pos:       ap,
			})
		}
	}
	return antinodes
}

func countDistinctPos(antinodes []Antinode) int {
	pos := make(map[helper.Vec2D[int]]bool)
	for _, a := range antinodes {
		pos[a.Pos] = true
	}
	return len(pos)
}
