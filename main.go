package main

import (
  top "main/topmovies"
)

func main() {
  // aa := top.CreateLinkedList([]int{34,8,91,6,10})
  // top.Display(aa)

	a := top.CreateLinkedList([]int{11, 41, 51})
	b := top.CreateLinkedList([]int{21,23,42})
	c := top.CreateLinkedList([]int{25,56,66,72})
	list1 := []*top.LinkedListNode{a, b, c}
	top.Display(top.MergeKSortedLists(list1))  
}
