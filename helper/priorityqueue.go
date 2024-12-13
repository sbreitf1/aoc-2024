package helper

import (
	"container/heap"
)

// inspired by https://pkg.go.dev/container/heap#example-package-PriorityQueue

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

type PriorityQueue[P Ordered, T any] struct {
	items priorityQueueItemList[P, T]
}

func NewPriorityQueue[P Ordered, T any]() *PriorityQueue[P, T] {
	return &PriorityQueue[P, T]{}
}

func (pq *PriorityQueue[P, T]) Push(priority P, obj T) {
	heap.Push(&pq.items, &priorityQueueItem[P, T]{Object: obj, Priority: priority})
}

// Pop removes and returns the element with lowest priority value.
func (pq *PriorityQueue[P, T]) Pop() (T, P) {
	item := heap.Pop(&pq.items).(*priorityQueueItem[P, T])
	return item.Object, item.Priority
}

func (pq *PriorityQueue[P, T]) Len() int {
	return pq.items.Len()
}

type priorityQueueItem[P Ordered, T any] struct {
	Object   T
	Priority P
	Index    int
}

type priorityQueueItemList[P Ordered, T any] []*priorityQueueItem[P, T]

func (pq priorityQueueItemList[P, T]) Len() int { return len(pq) }

func (pq priorityQueueItemList[P, T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq priorityQueueItemList[P, T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *priorityQueueItemList[P, T]) Push(x any) {
	n := len(*pq)
	item := x.(*priorityQueueItem[P, T])
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueueItemList[P, T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
