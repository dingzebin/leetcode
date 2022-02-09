package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	size, cur := 0, head
	for cur != nil {
		size++
		cur = cur.Next
	}
	k %= size
	end := head
	for i := 0; i < k; i++ {
		end = end.Next
	}
	s := head
	for end.Next != nil {
		end = end.Next
		s = s.Next
	}
	end.Next = head
	res := s.Next
	s.Next = nil
	return res
}
