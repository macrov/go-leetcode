package addtwonumbers

import (
	"fmt"
	. "leetcode/utils/list"
	"testing"
)

var listNode1, listNode2, p1, p2 *ListNode

func TestAddTwoNumbers(t *testing.T) {
	listNode1Numbers := []int{2,4,3}
	listNode2Numbers := []int{5,6,4}

	p1 = &ListNode{}
	listNode1 = p1
	p2 = &ListNode{}
	listNode2 = p2
	for i, num := range listNode1Numbers {
		p1.Val = num
		if i< len(listNode2Numbers)-1 {
			p1.Next = &ListNode{}
			p1 = p1.Next
		}
	}

	for i, num := range listNode2Numbers {
		p2.Val = num
		if i< len(listNode2Numbers)-1 {
			p2.Next = &ListNode{}
			p2 = p2.Next
		}
	}
	fmt.Println(listNode1)
	fmt.Println(listNode2)

	fmt.Println(addTwoNumbers(listNode1, listNode2))
}
