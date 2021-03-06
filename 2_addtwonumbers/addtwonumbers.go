package addtwonumbers

import (
	. "leetcode/utils/list"
)

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	var resListNode, currentNode *ListNode
//	var carry int = 0
//	for l1!=nil || l2!=nil || carry!=0 {
//		var n1, n2, n3 int = 0, 0, 0
//		if l1 != nil {
//			n1 = l1.Val
//			l1 = l1.Next
//		}
//		if l2 != nil {
//			n2 = l2.Val
//			l2 = l2.Next
//		}
//		n3 = (n1 + n2 + carry) % 10
//		carry = (n1 + n2 + carry) / 10
//		if resListNode == nil {
//			resListNode = &ListNode{n3, nil}
//			currentNode = resListNode
//		} else {
//			currentNode.Next = &ListNode{n3, nil}
//			currentNode = currentNode.Next
//		}
//	}
//	return resListNode
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoList(l1,l2,0)
}

func addTwoList(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		if carry > 0 {
			return &ListNode{carry, nil}
		} else {
			return  nil
		}
	}

	if l1 == nil {
		l1 = &ListNode{0, nil}
	}
	if l2 == nil {
		l2 = &ListNode{0, nil}
	}
	sum := l1.Val + l2.Val + carry
	current, carry := sum % 10, sum / 10
	return &ListNode{current, addTwoList(l1.Next, l2.Next, carry)}
}
