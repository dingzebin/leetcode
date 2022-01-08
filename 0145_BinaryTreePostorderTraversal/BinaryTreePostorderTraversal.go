package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	r := make([]int, 0)
	if root != nil {
		r = append(r, postorderTraversal(root.Left)...)
		r = append(r, postorderTraversal(root.Right)...)
		r = append(r, root.Val)
	}
	return r
}
