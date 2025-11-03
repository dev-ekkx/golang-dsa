package main

import (
	"errors"
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

// prepend adds a new node to the beginning of the list
func (l *linkedList) prepend(data int) {
	newNode := &node{data: data}
	second := l.head
	l.head = newNode
	l.head.next = second
	l.length++
}

// printListData prints all elements in the list
func (l linkedList) printListData() {
	if l.length == 0 {
		fmt.Println("Empty list")
		return
	}

	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

// deleteWithValue removes the first occurrence of a value from the list
func (l *linkedList) deleteWithValue(value int) error {
	if l.length == 0 {
		return errors.New("list is empty")
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return nil
	}

	current := l.head
	for current.next != nil && current.next.data != value {
		current = current.next
	}

	if current.next == nil {
		return fmt.Errorf("value %d not found in list", value)
	}

	current.next = current.next.next
	l.length--
	return nil
}

func main() {
	myLinkedList := linkedList{}

	// Add elements
	values := []int{8, 31, 2, 16, 18, 48}
	for _, v := range values {
		myLinkedList.prepend(v)
	}

	fmt.Printf("Initial list: ")
	myLinkedList.printListData()

	// Delete a value
	if err := myLinkedList.deleteWithValue(8); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("After deleting 8: ")
	myLinkedList.printListData()
}
