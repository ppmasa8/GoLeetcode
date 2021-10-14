package main
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head

	ptr := dummyHead

	for ptr.Next != nil && ptr.Next.Next != nil {
		if ptr.Next.Val == ptr.Next.Next.Val {
			cp := ptr.Next
			for cp.Next != nil && cp.Val == cp.Next.Val{
				cp = cp.Next
			}
			ptr.Next = cp.Next
		} else {
			ptr = ptr.Next
		}
	}

	return dummyHead.Next
}
