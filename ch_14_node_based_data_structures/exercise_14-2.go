package main

import "fmt"

type DoubleNode[T any] struct {
  data T
  next *DoubleNode[T]
  prev *DoubleNode[T]
}

type DoublyLinkedList[T any] struct {
  head *DoubleNode[T]
  tail *DoubleNode[T]
}

func (dl *DoublyLinkedList[T]) insertAtEnd(n *DoubleNode[T]) {
  cursor := dl.tail

  if cursor == nil {
    dl.head = n
    dl.tail = n
    return
  }

  cursor.next = n
  n.prev = cursor
  dl.tail = n
}

func (dl *DoublyLinkedList[T]) PrintAllReverse() {
  cursor := dl.tail

  for cursor != nil {
    fmt.Println(cursor.data)
    cursor = cursor.prev
  }
}

func (dl *DoublyLinkedList[T]) PrintAll() {
  cursor := dl.head

  for cursor != nil {
    fmt.Println(cursor.data)
    cursor = cursor.next
  }
}

func main() {
	doubleLinkList := new(DoublyLinkedList[string])

	doubleLinkList.insertAtEnd(&DoubleNode[string]{data: "list"})
	doubleLinkList.insertAtEnd(&DoubleNode[string]{data: "linked"})
	doubleLinkList.insertAtEnd(&DoubleNode[string]{data: "doubly"})
	doubleLinkList.insertAtEnd(&DoubleNode[string]{data: "a"})
	doubleLinkList.insertAtEnd(&DoubleNode[string]{data: "is"})
	doubleLinkList.insertAtEnd(&DoubleNode[string]{data: "This"})

	doubleLinkList.PrintAll()
	doubleLinkList.PrintAllReverse()
}
