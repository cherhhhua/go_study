package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var slow *ListNode
	var fast *ListNode
	var pre *ListNode
	var prepre *ListNode

	if head == nil || head.Next == nil {
		return true
	}
	slow, fast, pre = head, head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
		pre.Next = prepre
		prepre = pre
	}
	if fast != nil {
		slow = slow.Next
	}
	for pre != nil && slow != nil {
		if pre.Val != slow.Val {
			return false
		}
		pre = pre.Next
		slow = slow.Next
	}
	return true
}
