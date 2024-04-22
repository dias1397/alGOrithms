package queue

import "errors"

type Queue[T any] struct {
	length int
	head   *node[T]
	tail   *node[T]
}

type node[T any] struct {
	value T
	next  *node[T]
}

func (queue *Queue[T]) Enqueue(value T) {
	newNode := node[T]{value: value}

	if queue.tail != nil {
		queue.tail.next = &newNode
	}

	queue.tail = &newNode
	queue.length++

	if queue.head == nil {
		queue.head = &newNode
	}
}

func (queue *Queue[T]) Dequeue() (T, error) {
	if queue.length == 0 {
		return *new(T), errors.New("Queue is empty, unable to dequeu")
	}

	node := queue.head
	queue.head = queue.head.next
	queue.length--

	if queue.head == nil {
		queue.tail = nil
	}

	return node.value, nil
}

func (queue *Queue[T]) Peek() T {
	if queue.head == nil {
		return *new(T)
	}

	return queue.head.value
}

func (queue *Queue[T]) IsEmpty() bool {
	return queue.head == nil && queue.tail == nil && queue.length == 0
}

func (queue *Queue[T]) Size() int {
	return queue.length
}
