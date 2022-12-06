package queue

import "container/heap"

// This queue implements the heap.Interface and holds Items.
type MinQ []*Item

func (pq MinQ) Len() int {
	return len(pq)
}

func (pq MinQ) Less(i, j int) bool {
	// for this MinPriorityQueue, we want the Pop() to return the **lowest**
	// priority, hence why we use greater than in this impl.
	return pq[i].Priority < pq[j].Priority
}

func (pq MinQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *MinQ) Push(i interface{}) {
	n := len(*pq)
	item := i.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *MinQ) Pop() interface{} {
	old := *pq
	n := len(*pq)
	item := old[n-1]
	old[n-1] = nil  // avoids memory leaks
	item.Index = -1 // set it at the bounds for memory safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority & value of an Item in the queue
func (pq *MinQ) update(i *Item, value interface{}, priority int) {
	i.Value = value
	i.Priority = priority
	heap.Fix(pq, i.Index)
}
