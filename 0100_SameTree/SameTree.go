package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p == nil && q == nil) || (p != nil && q != nil && p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)) {
		return true
	}
	return false
}
