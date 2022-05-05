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

func (ll *LinkedList[T]) getLastNode() *Node[T] {
	if ll.head == nil { return nil }

	cursor := ll.head

	for cursor.next != nil {
		cursor = cursor.next
	}

	return cursor
}

func main() {
	linkList := new(LinkedList[string])
	fmt.Println(linkList.getLastNode()) // nil: no nodes in linked list

	linkList.insertAtEnd(&Node[string]{data: "Nodes"})
	linkList.insertAtEnd(&Node[string]{data: "are"})
	linkList.insertAtEnd(&Node[string]{data: "spaced-out"})
	linkList.insertAtEnd(&Node[string]{data: "data"})
	linkList.insertAtEnd(&Node[string]{data: "structures"})

	linkList.PrintAll() // "Nodes" "are" "spaced-out" "data" "structures" 
	fmt.Println(*linkList.getLastNode()) // node containing "structures"
}
