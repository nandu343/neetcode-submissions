/**
 * Definition for singly-linked list.
 * type ListNode struct {
 * Val int
 * Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    // FIX: Moving the nil/short check to the very top prevents fast.Next panic
    if head == nil || head.Next == nil || head.Next.Next == nil {
        return
    }

    fast := head
    slow := head
    
    // Find the middle of the linked list
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }

    // Split the list into two halves
    head2 := slow.Next
    slow.Next = nil
    
    // FIX: Standard, panic-safe iterative linked list reversal
    reverse := func (h *ListNode) *ListNode {
        var prev *ListNode
        curr := h
        for curr != nil {
            nextTemp := curr.Next
            curr.Next = prev
            prev = curr
            curr = nextTemp
        }
        return prev
    }

    head2 = reverse(head2)

    // Merge the two halves together
    curr := head
    for head2 != nil {
        temp1 := curr.Next
        temp2 := head2.Next
        
        curr.Next = head2
        head2.Next = temp1
        
        curr = temp1
        head2 = temp2
    }
}