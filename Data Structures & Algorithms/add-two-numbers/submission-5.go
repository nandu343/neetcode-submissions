/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	curr := &ListNode{}
	head := curr
	a, b := l1, l2
	carryover := 0
	for a != nil && b != nil {
		curr.Next = &ListNode{}
		curr = curr.Next
		sum := a.Val + b.Val + carryover
		if sum >= 10 {
			sum = sum % 10
			curr.Val = sum
			carryover = 1
		} else {
			curr.Val = sum
			carryover = 0
		}
		a = a.Next
		b = b.Next
	}


	for a != nil {
		curr.Next = &ListNode{}
		curr = curr.Next
		curr.Val = (a.Val + carryover) % 10
		if (a.Val + carryover) >= 10 {
			carryover = 1
		} else {
			carryover = 0
		}
		a = a.Next
	}
	for b != nil {
		curr.Next = &ListNode{}
		curr = curr.Next
		curr.Val = (b.Val + carryover) % 10
		if (b.Val + carryover) >= 10 {
			carryover = 1
		} else {
			carryover = 0
		}
		b = b.Next
	}

	if a == nil && b == nil && carryover == 1{
		curr.Next = &ListNode{Val: 1}
	}

	return head.Next
}
