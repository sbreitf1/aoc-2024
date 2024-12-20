package dijkstra

import "aoc/helper"

type Successor[D helper.Number, T comparable] struct {
	Obj  T
	Dist D
}

type Params[D helper.Number, T comparable] struct {
	SuccessorGenerator func(current T, currentDist D) []Successor[D, T]
	Equals             func(obj1, obj2 T) bool
}

func MustFindPath[D helper.Number, T comparable](from, to T, params Params[D, T]) ([]T, D) {
	path, dist, ok := FindPath(from, to, params)
	if !ok {
		helper.ExitWithMessage("no path found!")
	}
	return path, dist
}

func FindPath[D helper.Number, T comparable](from, to T, params Params[D, T]) ([]T, D, bool) {
	if params.SuccessorGenerator == nil {
		panic("params.SuccessorGenerator must be set")
	}
	if params.Equals == nil {
		params.Equals = func(obj1, obj2 T) bool { return obj1 == obj2 }
	}

	type Crumb struct {
		Obj      T
		Previous *Crumb
	}

	queue := helper.NewPriorityQueue[D, Crumb]()
	queue.Push(0, Crumb{Obj: from, Previous: nil})
	seen := make(map[T]Crumb)
	for queue.Len() > 0 {
		c, dist := queue.Pop()
		if params.Equals(c.Obj, to) {
			seen[c.Obj] = c
			path := make([]T, 0)
			cur := c.Obj
			for {
				crumb := seen[cur]
				path = append(path, crumb.Obj)
				if crumb.Previous == nil {
					break
				}
				cur = crumb.Previous.Obj
			}
			helper.ReverseSlice(path)
			return path, dist, true
		}

		if _, ok := seen[c.Obj]; ok {
			continue
		}
		seen[c.Obj] = c

		for _, obj := range params.SuccessorGenerator(c.Obj, dist) {
			if _, ok := seen[obj.Obj]; !ok {
				queue.Push(obj.Dist, Crumb{Obj: obj.Obj, Previous: &c})
			}
		}
	}

	return nil, 0, false
}

func NewDefaultFieldSuccessorGenerator(field [][]rune, floorRunes, wallRunes []rune) func(current helper.Vec2D[int], currentDist int) []Successor[int, helper.Vec2D[int]] {
	wallMap := make(map[rune]bool)
	for _, r := range wallRunes {
		wallMap[r] = true
	}

	return func(current helper.Vec2D[int], currentDist int) []Successor[int, helper.Vec2D[int]] {
		successors := make([]Successor[int, helper.Vec2D[int]], 0)
		for _, dir := range []helper.Vec2D[int]{helper.NewVec2D(-1, 0), helper.NewVec2D(1, 0), helper.NewVec2D(0, -1), helper.NewVec2D(0, 1)} {
			p := current.Add(dir)
			if p.X >= 0 && p.Y >= 0 && p.X < len(field[0]) && p.Y < len(field) {
				if !wallMap[field[p.Y][p.X]] {
					successors = append(successors, Successor[int, helper.Vec2D[int]]{
						Obj:  p,
						Dist: currentDist + 1,
					})
				}
			}
		}
		return successors
	}
}
