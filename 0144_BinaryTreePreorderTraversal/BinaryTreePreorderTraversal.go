package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	r := make([]int, 0)
	if root != nil {
		r = append(r, root.Val)
		r = append(r, preorderTraversal(root.Left)...)
		r = append(r, preorderTraversal(root.Right)...)
	}
	return r
}
