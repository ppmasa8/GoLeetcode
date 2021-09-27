/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil { return nil }
	cp := head
	for cp.Next != nil {
		if cp.Val == cp.Next.Val {
			cp.Next = cp.Next.Next
		} else {
			cp = cp.Next
		}
	}
	return head
}