/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var head *ListNode
    var prev *ListNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
    if l1.Val < l2.Val {
        head = l1
        l1 = l1.Next
    } else {
        head = l2
        l2 = l2.Next
    }
    prev = head
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            prev.Next = l1
            l1 = l1.Next
            prev = prev.Next
        } else {
            prev.Next = l2
            l2 = l2.Next
            prev = prev.Next
        }
    }

    if l1 != nil {
        prev.Next = l1
    }

    if l2 != nil {
        prev.Next = l2
    }
    
    return head
}
