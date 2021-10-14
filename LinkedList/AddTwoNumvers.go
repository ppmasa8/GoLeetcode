package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := new(ListNode)
	ans := root
	carry := 0

	for l1 != nil || l2 != nil {
		n := 0
		if l1 != nil && l2 != nil {
			n = (l1.Val + l2.Val + carry) % 10
			carry = (l1.Val + l2.Val + carry) / 10
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			n = (l1.Val + carry) % 10
			carry = (l1.Val + carry) / 10
			l1 = l1.Next
		} else if l2 != nil {
			n = (l2.Val + carry) % 10
			carry = (l2.Val + carry) / 10
			l2 = l2.Next
		} else {
			// Exception
		}
		ans.Val = n
		if l1 != nil || l2 != nil {
			ans.Next = new(ListNode)
			ans = ans.Next
		}
	}
	if carry == 1 {
		ans.Next = new(ListNode)
		ans.Next.Val++
	}

	return root
}
