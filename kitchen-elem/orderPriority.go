package kitchen_elem

import (
	"container/heap"
)

//// An Item is something we manage in a priority queue.
//type Item struct {
//	value    string // The value of the item; arbitrary.
//	priority int    // The priority of the item in the queue.
//	// The index is needed by update and is maintained by the heap.Interface methods.
//	index int // The index of the item in the heap.
//}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*OrderInKitchen

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	//Pop will return the item will lowest priority
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*OrderInKitchen)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) Update(item *OrderInKitchen) {

	heap.Fix(pq, item.Index)
}
