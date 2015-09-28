package stack

import "testing"

func TestSize(t *testing.T) {
	var s Stack

	if s.Size() != 0 {
		t.Fatalf("Size of stack should be 0")
	}
}

func TestIsEmpty(t *testing.T) {
	var s Stack

	if !s.IsEmpty() {
		t.Fatalf("Stack should be empty")
	}
}

func TestPush(t *testing.T) {
	var s Stack
	s.Push(3)

	if s.IsEmpty() {
		t.Fatalf("Stack should not be empty")
	}
}

func TestPop(t *testing.T) {
	var s Stack
	blank := s.Pop()
	if blank != nil {
		t.Fatalf("Pop'd item should be nil")
	}

	s.Push(3)
	s.Push(5)
	s.Push(7)
	item := s.Pop()
	item = item.(int)
	if item != 7 {
		t.Fatalf("Pop'd item should be 7")
	}

	if s.Size() != 2 {
		t.Fatalf("Stack size should be 2")
	}
}
