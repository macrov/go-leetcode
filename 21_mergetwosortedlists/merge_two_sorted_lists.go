package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type ListNode struct {
	Val int
	Next *ListNode
}

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

func printList(l *ListNode) string {
	sb := bytes.NewBufferString("{")
	for l.Next != nil {
		sb.WriteString(strconv.Itoa(l.Val))
		sb.WriteString("->")
		l = l.Next
	}
	sb.WriteString(strconv.Itoa(l.Val))
	sb.WriteString("}")
	return sb.String()
}

func main() {
	var l1, l2, head, curr *ListNode
	curr = &ListNode{}
	head = curr
	for _, n := range []int{1,2,4} {
		curr.Next = &ListNode{n, nil}
		curr = curr.Next
	}
	head = head.Next
	l1 = head

	curr = &ListNode{}
	head = curr
	for _, n := range []int{1,3,4} {
		curr.Next = &ListNode{n, nil}
		curr = curr.Next
	}
	head = head.Next
	l2 = head
	fmt.Println(printList(mergeTwoLists(l1, l2)))
}
