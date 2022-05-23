package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	newHead := reverseList(next)
	next.Next = head
	head.Next = nil
	return newHead
}
