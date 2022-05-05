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

func main() {
  linkList := new(LinkedList[string])

  linkList.insertAtEnd(&Node[string]{data: "This"})
  linkList.insertAtEnd(&Node[string]{data: "is"})
  linkList.insertAtEnd(&Node[string]{data: "a"})
  linkList.insertAtEnd(&Node[string]{data: "linked"})
  linkList.insertAtEnd(&Node[string]{data: "list"})

  linkList.PrintAll()
}
