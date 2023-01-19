package structure

// ListNode represents a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}
	first := &ListNode{}
	loop := first
	i := 0
	for {
		loop.Val = a[i]
		if i == len(a)-1 {
			return first
		}
		i += 1
		loop.Next = &ListNode{}
		loop = loop.Next
	}
}

func (first *ListNode) ToSlice() []int {
	var ret []int
	loop := first
	for loop != nil {
		ret = append(ret, loop.Val)
		loop = loop.Next
	}
	return ret
}
