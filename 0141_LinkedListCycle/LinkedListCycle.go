package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}
	p1, p2 := head, head.Next
	for p1 != p2 {
		if p2 == nil && p2.Next == nil {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return true
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	p1, p2 := head, head.Next
	for i := 1; p2 != nil; i++ {
		if i%2 == 0 {
			p1 = p1.Next
		}
		p2 = p2.Next
		if p1 == p2 {
			return true
		}
	}
	return false
}
