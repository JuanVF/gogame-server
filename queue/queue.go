package queue

import (
	"fmt"
)

// Void is a type to send void functions as arguments
type Void func()

// Queue is a basic implementation of a queue of void functions
type Queue struct {
	queue []Void
	data  chan bool
}

// NewQueue is a constructor for Queue data structure
func NewQueue() *Queue {
	queue := &Queue{
		queue: make([]Void, 0),
		data:  make(chan bool),
	}

	return queue
}

// Run will be a listener that will execute all the tasks
func (q *Queue) Run() {
	defer close(q.data)

	for {
		if status := <-q.data; !status {
			break
		}

		q.runTasks()
	}
}

// Stop will stop the queue, of course will wait for all running tasks to be executed
func (q *Queue) Stop() {
	q.data <- false
}

// runTasks will execute all the tasks has been enqueued
func (q *Queue) runTasks() {
	queue, _ := q.Dequeue()

	for {
		if queue == nil {
			break
		}

		queue()

		queue, _ = q.Dequeue()
	}
}

// Dequeue this will dequeue if a element exists, otherwise will return nil
// if queue has not been initialized it will return an error
func (q *Queue) Dequeue() (Void, error) {
	if !q.IsInitialized() {
		return nil, fmt.Errorf("Queue has not been initialized...\n")
	}

	if q.IsEmpty() {
		return nil, nil
	}

	el := q.queue[0]
	q.queue = q.queue[1:]

	return el, nil
}

// Enqueue this will enqueue a void function
// this will return error if queue has not been initialized
func (q *Queue) Enqueue(v Void) error {
	if !q.IsInitialized() {
		return fmt.Errorf("Queue has not been initialized...\n")
	}

	q.queue = append(q.queue, v)
	q.data <- true

	return nil
}

// Peek will return the first element if a exists without removing it
func (q *Queue) Peek() (Void, error) {
	if !q.IsInitialized() {
		return nil, fmt.Errorf("Queue has not been initialized...\n")
	}

	if q.IsEmpty() {
		return nil, nil
	}

	return q.queue[0], nil
}

// IsEmpty will return true if the queue is empty
// Will also return true if it has not been initialized
func (q *Queue) IsEmpty() bool {
	if !q.IsInitialized() {
		return true
	}

	return len(q.queue) == 0
}

// IsInitialized This will return true if the queue has been initialized
func (q *Queue) IsInitialized() bool {
	return q.queue != nil
}

// Length will return the amount of functions the queue has
func (q *Queue) Length() int {
	if !q.IsInitialized() {
		return 0
	}

	return len(q.queue)
}
