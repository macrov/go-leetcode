package main

import (
	"fmt"
	. "leetcode/utils/list"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fakeHead := &ListNode{0, head}
	start, end := fakeHead, fakeHead
	for i := 0; i <= n; i++ {
		end = end.Next
	}

	for end != nil {
		end = end.Next
		start = start.Next
	}

	start.Next = start.Next.Next
	return fakeHead.Next
}

func main() {
	fmt.Println(removeNthFromEnd(FromInt(1,2), 1))
}
