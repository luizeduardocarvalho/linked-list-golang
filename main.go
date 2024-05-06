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

func (list *List[T]) removeInPosition(position int64) {
	if list.head != nil {
		current := list.head

		if current.nextValue == nil {
			list.head = nil
			return
		}

		next := current.nextValue

		for i := int64(0); i < position-1; i++ {
			if next.nextValue == nil {
				break
			}

			current = next
			next = current.nextValue
		}

		if position == 0 {
			list.head = current.nextValue
			return
		}

		if next.nextValue == nil {
			current.nextValue = nil
		} else {
			current.nextValue = next.nextValue
			next.nextValue = nil
		}
	}
}

func (list *List[T]) pop() {
	if list.head != nil {
		current := list.head

		if current.nextValue == nil {
			list.head = nil
		}

		for current.nextValue != nil {
			if current.nextValue.nextValue == nil {
				current.nextValue = nil
				break
			}

			current = current.nextValue
		}
	}
}

func (list *List[T]) search(position int64) T {
	current := list.head

	if list.head == nil {
		panic("Index out of bounds")
	}

	if position == 0 {
		return list.head.value
	}

	if current.nextValue == nil {
		panic("Index out of bounds")
	}

	for i := int64(0); i < position; i++ {
		if current.nextValue == nil {
			break
		}

		current = current.nextValue
	}

	return current.value
}

func (list *List[T]) update(value T, position int64) {
	current := list.head

	for i := int64(0); i < position; i++ {
		if current.nextValue == nil {
			panic("Index out of bounds")
		}

		current = current.nextValue
	}

	current.value = value
}

func (list List[T]) isEmpty() bool {
	if list.head == nil {
		return true
	}

	return false
}

func (list List[T]) length() int64 {
	length := int64(0)

	current := list.head

	if current != nil {
		length++
	} else {
		return length
	}

	if current.nextValue == nil {
		return length
	}

	for {
		if current.nextValue == nil {
			break
		}

		current = current.nextValue
		length++
	}

	return length
}

func (list *List[T]) join(listToJoin *List[T]) *List[T] {
	if list.head == nil {
		list.head = listToJoin.head
		return list
	}

	current := list.head

	for {
		if current.nextValue == nil {
			break
		}

		current = current.nextValue
	}

	current.nextValue = listToJoin.head
	return list
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

	// var secondList List[int]

	// secondList.add(3)
	// secondList.add(4)

	// thirdList := list.join(&secondList)

	// thirdList.print()

	// list.update(5, 2)

	// list.removeInPosition(0)

	// value := list.search(3)
	// fmt.Println(value)

	// fmt.Println(list.isEmpty())

	// length := list.length()
	// fmt.Println(length)

	list.print()
}
