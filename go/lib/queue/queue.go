package queue

import "sort"

// Basic minQ where the node w/ the lowest priority is the first elem in the
// queue
type MinQ struct {
	keys  []string
	nodes map[string]int
}

// Len impls the sort.Interface
func (q *MinQ) Len() int {
	return len(q.keys)
}

// Swap impls the sort.Interface
func (q *MinQ) Swap(i, j int) {
	q.keys[i], q.keys[j] = q.keys[j], q.keys[i]
}

// Less impls the sort.Interface
func (q *MinQ) Less(i, j int) bool {
	x := q.keys[i]
	y := q.keys[j]
	return q.nodes[x] < q.nodes[y]
}

// Push inserts a new key in the queue
func (q *MinQ) Push(key string, pri int) {
	// insert new key if it doesn't exist
	if _, ok := q.nodes[key]; !ok {
		q.keys = append(q.keys, key)
	}

	// set the priority
	q.nodes[key] = pri
	// sort the keys
	sort.Sort(q)
}

// Pop removes the first elem in the queue and returns the key AND the priority
func (q *MinQ) Pop() (key string, pri int) {
	// grab the key
	key, keys := q.keys[0], q.keys[1:]
	q.keys = keys

	pri = q.nodes[key]
	delete(q.nodes, key)

	return key, pri
}

// Empty return true if the queue is empty
func (q *MinQ) Empty() bool {
	return len(q.keys) == 0
}

// Get returns the priority of the key that's passed in
func (q *MinQ) Get(key string) (pri int, ok bool) {
	pri, ok = q.nodes[key]
	return
}

// New creates an empty queue
func New() *MinQ {
	var q MinQ
	q.nodes = make(map[string]int)
	return &q
}
