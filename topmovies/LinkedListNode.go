package topmovies

import (
	"fmt"
	"rand"
)

type LinkedListNode struct {
  key               int
  data              int
  next              *LinkedListNode
  arbitrartPointer  *LinkedListNode
}

func createLinkedList(lst []int) *LinkedListNode {
  var head *LinkedListNode
  var tail *LinkedListNode

  for _, x := range lst {
    newNode := &LinkedListNode{ data: x }
    if head == nil {
      head = newNode
    } else {
      tail.next = newNode
    }
    tail = newNode
  }

  return head
}

func insertAtHead(head *LinkedListNode, data int) *LinkedListNode {
  newNode := &LinkedListNode{data: data}
  newNode.next = head
  return newNode
}

func insertAtTail(head *LinkedListNode, data int) *LinkedListNode {
  newNode := &LinkedListNode{data: data}
  if head == nil {
    return newNode
  }

  temp := head
  for temp.next != nil {
    temp = temp.next
  }
  temp.next = newNode
  return head
}

func createRandomList(length int) *LinkedListNode {
  var listHead *LinkedListNode
  for i := 0; i < length; i++ {
    listHead = insertAtHead(listHead, rand.Intn(100))
  }
  return listHead
}

func toList(head *LinkedListNode) []int {
  var lst []int
  temp := head
  for temp != nil {
    lst = append(lst, temp.data)
  }

  return lst
}

func display(head *LinkedListNode) {
  temp := head
  for temp != nil {
    fmt.Println("%d", temp.data)
    temp = temp.next
    if temp != nil {
      fmt.Println(", ")
    }
  }
  fmt.Println("\n")
}

func mergeAlternating(list1, list2 *LinkedListNode) *LinkedListNode {
  
}

func isEqual(list1, list2 *LinkedListNode) bool {}