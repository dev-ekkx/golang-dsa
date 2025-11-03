package main

import (
	"errors"
	"fmt"
)

type stack struct {
	items []int
}

// push adds an item to the top of the stack
func (s *stack) push(item int) {
	s.items = append(s.items, item)

}

// pop removes and returns the top item of the stack
func (s *stack) pop() (int, error) {
	if len(s.items) == 0 {
		return -1, errors.New("stack is empty")
	}

	l := len(s.items) - 1

	item := s.items[l]
	s.items = s.items[:l]
	return item, nil
}

func main() {
	myStack := stack{}
	myStack.push(10)
	myStack.push(20)
	myStack.push(30)

	fmt.Println("Stack items:", myStack.items)

	_, err := myStack.pop()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Stack items after pop:", myStack.items)
}
