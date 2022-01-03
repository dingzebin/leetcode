package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tmp := map[int]bool{
		head.Val: true,
	}
	prev := head
	for prev.Next != nil {
		if _, ok := tmp[prev.Next.Val]; ok {
			prev.Next = prev.Next.Next
		} else {
			tmp[prev.Next.Val] = true
			prev = prev.Next
		}
	}
	return head
}
