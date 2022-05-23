/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    if head == nil || head.Val == 0 && head.Next == nil {
        return nil
    }
    curr := head
    pointer := &ListNode{}
    for curr.Next != nil {
        if curr.Val == 0 {
            pointer = curr
            curr = curr.Next
        } else {
            pointer.Val += curr.Val
            *curr = *curr.Next
        }
    }
    pointer.Next = nil
    return head
}
