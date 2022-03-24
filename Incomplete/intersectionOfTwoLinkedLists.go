/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    longer, shorter := headA, headB
    p1, p2 := headA, headB
    cnt := 0
    for p1 != nil {
        cnt++
        p1 = p1.Next
    }
    
    for p2 != nil {
        cnt--
        p2 = p2.Next
    }
    
    if cnt < 0 {
        longer, shorter = headB, headA
        cnt = -cnt
    }
    
    for i := 0; i < cnt; i++ {
        longer = longer.Next
    }
    
    for longer != shorter {
        longer = longer.Next
        shorter = shorter.Next
    }
    
    return longer
}
