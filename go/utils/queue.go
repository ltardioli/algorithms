package utils

import "fmt"

type QueueType interface {
	rune | string | int
}

type Queue[T QueueType] struct {
	elements []T
}

func (q *Queue[T]) PreAlloc(size int) {
	q.elements = make([]T, 0, size)
}

func (q *Queue[T]) Enqueue(element T) {
	q.elements = append(q.elements, element)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, fmt.Errorf("empty queue")
	}

	e := q.elements[0]
	q.elements = q.elements[1:]

	return e, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}
