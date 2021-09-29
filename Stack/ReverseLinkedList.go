/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
var res *ListNode
tmp1, tmp2 := head, head
for tmp1 != nil {
tmp2 = tmp1.Next
tmp1.Next = res
res, tmp1 = tmp1, tmp2
}
return res
}