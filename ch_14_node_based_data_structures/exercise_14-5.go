/*
	Exercise 14.5

	Let's say you have access to a node from somewhere in the middle of a linked list but not the linked list itself; that is, you have a variable that points to an instance of Node but not the LinkedList instance itself. 

	In this situation, if you follow this node's link, you can find all the items from this middle node until the end, but you have no way to find the nodes that precede this node in the list.

	Write code that will effectively delete this node from the list. The entire remaining list should remain complete, with only this node removed.

	input: node (to remove from list)
	output: none/SIDE EFFECT: alter list so that node passed-in is removed from list

	note: can traverse through list (beginning at middle node), peek at following data and reassign current node's data to the following node's data, and then continue. When the node two nodes down

	algo:
		get the data in the node following the input node and reassign current node's data to it
		get the address of the node following the next one (two nodes away) and reassign the next pointer of current node to that node (skipping what had previously been the next node)
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

func removeMiddleNodeFromList[T any] (middleNode *Node[T]) {
	middleNode.data = middleNode.next.data
	middleNode.next = middleNode.next.next
}

func main() {
	linkList := new(LinkedList[string])

	linkList.insertAtEnd(&Node[string]{data: "apple"})
	linkList.insertAtEnd(&Node[string]{data: "banana"})
	linkList.insertAtEnd(&Node[string]{data: "cherry"})

	middleLink := &Node[string]{data: "diamond"}
	linkList.insertAtEnd(middleLink)
	linkList.insertAtEnd(&Node[string]{data: "elderberry"})
	linkList.insertAtEnd(&Node[string]{data: "fig"})
	linkList.insertAtEnd(&Node[string]{data: "grape"})

	linkList.PrintAll() // apple banana cherry diamond elderberry fig grape
	fmt.Println("\nremoving diamond node...")
	removeMiddleNodeFromList(middleLink)
	linkList.PrintAll() // apple banana cherry elderberry fig grape
}
