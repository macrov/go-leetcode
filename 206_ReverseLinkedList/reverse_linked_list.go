package main

import (
	"fmt"
	. "leetcode/utils/list"
)

/*
 1->2->3->4->5->nil to 5->4->3->2->1->nil

 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {return nil}
	dummyHead := &ListNode{0, head}
	var pre, curr = dummyHead, dummyHead.Next
	for curr != nil {
		tmp := curr.Next
		curr.Next = pre
		pre = curr
		curr = tmp
	}
	head.Next = nil
	return pre
}

func main() {
	fmt.Println(reverseList(FromInt(1,2,3,4,5)))
}