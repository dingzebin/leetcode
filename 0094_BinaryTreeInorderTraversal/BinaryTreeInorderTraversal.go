package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	r := []int{}
	if root != nil {
		if root.Left != nil {
			r = append(r, inorderTraversal(root.Left)...)
		}
		r = append(r, root.Val)
		if root.Right != nil {
			r = append(r, inorderTraversal(root.Right)...)
		}
	}
	return r
}
