package model

import "container/heap"

type GetheringPriorityQueue [][]Gethering

func (pq GetheringPriorityQueue) Len() int { return len(pq) }

func (pq GetheringPriorityQueue) Less(i, j int) bool {
	// Pop empty lists first, to reduce computation complexity
	if len(pq[i]) == 0 {
		return true
	}
	if len(pq[j]) == 0 {
		return false
	}
	// We want Pop to give us the latest, not earliest, gethering so we use greater than here.
	return pq[i][0].CreatedAt.After(pq[j][0].CreatedAt)
}

func (pq GetheringPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *GetheringPriorityQueue) Push(x interface{}) {
	item := x.([]Gethering)
	*pq = append(*pq, item)
}

func (pq *GetheringPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func MergeGetherings(pq GetheringPriorityQueue, offset, limit int) []Gethering {
	merged := make([]Gethering, 0, limit)
	heap.Init(&pq)
	numVisitedGetherings := 0

	for len(pq) > 0 && numVisitedGetherings < offset+limit {
		list := pq[0]

		if len(list) == 0 {
			heap.Pop(&pq)
		} else {
			if numVisitedGetherings >= offset {
				gethering := list[0]
				merged = append(merged, gethering)
			}
			pq[0] = list[1:]
			heap.Fix(&pq, 0)
			numVisitedGetherings++
		}
	}
	return merged
}
