package linkedlist

import (
	"errors"
	"fmt"
)

type Singly[T any] struct {
	length int
	head   *element[T]
}

type element[T any] struct {
	value T
	next  *element[T]
}

func (list *Singly[T]) Append(val T) {
	node := element[T]{value: val}

	if list.head == nil {
		list.head = &node
	} else {
		current := list.head
		for ; current.next != nil; current = current.next {
		}
		current.next = &node
	}

	list.length++
}

func (list *Singly[T]) Prepend(val T) {
	node := element[T]{value: val}

	if list.head == nil {
		list.head = &node
	} else {
		node.next = list.head
		list.head = &node
	}

	list.length++
}

func (list *Singly[T]) InsertAt(val T, index int) error {
	if index > list.length {
		return errors.New("Index out of bounds")
	} else {
		node := element[T]{value: val}

		current := list.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}

		if current == nil {
			list.head = &node
		} else {
			node.next = current.next
			current.next = &node
		}

		list.length++

		return nil
	}
}

func (list *Singly[T]) Remove() (element[T], error) {

	return element[T]{}, nil
}

func (list *Singly[T]) RemoveAt(index int) (element[T], error) {

	return element[T]{}, nil
}

func (list *Singly[T]) Get(index int) (T, error) {
	if list.head == nil {
		return *new(T), errors.New("Linked list is empty")
	} else if index >= list.length {
		return *new(T), errors.New("Index out of bounds")
	} else {
		current := list.head
		for i := 0; i < index; i++ {
			current = current.next
		}

		return current.value, nil
	}
}

func (list *Singly[T]) Size() int {
	return list.length
}

func (list *Singly[T]) Display() {
	for node := list.head; node != nil; node = node.next {
		fmt.Print(node, " ")
	}
	fmt.Println()
}
