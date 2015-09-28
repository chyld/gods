package queue

// Queue todo
type Queue struct {
	data []interface{}
}

// Enqueue todo
func (q *Queue) Enqueue(item interface{}) {
	q.data = append(q.data, item)
}

// Dequeue todo
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	item := q.data[0]
	q.data = q.data[1:]
	return item
}

// Size todo
func (q *Queue) Size() int {
	return len(q.data)
}

// IsEmpty todo
func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}
