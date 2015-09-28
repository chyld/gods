package queue

import "testing"

func TestSize(t *testing.T) {
	var q Queue

	if q.Size() != 0 {
		t.Fatalf("Size of queue should be 0")
	}
}

func TestIsEmpty(t *testing.T) {
	var q Queue

	if !q.IsEmpty() {
		t.Fatalf("Queue should be empty")
	}
}

func TestEnqueue(t *testing.T) {
	var q Queue
	q.Enqueue(3)

	if q.IsEmpty() {
		t.Fatalf("Enqueued item, queue should NOT be empty")
	}
}

func TestDequeue(t *testing.T) {
	var q Queue
	blank := q.Dequeue()
	if blank != nil {
		t.Fatalf("Dequeued nothing")
	}

	q.Enqueue(3)
	one := q.Dequeue()
	one = one.(int)
	if one != 3 {
		t.Fatalf("Dequeued item should be 1")
	}

	q.Enqueue(5)
	q.Enqueue(7)
	first := q.Dequeue()
	first = first.(int)

	if first != 5 {
		t.Fatalf("First dequeued item should be a 5")
	}
	if q.Size() != 1 {
		t.Fatalf("After dequeue, should have 1 items remaining")
	}
}
