package stack

import "errors"

type Stack[T any] struct {
	length int
	head   *stackItem[T]
}

type stackItem[T any] struct {
	value T
	prev  *stackItem[T]
}

func (stack *Stack[T]) Push(value T) {
	stack.head = &stackItem[T]{value: value, prev: stack.head}
	stack.length++
}

func (stack *Stack[T]) Pop() (T, error) {
	if stack.length == 0 {
		var defaultValue T
		return defaultValue, errors.New("Stack is empty, unable to pop")
	}

	item := stack.head
	stack.head = item.prev
	stack.length--
	item.prev = nil

	return item.value, nil
}

func (stack *Stack[T]) Peek() T {
	if stack.head == nil {
		var defaultValue T
		return defaultValue
	}
	return stack.head.value
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.head == nil && stack.length == 0
}

func (stack *Stack[T]) Size() int {
	return stack.length
}
