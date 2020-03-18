package main

import (
	"fmt"
	. "leetcode/utils/list"
)

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var dummyHead, currentHead *ListNode
	dummyHead = &ListNode{0, head}
	currentHead = dummyHead
	left := currentHead.Next
	right := left.Next
	for left != nil && right != nil {
		end := right.Next
		currentHead.Next = right
		left.Next = end
		right.Next = left

		currentHead = currentHead.Next.Next
		left = currentHead.Next
		if left != nil {
			right = left.Next
		}
	}
	return dummyHead.Next
}

func main() {
	fmt.Println(swapPairs(FromArray([]int{})).String())
}
