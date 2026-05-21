// queue/queue.go
package queue

type Queue[T any] struct {
	items []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue adds an element at the end
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes the first element
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Front returns the first element without removing it
func (q *Queue[T]) Front() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	return q.items[0], true
}
