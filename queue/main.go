package main

import "fmt"

// linear data structure
// =<-dequeue [=======================] back/tail/rear =<- enqueue

type Queue struct {
	data []interface{}
}

func (q *Queue) Len() int {
	return len(q.data)

}

func (q *Queue) Enqueue(val interface{}) {
	q.data = append(q.data, val)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.data) == 0 {
		return nil // Queue is empty
	}
	item := q.data[0]
	q.data = q.data[1:]
	return item

}

func (q *Queue) Front() interface{} {
	if len(q.data) == 0 {
		return nil // Queue is empty
	}
	return q.data[0]
}

// IsEmpty returns true if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func main() {
	var q = Queue{}
	q.Enqueue(12)
	q.Enqueue(13)
	q.Enqueue(14)
	q.Enqueue(15)
	q.Enqueue(16)
	q.Enqueue(17)
	fmt.Println(">>>>", q)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	fmt.Println(">>>>", q)

}
