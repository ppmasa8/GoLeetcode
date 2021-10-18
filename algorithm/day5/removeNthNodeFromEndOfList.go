package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cnt, i := 0, 0
	tmp, cp := head, head

	for tmp.Next != nil {
		tmp = tmp.Next
		cnt++
	}

	if cnt == 0 {
		head = nil
	} else if cnt == 1 {
		if n == 1 {
			head.Next = nil
		} else if n == 2 {
			head = head.Next
		}
	} else if cnt - n < 0 {
		head = head.Next
	}

	for cp.Next != nil {
		if i == cnt - n {
			cp.Next = cp.Next.Next
		} else {
			cp = cp.Next
		}
		i++
	}

	return head
}
