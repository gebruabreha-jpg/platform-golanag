// stack/stack.go
package stack

type Stack[T any] struct {
	items []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push an element onto the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop an element from the stack
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, true
}

// Peek at the top element without removing it
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}
