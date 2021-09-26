/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Two pointers solution
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil { return false }
	fast, slow := head, head
	for fast != nil {
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
	return false
}