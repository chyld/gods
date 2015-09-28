package linkedlist

import "testing"

func TestAdd(t *testing.T) {
	var l LinkedList
	l.Add("Alpha")
	l.Add("Beta")
	l.Add("Gamma")

	if l.Size() != 3 {
		t.Fatalf("List should have 3 items")
	}
}

func TestString(t *testing.T) {
	var l LinkedList
	l.Add("Alpha")
	l.Add("Beta")
	l.Add("Gamma")
	t.Logf("String: %#v", l.String())
}

func TestIsEmpty(t *testing.T) {
	var l LinkedList

	if !l.IsEmpty() {
		t.Fatalf("List should be empty")
	}
}

func TestSize(t *testing.T) {
	var l LinkedList

	if l.Size() != 0 {
		t.Fatalf("List should have size 0")
	}
}
