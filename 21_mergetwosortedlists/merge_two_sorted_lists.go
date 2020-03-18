package main

import (
	"fmt"
	. "leetcode/utils/list"
)

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dumyHead := &ListNode{}
	curr := dumyHead
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
	return dumyHead.Next
}

func main() {
	fmt.Println(mergeTwoLists(FromInt(1,2,4), FromInt(1,3,4)).String())
}
