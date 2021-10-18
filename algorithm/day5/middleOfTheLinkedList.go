package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// import "fmt"
func middleNode(head *ListNode) *ListNode {
	cnt := 1
	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
		cnt++
		// fmt.Println(cnt)
	}

	for i := 0; i < cnt / 2; i++{
		head = head.Next
	}

	return head
}
