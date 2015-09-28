package linkedlist

import "fmt"

type node struct {
	body       interface{}
	prev, next *node
}

// LinkedList todo
type LinkedList struct {
	head, tail *node
	size       int
}

/* -------------------------------------------------------------------------- */
/* -------------------------------------------------------------------------- */
/* -------------------------------------------------------------------------- */

// Add todo
func (l *LinkedList) Add(item interface{}) {
	n := &node{body: item}

	if l.IsEmpty() {
		l.head, l.tail = n, n
	} else {
		l.tail.next = n
		n.prev = l.tail
		l.tail = n
	}

	l.size++
}

// IsEmpty tooo
func (l *LinkedList) IsEmpty() bool {
	return l.Size() == 0
}

// Size todo
func (l *LinkedList) Size() int {
	return l.size
}

func (l LinkedList) String() string {
	slice := toSlice(&l)
	return fmt.Sprintf("%#v", slice)
}

/* -------------------------------------------------------------------------- */
/* -------------------------------------------------------------------------- */
/* -------------------------------------------------------------------------- */

func toSlice(l *LinkedList) []interface{} {
	var slice []interface{}

	if l.head == nil {
		return slice
	}

	descend(l.head, &slice)
	return slice

}

func descend(n *node, s *[]interface{}) {
	*s = append(*s, n.body)

	if n.next == nil {
		return
	}

	descend(n.next, s)
}
