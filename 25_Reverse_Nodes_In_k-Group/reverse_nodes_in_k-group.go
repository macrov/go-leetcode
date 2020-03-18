package main

import (
	"fmt"
	. "leetcode/utils/list"
)

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Val: -1,
		Next: head,
	}
	pre := dummy
	cur := dummy.Next

	for {
		n := k
		nextPart := cur
		for nextPart != nil && n > 0 {
			nextPart = nextPart.Next
			n--
		}

		if n > 0 {
			break
		}

		nextPre := cur
		for n < k {
			temp := cur.Next
			cur.Next = nextPart
			nextPart = cur
			cur = temp
			n++
		}
		pre.Next = nextPart
		pre = nextPre
		cur = pre.Next
	}
	return dummy.Next
}

func main() {
	fmt.Println(reverseKGroup(FromInt(1,2,3,4,5,6,7), 3).String())
}
