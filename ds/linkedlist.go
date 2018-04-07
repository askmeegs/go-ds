package ds

import (
	"fmt"
)

type Node struct {
	key  interface{}
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// adds Node to head of list
func (L *DoublyLinkedList) Insert(key interface{}) {
	temp := &Node{
		next: L.head,
		key:  key,
	}
	// as long as list is nonempty, update the previous head to point to new head
	if L.head != nil {
		L.head.prev = temp
	}
	L.head = temp //update with new head

	// update tail
	z := L.head
	for z.next != nil {
		z = z.next
	}
	L.tail = z
}

func (L *DoublyLinkedList) ToString() string {
	if L == nil {
		return ""
	}
	if L.head == nil {
		return ""
	}
	cur := L.head
	output := fmt.Sprintf("%v", cur.key)

	for cur.next != nil {
		output = fmt.Sprintf("%s -> %v", output, cur.next.key)
		cur = cur.next
	}
	return output
}

func (L *DoublyLinkedList) Reverse() *DoublyLinkedList {
	R := &DoublyLinkedList{}
	if L == nil || L.tail == nil {
		return R
	}

	cur := L.head
	R.Insert(cur.key)
	for cur.next != nil {
		R.Insert(cur.next.key)
		cur = cur.next
	}
	return R
}
