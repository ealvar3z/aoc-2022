package main

type task struct {
	left, right int
}

// Contains return true iff the given pair can be exactly included inside the
// current task's range.
func (t task) Contains(pair task) bool {
	return pair.left >= t.left && pair.right <= t.right
}

// Overlaps return true iff the given pair of tasks can exactly included inside
// the current task range on either end.
func (t task) Overlaps(pair task) bool {
	if pair.left > t.right {
		return false
	}
	if pair.right < t.left {
		return false
	}

	return pair.left >= t.left || pair.right <= t.right
}
