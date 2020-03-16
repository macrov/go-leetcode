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

func (l *ListNode) ToString() string {
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
	var l1, l2, l3, head, curr *ListNode
	curr = &ListNode{}
	head = curr
	for _, n := range []int{1,4,5} {
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

	curr = &ListNode{}
	head = curr
	for _, n := range []int{2,6} {
		curr.Next = &ListNode{n, nil}
		curr = curr.Next
	}
	head = head.Next
	l3 = head
	fmt.Println(mergeKLists([]*ListNode{l1,l2,l3}).ToString())
}
