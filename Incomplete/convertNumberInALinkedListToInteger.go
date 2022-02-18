/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    curr, result := head, 0
    for curr != nil {
        result = result<<1
        result += curr.Val
        curr = curr.Next
    }
    return result
}
