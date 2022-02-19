package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := &ListNode{0, head}
	prev := res
	for head != nil {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		if prev.Next == head {
			prev = prev.Next
		} else {
			prev.Next = head.Next
		}
		head = head.Next
	}
	return res.Next
}

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := &ListNode{0, nil}
	tmp, isDuplicate := res, false
	for head.Next != nil {
		if head.Val == head.Next.Val {
			isDuplicate = true
		} else {
			if !isDuplicate {
				tmp.Next = head
				tmp = head
			}
			isDuplicate = false
		}
		head = head.Next
	}
	if !isDuplicate {
		tmp.Next = head
	} else {
		tmp.Next = nil
	}
	return res.Next
}
