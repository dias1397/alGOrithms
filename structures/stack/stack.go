package stack

import "errors"

type Stack[T any] struct {
	length int
	head   *item[T]
}

type item[T any] struct {
	value T
	prev  *item[T]
}

func (stack *Stack[T]) Push(value T) {
	item := item[T]{value: value, prev: stack.head}

	stack.head = &item

	stack.length++
}

func (stack *Stack[T]) Pop() (T, error) {
	if stack.length == 0 {
		return *new(T), errors.New("Stack is empty, unable to pop")
	}

	item := stack.head

	stack.head = item.prev
	stack.length--

	item.prev = nil

	return item.value, nil
}

func (stack *Stack[T]) Peek() T {
	return stack.head.value
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.head == nil && stack.length == 0
}

func (stack *Stack[T]) Size() int {
	return stack.length
}
