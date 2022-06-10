package topmovies


// list1 := []*LinkedListNode{}
func merge2SortedLists(l1, l2 *LinkedListNode) *LinkedListNode {
  //flow ~ 
  dummy := &LinkedListNode{data: -1}
  prev := dummy

  for l1 != nil && l2 != nil {
    if l1.data <= l2.data {
      prev.next = l1
      l1 = l1.next
    } else {
      prev.next = l2
      l2 = l2.next
    }
    prev = prev.next
  }

  if l1 == nil {
    prev.next = l1
  } else {
    prev.next = l2
  }

  return prev.next
}

func MergeKSortedLists(lists []*LinkedListNode) *LinkedListNode {
  if len(lists) > 0 {
    res := lists[0]
    
    for i := 1; i < len(lists); i++ {
      res = merge2SortedLists(res, lists[i])
    }
    
    return res
  }

  return &LinkedListNode{data: -1}
}