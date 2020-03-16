package list

import (
	"bytes"
	"strconv"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func (l *ListNode) ToString() string {
	if l == nil {
		return "{}"
	}
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

func FromArray(a []int) *ListNode {
	var head, curr *ListNode
	curr = &ListNode{}
	head = curr
	for _, n := range a {
		curr.Next = &ListNode{n, nil}
		curr = curr.Next
	}
	return head.Next
}

func FromInt(ints ...int) *ListNode {
	var head, curr *ListNode
	curr = &ListNode{}
	head = curr
	for _, n := range ints {
		curr.Next = &ListNode{n, nil}
		curr = curr.Next
	}
	return head.Next
}

