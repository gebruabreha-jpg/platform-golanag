// main.go
package main

import (
	"fmt"
	"generic_type/queue"
	"generic_type/stack"
)

func main() {
	s := stack.New[int]()
	s.Push(10)
	fmt.Println(s.Pop())

	q := queue.New[string]()
	q.Enqueue("apple")
	fmt.Println(q.Dequeue())
}
