package BreadthFirstSearch

// Queue represents a queue data structure.
type Queue struct {
	items [][][]int
}

// NewQueue creates a new Queue.
func NewQueue() *Queue {
	return &Queue{items: [][][]int{}}
}

// Enqueue adds an item to the queue.
func (q *Queue) Enqueue(item [][]int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front item from the queue.
func (q *Queue) Dequeue() [][]int {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
