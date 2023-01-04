package main

import (
	"fmt"
	"sync"
)

type Item interface{}

type Queue struct {
	items []Item
	mutex sync.Mutex
}

func (queue *Queue) Peek() Item {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	if len(queue.items) == 0 {
		return nil
	}

	return queue.items[len(queue.items)-1]
}
func (queue *Queue) Dump() []Item {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	var copiedqueue = make([]Item, len(queue.items))
	copy(copiedqueue, queue.items)

	return copiedqueue
}
func (queue *Queue) IsEmpty() bool {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	return len(queue.items) == 0
}
func (queue *Queue) Reset() {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	queue.items = nil
}
func (queue *Queue) Enqueue(item Item) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	queue.items = append(queue.items, item)
}

func (queue *Queue) Dequeue() Item {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	if len(queue.items) == 0 {
		return nil
	}

	lastItem := queue.items[0]
	queue.items = queue.items[1:]

	return lastItem
}
func main() {
	var queue Queue

	queue.Enqueue(5)
	queue.Enqueue(4)
	queue.Enqueue(3)
	queue.Enqueue(2)
	queue.Enqueue(1)

	fmt.Println("Queue:", queue.Dump())
	fmt.Println("The last item:", queue.Peek())

	queue.Dequeue()

	fmt.Println("Queue:", queue.Dump())
	fmt.Println("Queue is empty:", queue.IsEmpty())

	queue.Reset()

	fmt.Println("Queue is empty:", queue.IsEmpty())
}
