package queue

type Item struct {
	Value    any
	Priority int
	Index    int // index of the item in the heap
}
