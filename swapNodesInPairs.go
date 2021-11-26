/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    node := head
    
    if node != nil && node.Next != nil {
        temp := node.Val
        node.Val = node.Next.Val
        node.Next.Val = temp
        swapPairs(head.Next.Next)
    }
    return head
}
