package dijkstra

import "aoc/helper"

type Successor[S helper.Number, T comparable] struct {
	Obj  T
	Dist S
}

func FindPath[S helper.Number, T comparable](from, to T, next func(current T, currentDist S) []Successor[S, T]) ([]T, S, bool) {
	type Crumb struct {
		Obj      T
		Previous *Crumb
	}

	queue := helper.NewPriorityQueue[S, Crumb]()
	queue.Push(0, Crumb{Obj: from, Previous: nil})
	seen := make(map[T]Crumb)
	for queue.Len() > 0 {
		c, dist := queue.Pop()
		if c.Obj == to {
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

		nextObjs := next(c.Obj, dist)
		for _, obj := range nextObjs {
			queue.Push(obj.Dist, Crumb{Obj: obj.Obj, Previous: &c})
		}
	}

	return nil, 0, false
}
