package p0002_add_two_numbers

import (
	"github.com/daominah/leetcode/structure"
)

type ListNode = structure.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	first := &ListNode{}
	loop := first
	carry := 0
	for {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		tmp := carry + val1 + val2
		loop.Val, carry = tmp%10, tmp/10
		if l1 == nil && l2 == nil {
			if carry == 0 {
				return first
			}
			loop.Next = &ListNode{Val: carry}
			return first
		}
		loop.Next = &ListNode{}
		loop = loop.Next
	}
}
