package bag

import "testing"

func TestSize(t *testing.T) {
	var b Bag

	if b.Size() != 0 {
		t.Fatalf("Size of bag should be 0")
	}
}

func TestIsEmpty(t *testing.T) {
	var b Bag

	if !b.IsEmpty() {
		t.Fatalf("Bag should be empty")
	}
}

func TestAdd(t *testing.T) {
	var b Bag
	b.Add(3)
	b.Add(4, 5, 6)

	if b.Size() != 4 {
		t.Fatalf("Bag should contain 4 items")
	}
}
