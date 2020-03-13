package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fakeHead := &ListNode{0, head}
	start, end := fakeHead, fakeHead
	for i:=0;i<=n;i++ {
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
	var head, curr *ListNode
	curr = &ListNode{}
	head = curr
	for _, n := range []int{1, 2} {
		curr.Next = &ListNode{n, nil}
		curr = curr.Next
	}
	head = head.Next
	fmt.Println(removeNthFromEnd(head, 1))
}
