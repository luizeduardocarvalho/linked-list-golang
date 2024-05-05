package main

import (
	"fmt"
)

type Node[T any] struct {
	nextValue *Node[T]
	value     T
}

type List[T any] struct {
	head *Node[T]
}

func (list *List[T]) add(value T) {
	newNode := &Node[T]{value: value}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head

		for current.nextValue != nil {
			current = current.nextValue
		}

		current.nextValue = newNode
	}
}

func (list *List[T]) addInPosition(value T, position int64) {
	newNode := &Node[T]{value: value}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for i := int64(0); i < position-1; i++ {
			if current.nextValue == nil {
				break
			}

			current = current.nextValue
		}

		newNode.nextValue = current.nextValue
		current.nextValue = newNode
	}
}

func (list List[T]) print() {
	if list.head != nil {
		value := list.head

		fmt.Println(value.value)

		for value.nextValue != nil {
			fmt.Println(value.nextValue.value)
			value = value.nextValue
		}
	}
}

func main() {
	var list List[int]

	list.add(1)
	list.add(2)

	list.addInPosition(3, 1)

	list.print()
}
