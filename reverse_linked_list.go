package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//loop
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}

	return pre
}

//recursion
func reverseList2(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	ele := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ele
}
