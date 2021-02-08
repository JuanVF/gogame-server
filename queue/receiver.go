package queue

import (
	"fmt"
	"runtime"
)

type Receiver struct {
	orderQueue   *Queue
	queues       []*Queue
	currentQueue int
}

// NewReceiver will create a new receiver
func NewReceiver() *Receiver {
	receiver := &Receiver{
		orderQueue:   NewQueue(),
		queues:       make([]*Queue, 0),
		currentQueue: 0,
	}

	receiver.init()

	return receiver
}

// init will create all the queues based on your CPU cores
func (r *Receiver) init() {
	go r.orderQueue.Run()

	for i := 0; i < runtime.NumCPU(); i++ {
		r.queues = append(r.queues, NewQueue())
		go r.queues[i].Run()
	}
}

// Stop will stop all the queues listener
func (r *Receiver) Stop() {
	r.orderQueue.Stop()

	for _, queue := range r.queues {
		queue.Stop()
	}
}

// Execute will append a function to a certain queue and will execute it
// This method can run several functions at the same time
func (r *Receiver) Execute(function Void) error {
	if r.orderQueue == nil {
		return fmt.Errorf("Order queue has not been initialized...\n")
	}

	return r.getAvailableQueue().Enqueue(function)
}

// ExecuteOrder will append a function to the order queue and will execute it
// this method execute all task as it comes
func (r *Receiver) ExecuteOrder(function Void) error {
	if r.orderQueue == nil {
		return fmt.Errorf("Order queue has not been initialized...\n")
	}

	return r.orderQueue.Enqueue(function)
}

// getAvailableQueue will return the queue that is available to execute tasks
func (r *Receiver) getAvailableQueue() *Queue {
	if len(r.queues) == 0 {
		return nil
	}

	queue := r.queues[r.currentQueue]
	r.currentQueue++

	if r.currentQueue >= len(r.queues) {
		r.currentQueue = 0
	}

	return queue
}
