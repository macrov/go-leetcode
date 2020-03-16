package main

import (
	"fmt"
	. "leetcode/utils/list"
)

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	curr := dummyHead
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				curr.Next = l1
				curr = curr.Next
				l1 = l1.Next
			} else {
				curr.Next = l2
				curr = curr.Next
				l2 = l2.Next
			}
		} else if l1 == nil && l2 != nil {
			curr.Next = l2
			curr = curr.Next
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			curr.Next = l1
			curr = curr.Next
			l1 = l1.Next
		}
	}
	return dummyHead.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	var res *ListNode
	for _, l := range lists {
		res = mergeTwoLists(res, l)
	}
	return res
}

func main() {
	fmt.Println(mergeKLists([]*ListNode{FromInt(1,4,5), FromInt(1,3,4), FromInt(2,6)}).ToString())
}
