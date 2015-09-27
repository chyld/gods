package bag

// Bag todo
type Bag struct {
	data []interface{}
}

// Add todo
func (b *Bag) Add(item ...interface{}) {
	b.data = append(b.data, item...)
}

// Size todo
func (b *Bag) Size() int {
	return len(b.data)
}

// IsEmpty todo
func (b *Bag) IsEmpty() bool {
	return b.Size() == 0
}
