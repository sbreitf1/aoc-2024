package main

// https://adventofcode.com/2024/day/12

import (
	"aoc/helper"
	"fmt"
)

func main() {
	lines := helper.ReadNonEmptyLines("example-4.txt")
	world := World(helper.LinesToRunes(lines))

	regions := world.GetRegions()

	solution1 := SumFencePrices1(regions)
	fmt.Println("-> part 1:", solution1)

	solution2 := SumFencePrices2(regions)
	fmt.Println("-> part 2:", solution2)
	// NOT 871596
}

type World [][]rune

func (w World) Width() int  { return len(w[0]) }
func (w World) Height() int { return len(w) }

type Region map[helper.Vec2D[int]]bool

func (w World) GetRegions() []Region {
	regions := make([]Region, 0)
	for y := range w {
		for x := range w[y] {
			p := helper.Vec2D[int]{X: x, Y: y}
			inArea := false
			for _, r := range regions {
				if r[p] {
					inArea = true
					break
				}
			}
			if !inArea {
				regions = append(regions, w.GetRegionAt(p))
			}
		}
	}
	return regions
}

func (w World) GetRegionAt(start helper.Vec2D[int]) Region {
	region := make(Region, 0)
	plant := w[start.Y][start.X]
	pointQueue := []helper.Vec2D[int]{start}
	for len(pointQueue) > 0 {
		p := pointQueue[len(pointQueue)-1]
		pointQueue = pointQueue[:len(pointQueue)-1]
		if region[p] {
			continue
		}
		region[p] = true

		for _, d := range []helper.Vec2D[int]{{X: 1, Y: 0}, {X: -1, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: -1}} {
			next := p.Add(d)
			if next.X >= 0 && next.Y >= 0 && next.X < w.Width() && next.Y < w.Height() {
				if w[next.Y][next.X] == plant {
					pointQueue = append(pointQueue, next)
				}
			}
		}
	}
	return region
}

func SumFencePrices1(regions []Region) int {
	var sum int
	for _, r := range regions {
		sum += r.FencePrice1()
	}
	return sum
}

func (r Region) FencePrice1() int {
	return r.Area() * r.Perimeter()
}

func (r Region) Area() int {
	return len(r)
}

func (r Region) Perimeter() int {
	var perimeter int
	for p := range r {
		for _, d := range []helper.Vec2D[int]{{X: 1, Y: 0}, {X: -1, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: -1}} {
			if !r[p.Add(d)] {
				perimeter++
			}
		}
	}
	return perimeter
}

func SumFencePrices2(regions []Region) int {
	var sum int
	for _, r := range regions {
		sum += r.FencePrice2()
	}
	return sum
}

func (r Region) FencePrice2() int {
	return r.Area() * r.SideCount()
}

func (r Region) SideCount() int {
	topLeft := r.FindTopLeftCorner()

	p := topLeft
	d := helper.Vec2D[int]{X: 1, Y: 0}

	sideCount := 1

	// assumption: the point left of p (=p+d.RotCCW) is never part of the region
	// traverse boundary of region clockwise
	for sideCount < 4 || p != topLeft {
		if r[p.Add(d)] {
			p = p.Add(d)
			if r[p.Add(d.RotCCW())] {
				sideCount++
				d = d.RotCCW()
				p = p.Add(d)
			}

		} else {
			sideCount++
			d = d.RotCW()
		}
	}

	return sideCount
}

func (r Region) FindTopLeftCorner() helper.Vec2D[int] {
	min, max := r.GetMinMax()
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			p := helper.Vec2D[int]{X: x, Y: y}
			if r[p] {
				return p
			}
		}
	}
	panic("not top-left corner found")
}

func (r Region) GetMinMax() (helper.Vec2D[int], helper.Vec2D[int]) {
	min := helper.Vec2D[int]{X: 1000000, Y: 1000000}
	max := helper.Vec2D[int]{X: 0, Y: 0}
	for p := range r {
		if p.X < min.X {
			min.X = p.X
		}
		if p.X > max.X {
			max.X = p.X
		}
		if p.Y < min.Y {
			min.Y = p.Y
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
	}
	return min, max
}
