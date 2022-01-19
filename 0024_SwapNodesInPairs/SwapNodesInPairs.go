package main

func main() {
	
}


type ListNode struct {
    Val int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	sentry := &ListNode{Val: 0, Next: head}
	prevs := sentry
	for prevs.Next != nil && prevs.Next.Next != nil {
		n1, n2 := prevs.Next, prevs.Next.Next
		prevs.Next, n1.Next, n2.Next, prevs = n2, n2.Next, n1, n1
	}
	return sentry.Next
}