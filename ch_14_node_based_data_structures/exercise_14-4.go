/*
	Exercise 14.4

	Add a method to LinkedList that reverses the list, e.g., if the original list is A -> B -> C, the list should change so that its order becomes C -> B -> A 

input: linked list
output: (no output)
	SIDE EFFECT: input linked list's order is reversed so that the former tail is now the head and vice versa (list and nodes are mutated)

notes: can either switch values or next pointers

algo:
	initialize 3 variables (type pointer to Node[T]): previous, cursor, and following
	if the linked list is empty (head points to nil), return
	while cursor.next points to a node (isn't nil), traverse through the linked list beginning at head, updating cursor to point to each node
		reassign following to the pointer currently assigned to the cursor.next field (save next node destination)
		reassign cursor.next to node pointed at by previous (reverse order of list for current node)
		reassign previous to cursor (save a reference to what will be the previous node on next iteration)
		reassign cursor -> following (moving to the next node in the list)

	reassign linked list's head to be cursor (the last node accessed)

	return
*/
package main

import (
	"fmt"
)

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
  head *Node[T]
}

func (ll *LinkedList[T]) insertAtEnd(n *Node[T]) {
  cursor := ll.head

  if cursor == nil {
    ll.head = n
    return
  }

  for cursor.next != nil {
    cursor = cursor.next
  }

  cursor.next = n
}

func (ll *LinkedList[T]) PrintAll() {
  cursor := ll.head

  for cursor != nil {
    fmt.Println(cursor.data)

    cursor = cursor.next
  }
}

func (ll *LinkedList[T]) Reverse() {
	if ll.head == nil { return }
	
	previous, following := (*Node[T])(nil), (*Node[T])(nil)
	cursor := ll.head

	for cursor != nil {
		following = cursor.next
		cursor.next = previous
		previous = cursor
		cursor = following
	}

	ll.head = previous
}

func main() {
	linkList := new(LinkedList[int])
	linkList.PrintAll() // no output: empty linked list
	linkList.Reverse()

	linkList.insertAtEnd(&Node[int]{data: 5})
	linkList.insertAtEnd(&Node[int]{data: 4})
	linkList.insertAtEnd(&Node[int]{data: 3})
	linkList.insertAtEnd(&Node[int]{data: 2})
	linkList.insertAtEnd(&Node[int]{data: 1})

	linkList.PrintAll() // sequence is 5 4 3 2 1
	fmt.Println("reversing...")
	linkList.Reverse()
	linkList.PrintAll() // now list's sequence is 1 2 3 4 5
}
