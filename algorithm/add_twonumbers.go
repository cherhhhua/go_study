package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{
		Val: 0,
	}
	head := node
	n := 0
	for l1 != nil || l2 != nil {
		num := n
		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}

		node.Next = &ListNode{
			Val: num % 10,
		}

		n = num / 10
		node = node.Next
	}

	if n > 0 {
		node.Next = &ListNode{
			Val: n,
		}
	}

	return head.Next
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
