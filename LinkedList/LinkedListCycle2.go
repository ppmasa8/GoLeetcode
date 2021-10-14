package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow := head, head

	for fast != nil {
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return nil
		}
		slow = slow.Next

		if fast == slow {
			break
		}
	}

	if fast == nil {
		return nil
	}

	slow = head
	for fast != slow {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
