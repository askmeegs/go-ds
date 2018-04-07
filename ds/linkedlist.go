package ds

import (
	"fmt"
	"reflect"
)

type Node struct {
	key  interface{}
	prev *Node
	next *Node
}

// DLL advantage = you can iterate in both directions, easy to reverse
// in a SLL you can't go backwards

// DLL disadvantages = more space intensive (have to store prev)
// Insert and Delete takes more time (more pointer operations)
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

// deletes the first instance of find
func (L *DoublyLinkedList) Delete(find interface{}) {
	cur := L.head
	for cur != nil {
		if reflect.DeepEqual(cur.key, find) {
			// case 1 = find is head
			if cur.prev == nil {
				if cur.next != nil {
					cur.next.prev = nil
					L.head = cur.next
				} else {
					L = &DoublyLinkedList{}
				}
			} else {
				// case 2 = find is tail
				// b->d
				if cur.next == nil {
					cur.prev.next = nil
					L.tail = cur.prev

				} else {
					// case 3 = find is in middle -- skip over cur
					cur.prev.next = cur.next
					cur.next.prev = cur.prev
				}
			}
			return
		}
		cur = cur.next
	}
	fmt.Printf("%v not found in list: %s\n", find, L.ToString())
	return
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
	// fmt.Printf("Node key=%v   prev=%#v,     next=%#v ----> \n", cur.key, cur.prev, cur.next)

	for cur.next != nil {
		// fmt.Printf("Node key=%v   prev=%#v,     next=%#v ----> \n", cur.next.key, cur.next.prev, cur.next.next)

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
