package main

import (
	"errors"
	"fmt"
)

type queue struct {
	items []int
}

// enqueue adds an item to the end of the queue
func (q *queue) enqueue(item int) {
	q.items = append(q.items, item)
}

// dequeue removes and returns the first item in the queue
func (q *queue) dequeue() (int, error) {
	if len(q.items) == 0 {
		return -1, errors.New("queue is empty")
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}
func main() {
	myQueue := queue{}
	myQueue.enqueue(10)
	myQueue.enqueue(20)
	myQueue.enqueue(30)

	fmt.Println("Queue items:", myQueue.items)

	_, err := myQueue.dequeue()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Queue items after dequeue:", myQueue.items)
}
